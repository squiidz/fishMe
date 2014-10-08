package database

import (
	"github.com/PushKids/module/utility"
	"database/sql"
	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	dbConfig, err := utility.LoadPage("article/dbconfig")
	utility.ShitAppend(err)
	DB, err := sql.Open("postgres", string(dbConfig.Body))
	utility.ShitAppend(err)
	return DB
}
