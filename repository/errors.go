package repository

import "errors"

var ErrForbidden = errors.New("forbidden")
var ErrNotFound = errors.New("Resource not found")
