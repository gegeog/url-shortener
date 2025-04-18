package storage

import "errors"

var (
	ErrorURLAlreadyExist = errors.New("URL already exist")
	ErrorURLNotFound     = errors.New("URL not found")
)
