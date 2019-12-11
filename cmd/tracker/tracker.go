package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/havr/time-tracker/internal/api"
	"github.com/havr/time-tracker/internal/stores"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	DatabaseURL            string `envconfig:"database_url"`
	ServeAt                string `envconfig:"serve_at"`
	StaticDir              string `envconfig:"static_dir"`
	MigrateFrom            string `envconfig:"migrate_from"`
	WaitForDatabaseSeconds int    `envconfig:"wait_for_database_seconds"`
}

func main() {
	var config Config
	if err := envconfig.Process("TT", &config); err != nil {
		log.Fatal("processing logs: ", err.Error())
	}

	db, err := dialDatabase(config.DatabaseURL, time.Duration(config.WaitForDatabaseSeconds)*time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	if config.MigrateFrom != "" {
		if err := applyMigrations(db, config.MigrateFrom); err != nil {
			log.Fatal(err.Error())
		}
	}

	sessionStore := stores.NewDatabaseSessionStore(db)
	handler := api.NewRouter(config.StaticDir, sessionStore)
	server := http.Server{
		Addr:    config.ServeAt,
		Handler: handler,
	}

	fmt.Printf("Serving on %s\n", config.ServeAt)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

func applyMigrations(db *sql.DB, src string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("creating postgres migrator driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+src,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("creating migrator instance: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("applying migrations: %s", err)
	}

	return nil
}

func dialDatabase(url string, waitInterval time.Duration) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("open database: %s", err.Error())
	}

	deadline := time.Now().Add(time.Duration(waitInterval) * time.Second)
	for time.Now().Before(deadline) {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("ping database: %s", err.Error())
	}

	return db, nil
}
