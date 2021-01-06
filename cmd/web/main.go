package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// For heroku deployment
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT is not defined, set to 4000")
		port = "4000"
	}

	// Initialize a new servemux
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir() function is relative to the
	// project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on :%v", port)

	// If http.ListenAndServe() returns an error we use the log.Fatal() function
	// to log the error message and exit. Note that any error returned by
	// http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":" + port, mux) // :port will listen on all available network interfaces
	log.Fatal(err)
}
