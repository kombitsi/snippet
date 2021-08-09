package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)
func (app *application) serverError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) ClientError(w http.ResponseWriter, status int)  {
	http.Error(w, http.StatusText(status), status)

}

func (app *application) NFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}