Sqlp
======

[![Go Reference](https://pkg.go.dev/badge/github.com/valord577/sqlp.svg)](https://pkg.go.dev/github.com/valord577/sqlp)
![GitHub](https://img.shields.io/github/license/valord577/sqlp?t=0)

Sqlp is a library which provides several features for `database/sql` and only depend on the stdlib `database/sql`.

Compatibility
------

- Compatible with go 1.16+

Features
------

- Use raw sql via custom handlers when needed
- Make object-relational mapping easier to use

Installing
------

go mod:

```shell
go get github.com/valord577/sqlp
```

Example
------

<details>
<summary>
- Do object-relational mapping
</summary>

```go
package main

import (
    "database/sql"
    "fmt"

    "github.com/valord577/sqlp"
)

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

func main() {
    do(func(s *sqlp.DBSession) error {

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

func do(f func(s *sqlp.DBSession) error) {
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
```
</details>

Changes
------

See the [CHANGES](CHANGE.md) for changes.

License
------

See the [LICENSE](LICENSE) for Rights and Limitations (MIT).
