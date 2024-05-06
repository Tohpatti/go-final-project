package main

import (
	"database/sql"
	"fmt"
	"go-final-project/api/routes"
	"go-final-project/databases"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	// Load .env file
	err = godotenv.Load("configs/.env")
	if err != nil {
		panic(err)
	}

	// Connect to database
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	// Check if connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Run migrations
	databases.DbMigrations(db)
	// Close database connection
	defer db.Close()

	fmt.Println("Connected to database")

	routes.StartServer().Run(":3000")
}
