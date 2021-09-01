package sqlp

import (
	"database/sql"
)

// @author valor.

func Open(db *sql.DB) (*DBSession, error) {

	err := db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return &DBSession{
		database: db,
		//aotCachedStmt: make(map[string]*fakeStmt),
		jitCachedStmt: make(map[string]*fakeStmt),
	}, nil
}
