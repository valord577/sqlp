package example

import (
	"fmt"
	"testing"

	"github.com/valord577/sqlp"
)

// @author valor.

func TestExecSqlAtDB(t *testing.T) {
	doTest(func(s *sqlp.DBSession) error {
		
		sql := "insert into `tc` (`name`) values (?);"
		ret, err := s.ExecSql(sql, "first name")
		if err != nil {
			return err
		}

		lastInsertId, err := ret.LastInsertId()
		if err != nil {
			return err
		}
		rowsAffected, err := ret.RowsAffected()
		if err != nil {
			return err
		}

		fmt.Printf("LastInsertId: %d, RowsAffected: %d\n", lastInsertId, rowsAffected)
		return nil
	})
}

func TestExecSqlAtTx(t *testing.T) {
	doTest(func(s *sqlp.DBSession) error {

		tx, err := s.BeginTx()
		if err != nil {
			return err
		}

		sql := "insert into `tc` (`name`) values (?);"
		ret, err := tx.ExecSql(sql, "first name")
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()

		lastInsertId, err := ret.LastInsertId()
		if err != nil {
			return err
		}
		rowsAffected, err := ret.RowsAffected()
		if err != nil {
			return err
		}

		fmt.Printf("LastInsertId: %d, RowsAffected: %d\n", lastInsertId, rowsAffected)
		return nil
	})
}
