package sqlp

import (
	goSql "database/sql"
	"errors"
)

// @author valor.

var errNoRawSqlHandler = errors.New("raw sql handler is nil")

type fakeStmt struct {
	stmtStr string
	stmtSql *goSql.Stmt
}

func (f *fakeStmt) queryAtDB(dbSession *DBSession, direct bool, args ...interface{}) (*goSql.Rows, error) {
	if dbSession.disablePreparedStmtAtDBSession {
		return f.queryRawAtDB(dbSession, args...)
	}

	if direct {
		return dbSession.database.Query(f.stmtStr, args...)
	}
	return f.stmtSql.Query(args...)
}

func (f *fakeStmt) queryRawAtDB(dbSession *DBSession, args ...interface{}) (*goSql.Rows, error) {
	if dbSession.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := dbSession.rawSqlHandler.ToRawSql(f.stmtStr, args...)
	if err != nil {
		return nil, err
	}
	return dbSession.database.Query(rawSql)
}

func (f *fakeStmt) queryAtTx(txSession *TxSession, direct bool, args ...interface{}) (*goSql.Rows, error) {
	if txSession.db.disablePreparedStmtAtTxSession {
		return f.queryRawAtTx(txSession, args...)
	}

	if direct {
		return txSession.tx.Query(f.stmtStr, args...)
	}
	return txSession.tx.Stmt(f.stmtSql).Query(args...)
}

func (f *fakeStmt) queryRawAtTx(txSession *TxSession, args ...interface{}) (*goSql.Rows, error) {
	if txSession.db.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := txSession.db.rawSqlHandler.ToRawSql(f.stmtStr, args...)
	if err != nil {
		return nil, err
	}
	return txSession.tx.Query(rawSql)
}

func (f *fakeStmt) execAtDB(dbSession *DBSession, direct bool, args ...interface{}) (goSql.Result, error) {
	if dbSession.disablePreparedStmtAtDBSession {
		return f.execRawAtDB(dbSession, args...)
	}

	if direct {
		return dbSession.database.Exec(f.stmtStr, args...)
	}
	return f.stmtSql.Exec(args...)
}

func (f *fakeStmt) execRawAtDB(dbSession *DBSession, args ...interface{}) (goSql.Result, error) {
	if dbSession.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := dbSession.rawSqlHandler.ToRawSql(f.stmtStr, args...)
	if err != nil {
		return nil, err
	}
	return dbSession.database.Exec(rawSql)
}

func (f *fakeStmt) execAtTx(txSession *TxSession, direct bool,  args ...interface{}) (goSql.Result, error) {
	if txSession.db.disablePreparedStmtAtTxSession {
		return f.execRawAtTx(txSession, args...)
	}

	if direct {
		return txSession.tx.Exec(f.stmtStr, args...)
	}
	return txSession.tx.Stmt(f.stmtSql).Exec(args...)
}

func (f *fakeStmt) execRawAtTx(txSession *TxSession, args ...interface{}) (goSql.Result, error) {
	if txSession.db.rawSqlHandler == nil {
		return nil, errNoRawSqlHandler
	}
	rawSql, err := txSession.db.rawSqlHandler.ToRawSql(f.stmtStr, args...)
	if err != nil {
		return nil, err
	}
	return txSession.tx.Exec(rawSql)
}
