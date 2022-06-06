package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73?param1=value&database=techcoach")

	if err != nil {
		panic(err)
	}

	// code style Guard Pattern

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	query := "select * from cover"
	// Exec => insert update delete
	// Query => select mutiple row2
	// QueryRow => select a row
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	id := 0
	name := ""
	ok := rows.Next()

	if ok {
		rows.Scan(&id, &name)
	}

	fmt.Println(id, name)
}
