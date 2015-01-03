package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/squiidz/fishMe/module/utility"
)

func SetupDB() *sql.DB {
	dbConfig, err := utility.LoadPage("article/dbconfig")
	utility.ShitAppend(err)
	DB, err := sql.Open("postgres", string(dbConfig.Body))
	utility.ShitAppend(err)
	return DB
}
