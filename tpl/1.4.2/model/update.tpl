
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

// custom templdate
// feat: update with Database transactions
func (m *default{{.upperStartCamelObject}}Model) UpdateTx(ctx context.Context,session sqlx.Session, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}})  (sql.Result,error) {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return nil, err
	}

	{{end}}	{{.keys}}
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}

// custom templdate
// feat: update with revision control, if the session is nil without transactions 
func (m *default{{.upperStartCamelObject}}Model) UpdateWithVersion(ctx context.Context,session sqlx.Session,{{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	
	var sqlResult sql.Result
	var err error
	
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}
	{{end}}
	
	oldRevision := data.Revision
	newData.Revision += 1

	{{if .withCache}}{{.keys}}
	
	sqlResult,err =  m.ExecCtx(ctx,func(ctx context.Context,conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}},oldRevision)
	}
	return conn.ExecCtx(ctx,query, {{.expressionValues}},oldRevision)
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		sqlResult,err  =  session.ExecCtx(ctx,query, {{.expressionValues}},oldRevision)
	}else{
		sqlResult,err  =  m.conn.ExecCtx(ctx,query, {{.expressionValues}},oldRevision)
	}
	{{end}}
	if err != nil {
		return err
	}
	updateCount , err := sqlResult.RowsAffected()
	if err != nil{
		return err
	}
	if updateCount == 0 {
		return  xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil {{end}}
}
