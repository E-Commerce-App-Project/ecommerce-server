package commons

import (
	"errors"
)

// ErrDBConn error type for Error DB Connection
var ErrDBConn = errors.New("ErrDBConn")

// ErrCacheConn error type for Error Cache Connection
var ErrCacheConn = errors.New("ErrCacheConn")

var MappingError = errors.New("MappingError")
