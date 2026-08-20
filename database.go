package belajar_golang_database

import (
	// "context"
	"database/sql"
	// "fmt"
	// "testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	dsn := "root:@tcp(localhost:3306)/belajar_golang_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db,nil
}
