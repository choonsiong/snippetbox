package main

/*
	- Parsing the runtime configuration settings for the application
	- Establishing the dependencies for the handlers
	- Running the HTTP server
*/

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages, use different flags to show file name
	// and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		config: new(config),
		errorLog: errorLog,
		infoLog: infoLog,
	}

	flag.StringVar(&app.config.addr, "addr", ":4000", "HTTP network address:port")
	flag.StringVar(&app.config.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// First check whether the application is running on Heroku, if yes then
	// use the port assigned by Heroku.
	port := os.Getenv("PORT")
	var httpListenAddr = ""

	if port != "" {
		httpListenAddr = ":" + port
	} else {
		infoLog.Printf("Environment variable $PORT is not defined, set http listening address to %v", cfg.addr)
		httpListenAddr = app.config.addr
	}

	// Initialize a new http.Server struct. Use the custom errorLog logger
	// in the event of any problems.
	httpServ := &http.Server{
		Addr:     httpListenAddr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	//log.Printf("Starting server on %v", httpListenAddr)
	infoLog.Printf("Starting server on %v", httpListenAddr)

	// Note that any error returned by http.ListenAndServe() is always non-nil.
	// :port will listen on all available network interfaces
	//err := http.ListenAndServe(httpListenAddr, mux)
	err := httpServ.ListenAndServe()
	//log.Fatal(err)
	errorLog.Fatal(err)
}
