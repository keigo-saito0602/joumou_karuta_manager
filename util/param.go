package util

import (
	"errors"
	"strconv"
)

// ParseUint64Param converts a string to uint64 safely
func ParseUint64Param(s string) (uint64, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}
	return id, nil
}
