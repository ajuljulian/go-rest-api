package main

import (
	"fmt"

	"github.com/ajuljulian/go-rest-api/internal/comment"
	"github.com/ajuljulian/go-rest-api/internal/db"
	transportHttp "github.com/ajuljulian/go-rest-api/internal/transport/http"
)

// Run - is going to be responsible for the instantiation
// and startup of our go application
func Run() error {
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
