package main

import (
	"database/sql"
	"errors"
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
	// db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.73?param1=value&database=techcoach")
	db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73)/techcoach")

	if err != nil {
		panic(err)
	}

	// cover := Cover{9, "New"}
	// err = AddCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	cover := Cover{9, "Chakrit"}
	err = UpdateCover(cover)
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

	// ------ MSSQL SERVER -----
	// cover, err := GetCover(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cover)

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

	// ----- MSSQL SERVER -----
	// query := "select id, name from cover where id=@id" //@id ตั้งชื่อให้ตรงกับ column นั้นๆ
	// row := db.QueryRow(query, sql.Named("id", id))

	// ----- MYSQL  -----
	query := "select id, name from cover where id=?" //@id ตั้งชื่อให้ตรงกับ column นั้นๆ
	row := db.QueryRow(query, id)

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil //ส่งตำแหน่ง cover ไป
}

func AddCover(cover Cover) error {
	query := "insert into cover (id, name) values (?,?)"
	result, err := db.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Cannot insert")
	}

	return nil
}

func UpdateCover(cover Cover) error {
	query := "update cover set name=? where id=?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Cannot update")
	}

	return nil
}
