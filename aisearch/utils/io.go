package utils

// Package utils provides utility functions for the AI search system.
// This file contains I/O related utilities.

import (
	"encoding/json"
	"io"
)

// ReadToEndJson reads JSON data from an io.Reader into a generic type.
// It handles nil readers and returns the decoded data or error.
//
// Type Parameters:
//   T: The target type to decode into
//
// Parameters:
//   r: io.Reader containing JSON data (can be nil)
//
// Returns:
//   T: Decoded data of type T
//   error: Decoding error or io.EOF if reader is nil
//
// Example:
//   type User struct { Name string }
//   var user User
//   user, err := ReadToEndJson[User](reader)
func ReadToEndJson[T any](r io.Reader) (T, error) {
	var t T
	if r == nil {
		return t, io.EOF
	}
	err := json.NewDecoder(r).Decode(&t)
	return t, err
}
