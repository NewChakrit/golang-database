package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

//  ------ Code style Guard Pattern ------
type Cover struct {
	Id   int
	Name string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73?param1=value&database=techcoach")
	// db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73)/techcoach")

	if err != nil {
		panic(err)
	}

	covers, err := GetCovers()
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, cover := range covers {
		fmt.Println(cover)
	}

}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id,name from cover" //input colunn name!!
	// Exec => insert update delete
	// Query => select mutiple row2
	// QueryRow => select a row
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, cover)
	}

	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, name from cover where id=@id" //@id ตั้งชื่อให้ตรงกับ column นั้นๆ
}
