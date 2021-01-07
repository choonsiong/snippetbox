package main

import "log"

// Define an application struct to hold the application-wide
// dependencies for the web application.
type Application struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
}
