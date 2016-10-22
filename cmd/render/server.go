package main

import (
	// "html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("output"))
	http.Handle("/output/", http.StripPrefix("/output/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Foo!\n"))
	})
	http.ListenAndServe(":8080", nil)
}
