package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentRepositoryInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	
	ctx := context.Background()
	comment := entity.Comment{
		Email: "aryo@gmail.com",
		Comment: "test comment",
	}
	result,err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRepositoryFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 35)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRepositoryFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	comments , err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _,comment := range comments {
		fmt.Println(comment)
	}
}