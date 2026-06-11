package utils

import "errors"

var ErrForbidden = errors.New("forbidden")
var ErrNotFound = errors.New("resource not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrEmailExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInternalServer = errors.New("internal server error")
