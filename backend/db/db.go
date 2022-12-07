package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func Init(dbString string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dbString)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
