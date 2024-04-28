package databases

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConn *sql.DB
)

//go:embed models/*.sql
var dbMigrations embed.FS

func DbMigrations(db *sql.DB) {
	// Run migrations
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "models",
	}

	// Run migrations
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error executing migration")
		panic(err)
	}

	DbConn = db
	fmt.Printf("Applied %d migrations!\n", n)
}
