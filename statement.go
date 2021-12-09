package sqlp

import (
	goSql "database/sql"
	"errors"
)

// @author valor.

var errNoRawSqlHandler = errors.New("raw sql handler is nil")

type fakeStmt string

func (f fakeStmt) queryAtDB(dbSession *DBSession, args ...interface{}) (*goSql.Rows, error) {
	if dbSession.disablePreparedStmtAtDBSession {
		return f.queryRawAtDB(dbSession, args...)
	}
	return dbSession.database.Query(string(f), args...)
}

func (f fakeStmt) queryRawAtDB(dbSession *DBSession, args ...interface{}) (*goSql.Rows, error) {
	if dbSession.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := dbSession.rawSqlHandler.ToRawSql(string(f), args...)
	if err != nil {
		return nil, err
	}
	return dbSession.database.Query(rawSql)
}

func (f fakeStmt) queryAtTx(txSession *TxSession, args ...interface{}) (*goSql.Rows, error) {
	if txSession.db.disablePreparedStmtAtTxSession {
		return f.queryRawAtTx(txSession, args...)
	}
	return txSession.tx.Query(string(f), args...)
}

func (f fakeStmt) queryRawAtTx(txSession *TxSession, args ...interface{}) (*goSql.Rows, error) {
	if txSession.db.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := txSession.db.rawSqlHandler.ToRawSql(string(f), args...)
	if err != nil {
		return nil, err
	}
	return txSession.tx.Query(rawSql)
}

func (f fakeStmt) execAtDB(dbSession *DBSession, args ...interface{}) (goSql.Result, error) {
	if dbSession.disablePreparedStmtAtDBSession {
		return f.execRawAtDB(dbSession, args...)
	}
	return dbSession.database.Exec(string(f), args...)
}

func (f fakeStmt) execRawAtDB(dbSession *DBSession, args ...interface{}) (goSql.Result, error) {
	if dbSession.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := dbSession.rawSqlHandler.ToRawSql(string(f), args...)
	if err != nil {
		return nil, err
	}
	return dbSession.database.Exec(rawSql)
}

func (f fakeStmt) execAtTx(txSession *TxSession, args ...interface{}) (goSql.Result, error) {
	if txSession.db.disablePreparedStmtAtTxSession {
		return f.execRawAtTx(txSession, args...)
	}
	return txSession.tx.Exec(string(f), args...)
}

func (f fakeStmt) execRawAtTx(txSession *TxSession, args ...interface{}) (goSql.Result, error) {
	if txSession.db.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := txSession.db.rawSqlHandler.ToRawSql(string(f), args...)
	if err != nil {
		return nil, err
	}
	return txSession.tx.Exec(rawSql)
}
