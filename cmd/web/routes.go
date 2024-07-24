package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle(http.MethodGet+" /static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc(http.MethodGet+" /", app.home)
	mux.HandleFunc(http.MethodGet+" /snippet/view/{id}", app.snippetView)
	mux.HandleFunc(http.MethodGet+" /snippet/create", app.snippetCreate)
	mux.HandleFunc(http.MethodPost+" /snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(mux)
}
