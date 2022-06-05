package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@127.0.0.1?param1=value&database=test")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
