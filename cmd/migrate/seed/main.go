package main

import (
	"SocioGo/internal/db"
	"SocioGo/internal/env"
	"SocioGo/internal/store"
	"log"
)

func main() {
	addr := env.GetStringEnv("DB_ADDR")
	conn, err := db.New(addr,3,3,"15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)

	db.Seed(store)
}
