package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email:= "budi@gmail.com"
	password := "budi"

	query := "SELECT email from user where email =  ? and password = ?"
	
	ctx := context.Background()
	rows, err := db.QueryContext(ctx , query, email,password)

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
		fmt.Println("gagal login ", email)
	}
}

func TestSqlInsertAdmin(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	email := "budi@gmail.com"
	password := "budi"

	query := "INSERT INTO user(email, password) VALUES(?, ?)"
	
	ctx := context.Background()
	_, err := db.ExecContext(ctx , query, email, password)

	if err != nil {
		panic(err)
	}
	fmt.Println("Data user berhasil dimasukan")
	
}

func TestSqlInsertComment(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	email := "budi@gmail.com"
	comment := "sangat membantu sekali tutorialnya, terimakasih banyak"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	
	ctx := context.Background()
	results, err := db.ExecContext(ctx , query, email, comment)

	if err != nil {
		panic(err)
	}
	resultId,err := results.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Data comment berhasil dimasukan")
	fmt.Println("Last Insert ID:", resultId)
}

func TestSqlPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	
	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx , query)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "budi" + strconv.Itoa(i) + "@gmail.com";
		comment := "sangat membantu sekali tutorialnya, terimakasih banyak" + strconv.Itoa(i);

		results, err := stmt.ExecContext(ctx , email, comment)
		if err != nil {
			panic(err)
		}
		resultId,err := results.LastInsertId()
		if err != nil {
			panic(err)
		}	
		fmt.Println("Data comment berhasil dimasukan")
		fmt.Println("Comment ID:", resultId)
	}
}

func TestSqlTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}	
	
for i := 0; i < 10; i++ {
		email := "budi" + strconv.Itoa(i) + "@gmail.com";
		comment := "sangat membantu sekali tutorialnya, terimakasih banyak" + strconv.Itoa(i);

		results, err := tx.ExecContext(ctx , query, email, comment)
		if err != nil {
			panic(err)
		}
		resultId,err := results.LastInsertId()
		if err != nil {
			panic(err)
		}	
		fmt.Println("Data comment berhasil dimasukan")
		fmt.Println("Comment ID:", resultId)
	}


	err = tx.Rollback()
	if err != nil {{
		tx.Rollback()
		panic(err)
	}
 }
}