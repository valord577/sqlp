package sqlp

import (
	goSql "database/sql"
)

// @author valor.

type TxSession struct {
	db *DBSession
	tx *goSql.Tx
}

/*
func (t *TxSession) Exec(id string, args ...interface{}) (goSql.Result, error) {
	stmt, err := t.db.getAotCachedStmt(id)
	if err != nil {
		return nil, err
	}
	return stmt.execAtTx(t, args...)
}
*/

func (t *TxSession) ExecSql(sql string, args ...interface{}) (goSql.Result, error) {
	stmt, err := t.db.getJitCachedStmt(sql)
	if err != nil {
		return nil, err
	}
	return stmt.execAtTx(t, args...)
}

/* 
func (t *TxSession) Query(dest interface{}, id string, args ...interface{}) error {
	stmt, err := t.db.getAotCachedStmt(id)
	if err != nil {
		return err
	}
	return t.queryTx(dest, stmt, args...)
}
*/

func (t *TxSession) QuerySql(dest interface{}, sql string, args ...interface{}) error {
	stmt, err := t.db.getJitCachedStmt(sql)
	if err != nil {
		return err
	}
	return t.queryTx(dest, stmt, args...)
}

func (t *TxSession) queryTx(dest interface{}, stmt *fakeStmt, args ...interface{}) error {
	rs, err := stmt.queryAtTx(t, args...)
	if err != nil {
		return err
	}
	// if something happens here, we want to make sure the rows are Closed
	defer func(rs *goSql.Rows) {
		_ = rs.Close()
	}(rs)
	
	return scanAny(dest, rs)
}

func (t *TxSession) Commit() error {
	return t.tx.Commit()
}

func (t *TxSession) Rollback() error {
	return t.tx.Rollback()
}
