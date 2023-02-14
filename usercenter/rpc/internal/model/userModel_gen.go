// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	globalkey "github.com/zhuangpeng/rabbit-go/pkg/globalKey"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/dbx"
	"github.com/zhuangpeng/rabbit-go/pkg/statuserr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cacheUserCenterUserIdPrefix     = "cache:userCenter:user:id:"
	cacheUserCenterUserEmailPrefix  = "cache:userCenter:user:email:"
	cacheUserCenterUserMobilePrefix = "cache:userCenter:user:mobile:"
	cacheUserCenterUserNamePrefix   = "cache:userCenter:user:name:"
)

type (
	userModel interface {
		Insert(ctx context.Context, dropZeroValue bool, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		FindOneByName(ctx context.Context, name string) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error
		Delete(ctx context.Context, id string) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id          string       `db:"id"`           // 用户标识
		Name        string       `db:"name"`         // 用户名
		Password    string       `db:"password"`     // 密码
		Nickname    string       `db:"nickname"`     // 昵称
		SideMode    string       `db:"side_mode"`    // 布局方式
		BaseColor   string       `db:"base_color"`   // 后台页面色调
		ActiveColor string       `db:"active_color"` // 当前激活颜色设定
		Mobile      string       `db:"mobile"`       // 联系电话
		Email       string       `db:"email"`        // 邮箱
		Avatar      string       `db:"avatar"`       // 头像
		Status      int64        `db:"status"`       // 状态 1--正常 2--禁用
		Deleted     int64        `db:"deleted"`      // 是否删除 1--否  2--是
		CreatedAt   time.Time    `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullTime `db:"updated_at"`   // 修改时间
		DeletedAt   sql.NullTime `db:"deleted_at"`   // 删除时间
		Revision    int64        `db:"revision"`     // 乐观锁修订版本号
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id string) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userCenterUserEmailKey := fmt.Sprintf("%s%v", cacheUserCenterUserEmailPrefix, data.Email)
	userCenterUserIdKey := fmt.Sprintf("%s%v", cacheUserCenterUserIdPrefix, id)
	userCenterUserMobileKey := fmt.Sprintf("%s%v", cacheUserCenterUserMobilePrefix, data.Mobile)
	userCenterUserNameKey := fmt.Sprintf("%s%v", cacheUserCenterUserNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userCenterUserEmailKey, userCenterUserIdKey, userCenterUserMobileKey, userCenterUserNameKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	userCenterUserIdKey := fmt.Sprintf("%s%v", cacheUserCenterUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userCenterUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and deleted = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	userCenterUserEmailKey := fmt.Sprintf("%s%v", cacheUserCenterUserEmailPrefix, email)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userCenterUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? and deleted = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	userCenterUserMobileKey := fmt.Sprintf("%s%v", cacheUserCenterUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userCenterUserMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? and deleted = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name string) (*User, error) {
	userCenterUserNameKey := fmt.Sprintf("%s%v", cacheUserCenterUserNamePrefix, name)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, userCenterUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? and deleted = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Insert
// if dropZeroValue is true then filter the zero value with struct
// if session not nil, the execute statement with transactions
func (m *defaultUserModel) Insert(ctx context.Context, dropZeroValue bool, session sqlx.Session, data *User) (sql.Result, error) {
	userCenterUserEmailKey := fmt.Sprintf("%s%v", cacheUserCenterUserEmailPrefix, data.Email)
	userCenterUserIdKey := fmt.Sprintf("%s%v", cacheUserCenterUserIdPrefix, data.Id)
	userCenterUserMobileKey := fmt.Sprintf("%s%v", cacheUserCenterUserMobilePrefix, data.Mobile)
	userCenterUserNameKey := fmt.Sprintf("%s%v", cacheUserCenterUserNamePrefix, data.Name)
	if dropZeroValue {
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			query, args, err := dbx.Insert(*data)
			if err != nil {
				return nil, err
			}
			if session != nil {
				return session.ExecCtx(ctx, query, args...)
			}
			return conn.ExecCtx(ctx, query, args...)
		}, userCenterUserEmailKey, userCenterUserIdKey, userCenterUserMobileKey, userCenterUserNameKey)
	} else {
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
			if session != nil {
				return session.ExecCtx(ctx, query, data.Id, data.Name, data.Password, data.Nickname, data.SideMode, data.BaseColor, data.ActiveColor, data.Mobile, data.Email, data.Avatar, data.Status, data.Deleted, data.DeletedAt, data.Revision)
			}
			return conn.ExecCtx(ctx, query, data.Id, data.Name, data.Password, data.Nickname, data.SideMode, data.BaseColor, data.ActiveColor, data.Mobile, data.Email, data.Avatar, data.Status, data.Deleted, data.DeletedAt, data.Revision)
		}, userCenterUserEmailKey, userCenterUserIdKey, userCenterUserMobileKey, userCenterUserNameKey)
	}
}

// Update when session not nil update with transactions
func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, newData *User) (sql.Result, error) {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return nil, err
	}
	userCenterUserEmailKey := fmt.Sprintf("%s%v", cacheUserCenterUserEmailPrefix, data.Email)
	userCenterUserIdKey := fmt.Sprintf("%s%v", cacheUserCenterUserIdPrefix, data.Id)
	userCenterUserMobileKey := fmt.Sprintf("%s%v", cacheUserCenterUserMobilePrefix, data.Mobile)
	userCenterUserNameKey := fmt.Sprintf("%s%v", cacheUserCenterUserNamePrefix, data.Name)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Nickname, newData.SideMode, newData.BaseColor, newData.ActiveColor, newData.Mobile, newData.Email, newData.Avatar, newData.Status, newData.Deleted, newData.DeletedAt, newData.Revision, newData.Id)
		}
		return conn.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Nickname, newData.SideMode, newData.BaseColor, newData.ActiveColor, newData.Mobile, newData.Email, newData.Avatar, newData.Status, newData.Deleted, newData.DeletedAt, newData.Revision, newData.Id)
	}, userCenterUserEmailKey, userCenterUserIdKey, userCenterUserMobileKey, userCenterUserNameKey)

}

// UpdateWithVersion update with revision control, if the session is not nil then with transactions
func (m *defaultUserModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *User) error {

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	oldRevision := data.Revision
	newData.Revision += 1

	userCenterUserEmailKey := fmt.Sprintf("%s%v", cacheUserCenterUserEmailPrefix, data.Email)
	userCenterUserIdKey := fmt.Sprintf("%s%v", cacheUserCenterUserIdPrefix, data.Id)
	userCenterUserMobileKey := fmt.Sprintf("%s%v", cacheUserCenterUserMobilePrefix, data.Mobile)
	userCenterUserNameKey := fmt.Sprintf("%s%v", cacheUserCenterUserNamePrefix, data.Name)

	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and revision = ? ", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Nickname, newData.SideMode, newData.BaseColor, newData.ActiveColor, newData.Mobile, newData.Email, newData.Avatar, newData.Status, newData.Deleted, newData.DeletedAt, newData.Revision, newData.Id, oldRevision)
		}
		return conn.ExecCtx(ctx, query, newData.Name, newData.Password, newData.Nickname, newData.SideMode, newData.BaseColor, newData.ActiveColor, newData.Mobile, newData.Email, newData.Avatar, newData.Status, newData.Deleted, newData.DeletedAt, newData.Revision, newData.Id, oldRevision)
	}, userCenterUserEmailKey, userCenterUserIdKey, userCenterUserMobileKey, userCenterUserNameKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return statuserr.NewNotFoundError(i18n.TargetNotFound)
	}

	return nil
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserCenterUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and deleted = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
