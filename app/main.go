package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ruziba3vich/todoGo/internal/handlers"
)

func main() {
	router := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "todo")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		PrintError(err)
	}
	defer db.Close()

	databases := []string{"users", "tasks"}

	for _, dbName := range databases {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		if err != nil {
			PrintError(err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			PrintError(err)
		}
	}

	router.POST("users/register", func(c *gin.Context) {
		handlers.Register(c, db)
	})

	router.POST("users/authorize", func(c *gin.Context) {
		handlers.Authorize(c, db)
	})

	router.POST("users/create-task", func(c *gin.Context) {
		handlers.CreateTask(c, db)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		PrintError(err)
	}
}

func PrintError(err error) {
	log.Fatal("Error :", err)
}
