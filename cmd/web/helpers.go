package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"
	"runtime/debug"
	_ "runtime/debug"
	_ "strconv"
)

func (app *application) serverError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int)  {
	http.Error(w, http.StatusText(status), status)

}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}