package main

import "github.com/choonsiong/snippetbox/pkg/models"

// Define a template data type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Snippet *models.Snippet
}
