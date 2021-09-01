package example

import (
	"database/sql"

	"github.com/valord577/sqlp"
)

// @author valor.

/*

create table `tc`
(
	id int auto_increment
		primary key,
	name varchar(16) default '' not null
);

*/

const (
	sqlType = ""
	sqlDsn  = ""
)

func doTest(f func(s *sqlp.DBSession) error) {
	db, err := sql.Open(sqlType, sqlDsn)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	dbSession, err := sqlp.Open(db)
	if err != nil {
		panic(err)
	}

	err = f(dbSession)
	if err != nil {
		panic(err)
	}
}
