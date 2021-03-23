package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// Middleware example to write to page
//func WriteToConsole(next http.Handler) http.Handler  {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//		fmt.Println("Hit the page")
//		next.ServeHTTP(w, r)
//	})
//}

// Adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler  {
	return session.LoadAndSave(next)
}