package main

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	if r.URL.Path != "/" {
		http.Error(w, "RequestNotFound", http.StatusNotFound)
	}

	ts, _ := template.ParseFiles("index.html")
	w.WriteHeader(http.StatusOK)
	ts.Execute(w, nil)

}

func ReportSheetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	if r.URL.Path != "/scoresheet" {
		http.Error(w, "RequestNotFound", http.StatusNotFound)
	}
	r.ParseForm()
	score := r.FormValue("score")

	result := ReportSheet(score)

	ts, _ := template.ParseFiles("index.html")

	ts.Execute(w, result)

}
