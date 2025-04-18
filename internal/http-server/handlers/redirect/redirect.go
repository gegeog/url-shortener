package redirect

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gegeog/url-shortener/internal/storage"
)

type URLGetter interface {
	GetURL(shortURL string) (string, error)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func New(urlGetter URLGetter) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.redirect.New"
		id := r.PathValue("id")
		longURL, err := urlGetter.GetURL(id)

		if errors.Is(err, storage.ErrorURLNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s: %s", op, err)))
			return
		}

		w.Header().Set("Location", longURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
