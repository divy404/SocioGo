package main

import (
	"SocioGo/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	var payload CreateCommentPayload

	// Parse JSON payload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRquestResponse(w, r, err)
		return
	}

	// Validate payload
	if err := Validate.Struct(payload); err != nil {
		app.badRquestResponse(w, r, err)
		return
	}

	// Create comment
	comment := &store.Comment{
		PostID:  id,
		UserID:  121, // replace with auth user ID in future
		Content: payload.Content,
	}

	ctx := r.Context()
	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Optionally attach user info
	comment.User = store.User{
		ID:       comment.UserID,
		Username: "anonymous", // replace with real username if needed
	}

	// Return the comment
	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
