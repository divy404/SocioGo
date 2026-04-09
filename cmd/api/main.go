package main

import (
	"SocioGo/internal/db"
	"SocioGo/internal/env"
	"SocioGo/internal/store"
	"time"

	"go.uber.org/zap"
)
const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR"),
		db : dbConfig{
			addr: env.GetStringEnv("DB_ADDR") ,
			maxOpenConns: env.GetIntEnv("DB_MAX_OPEN_CONNS"),
			maxIdleConns: env.GetIntEnv("DB_MAX_IDLE_CONNS"),
			maxIdleTime: env.GetStringEnv("DB_MAX_IDLE_TIME"),
		},
		env: env.GetStringEnv("ENV"),
		mail : mailConfig{
			exp: time.Hour * 24 * 3,
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store: store,
		logger: logger,
	}
	
	mux := app.mount()
	logger.Fatal(app.run(mux))
}
