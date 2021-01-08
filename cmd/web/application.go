package main

import (
	"github.com/choonsiong/snippetbox/pkg/models/mysql"
	"log"
)

// Define an application struct to hold the application-wide
// dependencies for the web application.
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetModel
}
