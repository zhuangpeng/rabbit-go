package config

import (
	"database/sql"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/zeromicro/go-zero/core/logx"
)

type CasbinConf struct {
    ModelText string `json:"ModelText"`
}

func (l CasbinConf) NewCasbin(dbType *sql.DB, driver, table string) (*casbin.Enforcer, error) {
	adapter, err := sqladapter.NewAdapter(dbType, driver, table)
	logx.Must(err)

	var text string
	if l.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = l.ModelText
	}

	m, err := model.NewModelFromString(text)
	logx.Must(err)

	enforcer, err := casbin.NewEnforcer(m, adapter)
	logx.Must(err)

	err = enforcer.LoadPolicy()
	logx.Must(err)

	return enforcer, nil
}
