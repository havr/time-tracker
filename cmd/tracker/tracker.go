package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/havr/time-tracker/internal/api"
	"github.com/havr/time-tracker/internal/stores"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	DatabaseURL string `envconfig:"database_url"`
	ServeAt     string `envconfig:"serve_at"`
	StaticDir   string `envconfig:"static_dir"`
	MigrateFrom string `envconfig:"migrate_from"`
	WaitForDatabaseSeconds int `envconfig:"wait_for_database_seconds"`
}

func main() {
	var config Config
	if err := envconfig.Process("TT", &config); err != nil {
		log.Fatal("processing logs: ", err.Error())
	}

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal("open database: ", err.Error())
	}

	deadline := time.Now().Add(time.Duration(config.WaitForDatabaseSeconds) * time.Second)
	for time.Now().Before(deadline){
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		log.Fatal("ping database: ", err.Error())
	}

	if config.MigrateFrom != "" {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatal("creating postgres migrator driver: ", err.Error())
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file:///" + config.MigrateFrom,
			"postgres", driver)
		if err != nil {
			log.Fatal("creating migrator instance: ", err.Error())
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("applying migrations: ", err.Error())
		}
	}

	defer db.Close()

	sessionStore := stores.NewDatabaseSessionStore(db)
	handler := api.NewRouter(config.StaticDir, sessionStore)
	server := http.Server{
		Addr:    config.ServeAt,
		Handler: handler,
	}

	fmt.Printf("serving on %s\n", config.ServeAt)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
