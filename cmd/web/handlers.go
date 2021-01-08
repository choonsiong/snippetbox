package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly  matches "/". If it
	// doesn't, use the http.NotFound() function to send a 404 response
	// to the client. We then return from the handler.
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containing the paths to the files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string {
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the files and store the templates
	// in a template set. Notice that we can pass the slice of file paths as a
	// variadic parameter.
	// If there's an error, we log the detailed error message
	// and use the http.Error() function to send a generic 500 Internal Server
	// Error response to the user.
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serveError(w, err)
		return
	}

	// Then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents
	// any dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)

	if err != nil {
		app.serveError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return an 404
	// page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	//w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with ID %d... 🤔", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note
	// that http.MethodPost is a constant equal to the string "POST".
	if r.Method != "POST" {
		// Use the Header().Set() method to add a header to the response
		// header map. The first parameter is the header name, and the second
		// parameter is the header value.
		// Note that changing the response header map after a call to w.WriteHeader()
		// or w.Write() will have no effect on the headers that the user receives.
		w.Header().Set("Allow", "POST")

		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a message response body.
		// We then return from the function so that the subsequent code is not
		// executed.
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed"))

		// Use the clientError helper.
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create dummy data.
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serveError(w, err)
		return
	}

	//w.Write([]byte("Create a new snippet..."))
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
