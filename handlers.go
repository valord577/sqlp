package sqlp

// @author valor.

type RawSqlHandler interface {
	ToRawSql(sql string, args ...interface{}) (string, error)
}
