package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}


func TextOpenConnection(t *testing.T) {
	dsn := "root:@tcp(localhost:3306)/belajar_golang_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Panic(err)
	}
	defer db.Close()
}