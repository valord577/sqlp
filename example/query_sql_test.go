package example

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/valord577/sqlp"
)

// @author valor.

func TestQuerySqlAtDB(t *testing.T) {
	doTest(func(s *sqlp.DBSession) error {

		type Tc struct {
			ID   uint   `sqlp:"id"`
			Name string `sqlp:"name"`
		}

		var (
			err  error
			sql1 = "select * from `tc` where `id` >= ?"
			sql2 = "select `id` from `tc` where `id` >= ?"
		)

		var a []Tc
		err = s.QuerySql(&a, sql1, 1)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", a)

		var b []*Tc
		err = s.QuerySql(&b, sql1, 1)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", b)

		var c []map[string]interface{}
		err = s.QuerySql(&c, sql1, 1)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", c)

		var d []*map[string]interface{}
		err = s.QuerySql(&d, sql1, 1)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", d)

		var e Tc
		err = s.QuerySql(&e, sql1, 1)
		if err != nil {
			if err == sqlp.ErrTooManyResults {
				fmt.Printf("%s\n", err.Error())
			} else {
				return err
			}
		}
		fmt.Printf("%v\n", e)

		var f map[string]interface{}
		err = s.QuerySql(&f, sql1, 1)
		if err != nil {
			if err == sqlp.ErrTooManyResults {
				fmt.Printf("%s\n", err.Error())
			} else {
				return err
			}
		}
		fmt.Printf("%v\n", f)

		var g int
		err = s.QuerySql(&g, sql2, 1)
		if err != nil {
			if err == sqlp.ErrTooManyResults {
				fmt.Printf("%s\n", err.Error())
			} else {
				return err
			}
		}
		fmt.Printf("%v\n", g)

		var h sql.NullString
		err = s.QuerySql(&h, sql2, 1)
		if err != nil {
			if err == sqlp.ErrTooManyResults {
				fmt.Printf("%s\n", err.Error())
			} else {
				return err
			}
		}
		fmt.Printf("%v\n", h)

		return nil
	})
}
