package models

import (
	"errors"
)

var (
	// Snippet model
	ErrNoRecord = errors.New("models: no matching record found")
	// User model
	ErrInvalidCredentails = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)
