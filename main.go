package main

import (
	asciiZ "ascii/art"
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	// fs := http.FileServer(http.Dir("web"))
	// http.Handle("/",fs)

	mux := http.NewServeMux()

	// Define handler functions
	mux.HandleFunc("/", nab)
	mux.HandleFunc("/ascii-art", zai)
	cssHandler := http.FileServer(http.Dir("web/"))
	mux.Handle("/web/", http.StripPrefix("/web/", cssHandler))

	// Create a new HTTP server instance
	http.ListenAndServe(":2004", mux)

	// Start the server
	fmt.Println("Server listening on :2004")

}

func nab(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Wrong path WLLLAAAKKK!!!", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "wrong method WLLLAAAKK!", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

type ArtTemplate struct {
	Madan string // فخر العرب
}

func zai(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	ban := r.FormValue("banne")
	artText := asciiZ.AsciiART(text, ban)

	finalart := ArtTemplate{Madan: artText}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, finalart)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
