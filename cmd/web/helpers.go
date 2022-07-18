package main

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
	"runtime/debug"
	"time"
)

// serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	// If debug mode is on, display the debug info to the user.
	if app.debug {
		http.Error(w, trace, http.StatusInternalServerError)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code and corresponding description to
// the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// isAuthenticated return true if the current request is from authenticated user,
// otherwise return false.
func (app *application) isAuthenticated(r *http.Request) bool {
	//return app.session.Exists(r, "authenticatedUserID")
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}

// notFound is simply a convenience wrapper around clientError which sends a
// 404 Not Found response to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w,fmt.Errorf("the template %s does not exist", name))
		return
	}

	// Below implements the template render in a two-stage process so that we won't send an
	// half-complete page to the user when there's a runtime error.

	// #1
	// Initialize a new buffer for the 'trial' render
	buf := new(bytes.Buffer)

	// Write the template to a buffer, instead of straight to the
	// http.ResponseWriter first. If there's an error, call serverError
	// and then return.
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	// #2
	// Now we can write the contents of the buffer to the
	// http.ResponseWriter.
	buf.WriteTo(w)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	// Injects current year into template data
	td.CSRFToken = nosurf.Token(r)
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	td.IsAuthenticated = app.isAuthenticated(r)
	return td
}