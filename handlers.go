package sqlp

// @author valor.

// RawSqlHandler formats stmt as raw sql,
//   more info: http://go-database-sql.org/prepared.html
type RawSqlHandler interface {
	ToRawSql(sql string, args ...interface{}) (string, error)
}
