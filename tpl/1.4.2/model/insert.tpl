// Insert
// if dropZeroValue is true then filter the zero value with struct
// if session not nil, the execute statement with transactions
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, dropZeroValue bool, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
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
		}, {{.keyValues}})
	} else {
		return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		if session != nil{
			return session.ExecCtx(ctx,query,{{.expressionValues}})
		}
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
		}, {{.keyValues}})
	}{{else}}
	if dropZeroValue {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		if session != nil{
			return session.ExecCtx(ctx,query,{{.expressionValues}})
		}
		return m.conn.ExecCtx(ctx, query, {{.expressionValues}})
	} else {
		query, args, err := dbx.Insert(*data)
		if err != nil {
			return nil, err
		}			
		if session != nil{
			return session.ExecCtx(ctx, query ,args...)
		}
		return m.conn.ExecCtx(ctx, query, args...)	
	}{{end}}
}
