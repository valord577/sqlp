package sqlp

import (
	"database/sql"
)

// @author valor.

// Open returns DBSession
func Open(db *sql.DB) (*DBSession, error) {
	err := db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	return &DBSession{database: db}, nil
}
