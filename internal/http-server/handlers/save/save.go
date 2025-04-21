package save

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gegeog/url-shortener/internal/storage"
)

type URLSaver interface {
	SaveURL(urlToSave string) (string, error)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func New(urlsaver URLSaver) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.save.New"
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s: %s", op, "failed to read request body")))
			return
		}

		short, err := urlsaver.SaveURL(string(body))
		if errors.Is(err, storage.ErrorURLAlreadyExist) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s: %s", op, err)))
			return
		}

		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", short)))
	}
}
