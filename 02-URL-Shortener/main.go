package main

import (
	"fmt"
	"net/http"

	"github.com/supreeth7/gophercises/02-URL-Shortener/handler"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/quiz":    "https://github.com/supreeth7/gophercises/tree/master/01-quiz-app",
		"/profile": "https://github.com/supreeth7",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)
	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello World")
}
