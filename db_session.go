package sqlp

import (
	goSql "database/sql"
)

// @author valor.

// DBSession provides a set of extensions on database/sql
type DBSession struct {
	database *goSql.DB

	disablePreparedStmtAtDBSession bool
	disablePreparedStmtAtTxSession bool
	// use raw sql, when disable prepared stmt
	rawSqlHandler RawSqlHandler
}

// EnablePrepStmtAtDBSession enable prepared statement at DBSession
func (s *DBSession) EnablePrepStmtAtDBSession() {
	s.disablePreparedStmtAtDBSession = false
}

// DisablePrepStmtAtDBSession disable prepared statement at DBSession
func (s *DBSession) DisablePrepStmtAtDBSession() {
	s.disablePreparedStmtAtDBSession = true
}

// EnablePrepStmtAtTxSession enable prepared statement at TxSession
func (s *DBSession) EnablePrepStmtAtTxSession() {
	s.disablePreparedStmtAtTxSession = false
}

// DisablePrepStmtAtTxSession disable prepared statement at TxSession
func (s *DBSession) DisablePrepStmtAtTxSession() {
	s.disablePreparedStmtAtTxSession = true
}

// UseRawSqlHandler formats prepared statement to raw sql
func (s *DBSession) UseRawSqlHandler(h RawSqlHandler) {
	s.rawSqlHandler = h
}

// BeginTx begin transaction
func (s *DBSession) BeginTx() (*TxSession, error) {
	tx, err := s.database.Begin()
	if err != nil {
		return nil, err
	}

	return &TxSession{
		db: s,
		tx: tx,
	}, nil
}

// ExecSql execute sql at DBSession
func (s *DBSession) ExecSql(sql string, args ...interface{}) (goSql.Result, error) {
	return s.ExecSqlDirect(sql, args...)
}

// ExecSql execute sql at DBSession without JIT
func (s *DBSession) ExecSqlDirect(sql string, args ...interface{}) (goSql.Result, error) {
	stmt := fakeStmt(sql)
	return stmt.execAtDB(s, args...)
}

// QuerySql execute query sql at DBSession
func (s *DBSession) QuerySql(dest interface{}, sql string, args ...interface{}) error {
	return s.QuerySqlDirect(dest, sql, args...)
}

// QuerySql execute query sql at DBSession without JIT
func (s *DBSession) QuerySqlDirect(dest interface{}, sql string, args ...interface{}) error {
	stmt := fakeStmt(sql)
	return s.query(dest, stmt, args...)
}

func (s *DBSession) query(dest interface{}, stmt fakeStmt, args ...interface{}) error {
	rs, err := stmt.queryAtDB(s, args...)
	if err != nil {
		return err
	}
	// if something happens here, we want to make sure the rows are Closed
	defer func(rs *goSql.Rows) {
		_ = rs.Close()
	}(rs)

	return scanAny(dest, rs)
}
