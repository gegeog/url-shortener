package memstorage

import (
	"fmt"

	"github.com/gegeog/url-shortener/internal/lib/random"
	"github.com/gegeog/url-shortener/internal/storage"
)

type MemStorage struct {
	storage map[string]string
}

const size = 5

func New() (*MemStorage, error) {
	ms := &MemStorage{
		storage: make(map[string]string),
	}

	return ms, nil
}

func (ms *MemStorage) SaveURL(urlToSave string) (string, error) {
	op := "storage.memstorage.SaveURL"

	if _, exists := ms.storage[urlToSave]; exists {
		return "", fmt.Errorf("%s: %w", op, storage.ErrorURLAlreadyExist)
	}

	shortened := random.NewPseudoRandomString(size)
	ms.storage[shortened] = urlToSave
	return shortened, nil
}

func (ms *MemStorage) GetURL(shortURL string) (string, error) {
	op := "storage.memstorage.GetURL"
	longURL, exists := ms.storage[shortURL]

	if !exists {
		return "", fmt.Errorf("%s: %w", op, storage.ErrorURLNotFound)
	}

	return longURL, nil
}
