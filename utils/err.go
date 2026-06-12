package utils

import "errors"

var ErrForbidden = errors.New("forbidden")
var ErrNotFound = errors.New("resource not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrEmailExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInternalServer = errors.New("internal server error")
var ErrInvalidAmount = errors.New("amount must be greater than zero")
var ErrDateInFuture = errors.New("date cannot be in the future")
var ErrDatabase = errors.New("database error")
