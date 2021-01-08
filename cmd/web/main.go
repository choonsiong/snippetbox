package main

/*
	- Parsing the runtime configuration settings for the application
	- Establishing the dependencies for the handlers
	- Running the HTTP server
*/

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	// Command-line arguments
	flag.StringVar(&app.config.addr, "addr", ":4000", "HTTP network address:port")
	flag.StringVar(&app.config.dsn, "dsn", "admin:pass@tcp(snippetbox_mysql_1:3306)/snippetbox?parseTime=true", "MySQL data source name")
	flag.StringVar(&app.config.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// First check whether the application is running on Heroku, if yes then
	// use the port assigned by Heroku.
	port := os.Getenv("PORT")
	var httpListenAddr = ""

	if port != "" {
		httpListenAddr = ":" + port
	} else {
		infoLog.Printf("Environment variable $PORT is not defined, set http listening address to %v", app.config.addr)
		httpListenAddr = app.config.addr
	}

	// Setup database connection
	db, err := openDB(app.config.dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

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
	err = httpServ.ListenAndServe()
	//log.Fatal(err)
	errorLog.Fatal(err)
}

// The OpenDB() function wraps sql.Open() and returns a sql.DB connection
// pool for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}