package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new mux to hadle incming HTTP request
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //define handler function for the root url
		fmt.Fprintf(w, "Hey, this is Ray!")//send a response with the message as shown
	})

	// Apply the logging middleware- this will log each request before passing it on to the actual handler
	loggedMux := loggingMiddleware(mux)

	// Apply the authentication middleware- this will check for the authentication header before allowing access
	authedMux := authMiddleware(loggedMux)

	fmt.Println("Starting server on :8080...") //starting the server
	http.ListenAndServe(":8080", authedMux)//the server listens for the incoming HTTP request and passes them through the middlware chain
}
