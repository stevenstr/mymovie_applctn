package internal

import "errors"

// ErrNotFound - is returned when a requested record
// is not found.
var ErrNotFound = errors.New("record was not found")
