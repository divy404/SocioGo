package main

import (
	"SocioGo/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetStringEnv("ADDR"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
