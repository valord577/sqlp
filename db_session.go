package sqlp

import (
	goSql "database/sql"
)

// @author valor.

// DBSession provides a set of extensions on database/sql
type DBSession struct {
	database *goSql.DB
	//aotCachedStmt map[string]*fakeStmt
	jitCachedStmt map[string]*fakeStmt

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

/*
func (s *DBSession) getAotCachedStmt(id string) (*fakeStmt, error) {
	fStmt, ok := s.aotCachedStmt[id]
	if ok {
		return fStmt, nil
	}
	return nil, errors.New("not found sql, ID: " + id)
}
*/

func (s *DBSession) getJitCachedStmt(sql string) (*fakeStmt, error) {
	fStmt, ok := s.jitCachedStmt[sql]
	if !ok {
		stmt, err := s.database.Prepare(sql)
		if err != nil {
			return nil, err
		}

		fStmt = &fakeStmt{
			stmtStr: sql,
			stmtSql: stmt,
		}
		s.jitCachedStmt[sql] = fStmt
	}
	return fStmt, nil
}

/*
func (s *DBSession) Exec(id string, args ...interface{}) (goSql.Result, error) {
	stmt, err := s.getAotCachedStmt(id)
	if err != nil {
		return nil, err
	}
	return stmt.execAtDB(s, args...)
}
*/

// ExecSql execute sql at DBSession
func (s *DBSession) ExecSql(sql string, args ...interface{}) (goSql.Result, error) {
	stmt, err := s.getJitCachedStmt(sql)
	if err != nil {
		return nil, err
	}
	return stmt.execAtDB(s, false, args...)
}

// ExecSql execute sql at DBSession without JIT
func (s *DBSession) ExecSqlDirect(sql string, args ...interface{}) (goSql.Result, error) {
	stmt := &fakeStmt{
		stmtStr: sql,
	}
	return stmt.execAtDB(s, true, args...)
}

/*
func (s *DBSession) Query(dest interface{}, id string, args ...interface{}) error {
	stmt, err := s.getAotCachedStmt(id)
	if err != nil {
		return err
	}
	return s.query(dest, stmt, args...)
}
*/

// QuerySql execute query sql at DBSession
func (s *DBSession) QuerySql(dest interface{}, sql string, args ...interface{}) error {
	stmt, err := s.getJitCachedStmt(sql)
	if err != nil {
		return err
	}
	return s.query(dest, false, stmt, args...)
}

// QuerySql execute query sql at DBSession without JIT
func (s *DBSession) QuerySqlDirect(dest interface{}, sql string, args ...interface{}) error {
	stmt := &fakeStmt{
		stmtStr: sql,
	}
	return s.query(dest, true, stmt, args...)
}

func (s *DBSession) query(dest interface{}, direct bool, stmt *fakeStmt, args ...interface{}) error {
	rs, err := stmt.queryAtDB(s, direct, args...)
	if err != nil {
		return err
	}
	// if something happens here, we want to make sure the rows are Closed
	defer func(rs *goSql.Rows) {
		_ = rs.Close()
	}(rs)

	return scanAny(dest, rs)
}
