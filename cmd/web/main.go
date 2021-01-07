package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := new(config)

	// Create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages, use different flags to show file name
	// and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing
	// the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address:port")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// First check whether the application is running on Heroku, if yes then
	// use the port assigned by Heroku.
	port := os.Getenv("PORT")
	var httpListenAddr = ""

	if port != "" {
		httpListenAddr = ":" + port
	} else {
		infoLog.Printf("Environment variable $PORT is not defined, set http listening address to %v", cfg.addr)
		httpListenAddr = cfg.addr
	}

	// Initialize a new servemux
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir() function is relative to the
	// project directory root.
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Initialize a new http.Server struct. Use the custom errorLog logger
	// in the event of any problems.
	httpServ := &http.Server{
		Addr:     httpListenAddr,
		ErrorLog: errorLog,
		Handler:  mux,
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
