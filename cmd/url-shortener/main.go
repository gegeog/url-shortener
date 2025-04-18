package main

import (
	"net/http"
	"os"

	"github.com/gegeog/url-shortener/internal/http-server/handlers/redirect"
	"github.com/gegeog/url-shortener/internal/storage/memstorage"
)

func main() {
	s, err := memstorage.New()

	if err != nil {
		os.Exit(1)
	}

	mux := http.NewServeMux()
	// mux.HandleFunc(`POST /`, saveHandler.New(s))
	mux.HandleFunc(`GET /{id}`, redirect.New(s))

	if err := http.ListenAndServe(`:8080`, mux); err != nil {
		panic(err)
	}

}
