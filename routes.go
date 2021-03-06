package main

// This file is responsible for setting up the REST controller and the handlers

import (
	"net/http"

	"github.com/sevren/test/middlewares"
	"github.com/sevren/test/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/go-chi/render"
)

type EmptyResponse struct{}

func Routes(store storage.ItemStore, challenge3features bool) (*chi.Mux, error) {

	// CORS configuration
	corsConf := corsConfig()

	r := chi.NewRouter()

	// Set up the middlewares which each request will pass through
	r.Use(corsConf.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// If rabbitmq is connected then we can use challenge 3 stuff
	// this sets a middleware which adds the challenge 3 context to the request.
	if challenge3features {
		r.Use(middlewares.Challenge3(challenge3features))
	}

	// Sets up the REST controller
	// endpoints are /{user} and /{user}/licenses
	r.Route("/{user}", func(r chi.Router) {
		r.Post("/", handleUser)
		r.Route("/licenses", func(r chi.Router) {
			r.Use(store.AuthUser)
			r.Post("/", store.HandleLicenses)
		})
	})
	r.Route("/usedlicenses", func(r chi.Router) {
		r.Get("/", store.HandleUsedLicenses)
	})

	return r, nil
}

// handles the use case for a simple /{user} rest call
func handleUser(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, EmptyResponse{})
}

func corsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token", "Cache-Control", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
