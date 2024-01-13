package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, "asd") {
		fmt.Fprintf(w, "gotcha")
		return
	}

	fmt.Fprintf(w, "came in here")
}

func ApiHandler(router *http.ServeMux) {
	router.Handle("/api/", apiHandler{})
}

func WebHandler(router *http.ServeMux) {
	router.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("pages/dashboard.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func StaticHandler(router *http.ServeMux) {
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("js/"))))
}
