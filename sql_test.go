package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestSqlExec(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	query := `INSERT INTO customer(id, name) VALUES('budi', 'Budi')`
	
	ctx := context.Background()
	_, err := db.ExecContext(ctx , query)

	if err != nil {
		panic(err)
	}
	fmt.Println("Data customer berhasil dimasukan")
	
}

func TestSqlQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	query := `SELECT id,name from customer`
	
	ctx := context.Background()
	rows, err := db.QueryContext(ctx , query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id,name string
		err := rows.Scan(&id,&name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:",id)
		fmt.Println("Name:",name)
	}
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	query := `SELECT id,name,email,balance,rating,birth_date,married,created_at from customer`
	
	ctx := context.Background()
	rows, err := db.QueryContext(ctx , query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("hai")
	for rows.Next() {
		var id,name string
		var email sql.NullString
		var balance int32
		var rating float32
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err := rows.Scan(&id,&name,&email,&balance,&rating,&birth_date,&married,&created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:",id)
		fmt.Println("Name:",name)
		if email.Valid{
			fmt.Println("Email:",email.String)
		}else{
			fmt.Println("Email: null")
		}
		fmt.Println("Balance:",balance)
		fmt.Println("Rating:",rating)
		if birth_date.Valid{
			fmt.Println("Birth Date:",birth_date.Time)
		}else{
			fmt.Println("Birth Date: null")
		}
		fmt.Println("Created At:",created_at)
		fmt.Println("====================")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email:= "admin'; #"
	password := "saalah"

	query := "SELECT email from user where email =  '" + email + "' and password = '" + password + "'"
	
	ctx := context.Background()
	rows, err := db.QueryContext(ctx , query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login ", email)
	}else{
		fmt.Println("gagal login")
	}
}