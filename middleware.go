package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler { //loggin middlware to log each incoming request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path) //log the http method and the url path of the incmoing request
		next.ServeHTTP(w, r)                                          //call the next handler in the chain
	})
}

func authMiddleware(next http.Handler) http.Handler { //Authenticatio middleware to enforce access control- checks the presence of the X-Auth token header and verifies it's values
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != "secret-token" {
			http.Error(w, "Forbidden", http.StatusForbidden) //if the token is incorrect or missing, send a Forbidden response
			return
		}
		next.ServeHTTP(w, r) //if the token is correct, call the next handler in the chain
	})
}
