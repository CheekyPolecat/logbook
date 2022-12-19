package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	static := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", static))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/student/add", app.addChild)
	mux.HandleFunc("/student/view", app.viewChild)

	return mux
}
