package config

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DatabaseConf struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
	Type     string
}

func (c DatabaseConf) NewSqlConn() (sqlConn sqlx.SqlConn) {
	switch c.Type {
	case "mysql":
		return sqlx.NewMysql(c.GetDSN())
	case "postgres":
		return postgres.New(c.GetDSN())
	default:
		return sqlx.NewMysql(c.GetDSN())
	}
}

func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

func (c DatabaseConf) GetDSN() string {
	switch c.Type {
	case "mysql":
		return c.MysqlDSN()
	case "postgres":
		return c.PostgresDSN()
	default:
		return "postgres"
	}
}
