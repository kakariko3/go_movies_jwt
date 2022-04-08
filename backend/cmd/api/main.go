package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"app/models"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config

	// フラグの設定 (コマンドのオプション)
	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	// DSN = "postgres://username:password@hostname:port/databasename?option..."
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable", "Postgres connection string")
	// flag.StringVar(&cfg.db.dsn, "dsn", "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", "Postgres connection string")
	flag.Parse()

	// ロギングの設定
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// データベースに接続
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatalln(err)
	}
	defer db.Close()

	app := application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// データベースの接続処理を定義
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
