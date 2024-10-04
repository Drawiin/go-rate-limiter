package main

import (
	"go-rate-limiter/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	urlController := controller.NewUrlController(&map[string]string{})

	r.Post("/url", urlController.CreateUrl)
	r.Get("/{urlId}", urlController.GetUrl)

	http.ListenAndServe(":3000", r)
}
