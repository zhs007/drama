package dramautils

import "errors"

var (
	// ErrInvalidArrayLength - invalid array length
	ErrInvalidArrayLength = errors.New("invalid array length")
	// ErrInvalidJsonString - invalid json string
	ErrInvalidJsonString = errors.New("invalid json string")
	// ErrInvalidJsonInt - invalid json int
	ErrInvalidJsonInt = errors.New("invalid json int")
)
