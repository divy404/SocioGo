package main

import "net/http"

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email   string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	
}