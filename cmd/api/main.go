package main

import (
	"SocioGo/internal/db"
	"SocioGo/internal/env"
	"SocioGo/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR"),
		db : dbConfig{
			addr: env.GetStringEnv("DB_ADDR") ,
			maxOpenConns: env.GetIntEnv("DB_MAX_OPEN_CONNS"),
			maxIdleConns: env.GetIntEnv("DB_MAX_IDLE_CONNS"),
			maxIdleTime: env.GetStringEnv("DB_MAX_IDLE_TIME"),
		},
	}
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store: store,
	}
	
	mux := app.mount()
	log.Fatal(app.run(mux))
}
