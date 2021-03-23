package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sebmattmuller/booking_app/pkg/config"
	"github.com/sebmattmuller/booking_app/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//Pat Router
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Static Files
	fileServer := http.FileServer(http.Dir("./static-files/"))
	mux.Handle("/static-files/*", http.StripPrefix("/static-files", fileServer))

	return mux
}
