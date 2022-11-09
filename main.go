package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Welcome !!!"))
	})
	http.ListenAndServe(":80", r)
}
