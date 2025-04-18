package save

import "net/http"

type URLSaver interface {
	SaveURL(urlToSave string) (string, error)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func New(urlsaver URLSaver) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
