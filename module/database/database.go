package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
)

func ShitAppend(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func loadPage(title string) []byte {
	filename := title + ".pk"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		ShitAppend(err)
	}
	return body
}

func SetupDB() *sql.DB {
	dbConfig := loadPage("article/dbconfig")
	DB, err := sql.Open("postgres", string(dbConfig))
	ShitAppend(err)
	return DB
}
