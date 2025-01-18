package storage

import "errors"

var ErrURLExists = errors.New("URL already exists")
var ErrURLNotFound = errors.New("URL not found")
