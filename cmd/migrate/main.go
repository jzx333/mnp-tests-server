package main

import (
	"log"

	"mnp-tests-server/internal/db"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := db.LoadConfig()
	log.Printf("migrating using DSN: %s", cfg.RedactedDSN())

	m, err := migrate.New(
		"file://migrations",
		cfg.DSN(),
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	// Просто запускаем Up; если всё уже применено, будет ErrNoChange — это ок
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("migration up failed: %v", err)
	}

	log.Println("migrations applied successfully")
}
