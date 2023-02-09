package dbx

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

func Insert(s interface{}) (sql string, args []interface{}, err error) {
	rv := reflect.ValueOf(s)
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		logx.Infof("params must a STRUCT, you past a: %v", rv)
		return "", nil, errors.New("input must be struct")
	}
	keys, vals := sKV(rv)
	sql, args, err = squirrel.Insert(rt.Name()).Columns(keys...).Values(vals...).ToSql()
	return
}

// // UpdateBuilder return update query segment
// func UpdateBuilder(s, pred interface{}, args []interface{}) (*squirrel.UpdateBuilder, error) {
// 	rv := reflect.ValueOf(s)
// 	rt := reflect.TypeOf(s)
// 	if rt.Kind() != reflect.Struct {
// 		logx.Infof("params must a STRUCT, you past a: %v", rv)
// 		return nil, errors.New("input must be struct")
// 	}
// 	updateMap := updateMap(rv)
// 	sqBuilder := squirrel.Update(rt.Name()).SetMap(updateMap).Where(pred, args...)
// 	return &sqBuilder, nil
// }


func sKV(v reflect.Value) ([]string, []interface{}) {
	var keys []string
	var values []interface{}
	t := v.Type()
	for n := 0; n < t.NumField(); n++ {
		tf := t.Field(n)
		vf := v.Field(n)
		//忽略非导出字段
		if tf.Anonymous {
			continue
		}
		//忽略无效、零值字段
		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type()).Interface()) {
			continue
		}
		for vf.Type().Kind() == reflect.Ptr {
			vf = vf.Elem()
		}
		//有时候根据需求会组合struct，这里处理下，支持获取嵌套的struct tag和value
		//如果字段值是time类型之外的struct，递归获取keys和values
		// if vf.Kind() == reflect.Struct && tf.Type.Name() != "Time" {
		// 	cKeys, cValues := sKV(vf)
		// 	keys = append(keys, cKeys...)
		// 	values = append(values, cValues...)
		// 	continue
		// }

		//根据字段的json tag获取key，忽略无tag字段
		key := strings.Split(tf.Tag.Get("db"), ",")[0]
		if key == "" {
			continue
		}
		value := format(vf)
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

func updateMap(v reflect.Value) map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	t := v.Type()
	for n := 0; n < t.NumField(); n++ {
		tf := t.Field(n)
		vf := v.Field(n)
		//忽略非导出字段
		if tf.Anonymous {
			continue
		}
		//忽略无效、零值字段
		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type()).Interface()) {
			continue
		}
		for vf.Type().Kind() == reflect.Ptr {
			vf = vf.Elem()
		}
		//有时候根据需求会组合struct，这里处理下，支持获取嵌套的struct tag和value
		//如果字段值是time类型之外的struct，递归获取keys和values
		// if vf.Kind() == reflect.Struct && tf.Type.Name() != "Time" {
		// 	cKeys, cValues := sKV(vf)
		// 	keys = append(keys, cKeys...)
		// 	values = append(values, cValues...)
		// 	continue
		// }

		//根据字段的json tag获取key，忽略无tag字段
		key := strings.Split(tf.Tag.Get("db"), ",")[0]
		if key == "" {
			continue
		}
		value := format(vf)
		ret[key] = value
	}
	return ret
}

func format(v reflect.Value) interface{} {
	//断言出time类型直接转unix时间戳
	if t, ok := v.Interface().(time.Time); ok {
		return fmt.Sprintf("FROM_UNIXTIME(%d)", t.Unix())
	}
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf(`%s`, v.Interface())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return fmt.Sprintf(`%d`, v.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%f`, v.Interface())
	// 如果是切片类型，遍历元素，递归格式化成"(, , , )"形式
	// case reflect.Slice:
	//     var values []string
	//     for i := 0; i < v.Len(); i++ {
	//         values = append(values, format(v.Index(i)))
	//     }
	//     return fmt.Sprintf(`(%s)`, strings.Join(values, ","))
	// 接口类型剥一层递归
	case reflect.Interface:
		return format(v.Elem())
	}
	return ""
}
