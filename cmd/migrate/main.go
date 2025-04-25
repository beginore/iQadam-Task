// cmd/migrate/main.go
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: migrate [up|down]")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // ← Добавьте file://
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Migration initialization failed:", err)
	}

	switch os.Args[1] {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		log.Fatal("Invalid command. Use 'up' or 'down'")
	}

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("No migrations to apply")
			return
		}
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migrations applied successfully!")
}
