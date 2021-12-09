package sqlp

import (
	goSql "database/sql"
)

// @author valor.

// TxSession provides a set of extensions on database/sql
type TxSession struct {
	db *DBSession
	tx *goSql.Tx
}

// ExecSql execute sql at TxSession
func (t *TxSession) ExecSql(sql string, args ...interface{}) (goSql.Result, error) {
	return t.ExecSqlDirect(sql, args...)
}

// ExecSql execute sql at TxSession without JIT
func (t *TxSession) ExecSqlDirect(sql string, args ...interface{}) (goSql.Result, error) {
	stmt := fakeStmt(sql)
	return stmt.execAtTx(t, args...)
}

// QuerySql execute query sql at TxSession
func (t *TxSession) QuerySql(dest interface{}, sql string, args ...interface{}) error {
	return t.QuerySqlDirect(dest, sql, args...)
}

// QuerySql execute query sql at TxSession without JIT
func (t *TxSession) QuerySqlDirect(dest interface{}, sql string, args ...interface{}) error {
	stmt := fakeStmt(sql)
	return t.queryTx(dest, stmt, args...)
}

func (t *TxSession) queryTx(dest interface{}, stmt fakeStmt, args ...interface{}) error {
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

// Commit commits the transaction.
func (t *TxSession) Commit() error {
	return t.tx.Commit()
}

// Rollback aborts the transaction.
func (t *TxSession) Rollback() error {
	return t.tx.Rollback()
}
