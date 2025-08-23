package main

import (
	"SocioGo/internal/env"
	"SocioGo/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store: store,
	}
	
	mux := app.mount()
	log.Fatal(app.run(mux))
}
