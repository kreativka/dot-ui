package desktop

import "errors"

var (
	ErrHiddenEntry  = errors.New("hidden entry")
	ErrInvalidEntry = errors.New("invalid entry")
)
