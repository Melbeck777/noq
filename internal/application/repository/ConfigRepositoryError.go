package repository

import "errors"

var (
	ErrConfigNotFound = errors.New("config not found")
	ErrConfigParse    = errors.New("config parse error")
	ErrConfigInvalid  = errors.New("config invalid")
)
