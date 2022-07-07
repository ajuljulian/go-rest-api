package main

import (
	"context"
	"fmt"

	"github.com/ajuljulian/go-rest-api/internal/comment"
	"github.com/ajuljulian/go-rest-api/internal/db"
)

// Run - is going to be responsible for the instantiation
// and startup of our go application
func Run() error {
	fmt.Println("starting up our applciation")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "b7ffc347-67a3-426d-b2eb-7b4b92e92f20",
			Slug:   "manual-test",
			Author: "Elliot",
			Body:   "Hello World",
		},
	)

	fmt.Println(cmtService.GetComment(context.Background(), "b7ffc347-67a3-426d-b2eb-7b4b92e92f20"))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
