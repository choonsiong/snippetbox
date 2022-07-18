package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"github.com/choonsiong/snippetbox/pkg/models/mysql"
	"github.com/golangcollege/sessions"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// application struct hold the application-wide dependencies for the web application.
// Note that we define the snippets and users as interfaces because we
// need to use them for the mock tests.
type application struct {
	debug         bool
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      mysql.SnippetInterface
	templateCache map[string]*template.Template
	users         mysql.UserInterface
}

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

func main() {
	// Command-line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:password@/snippetbox?parseTime=true", "MySQL data source name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key") // 32 bytes random key to encrypt and authenticate session cookies
	debug := flag.Bool("debug", false, "Enable debug mode")

	flag.Parse()

	// Custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	// Create database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Close the connection pool
	defer db.Close()

	// Initialize a template cache
	//templateCache, err := newTemplateCache("./ui/html/")
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize a new session manager
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	// Initialize a new instance of application containing the dependencies
	app := &application{
		debug:         *debug,
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		users:         &mysql.UserModel{DB: db},
	}

	// Initialize a tls.Config struct to hold the non-default TLS settings
	// we want the server to use.
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Initialize a new http.Server struct (so that we can use our own custom logger)
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	//err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

// openDB wraps sql.Open() and returns a sql.DB connection pool for a given DSN
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
