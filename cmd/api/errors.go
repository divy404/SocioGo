package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusInternalServerError,"the server encountered a problem")
}

func (app *application) badRquestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	app.logger.Warnf("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusBadRequest,err.Error())
}
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound,"not found")
}
func (app *application) conflictError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("Conflict Error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Errorw("Conflict Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}