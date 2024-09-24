package errors

import "errors"

// New returns an error that formats as the given text.
var New = errors.New

// known errors
var (
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrExpiredAccessToken = errors.New("expired access token")
)
