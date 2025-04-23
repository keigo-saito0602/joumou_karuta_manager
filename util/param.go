package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
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

// ConvertIntToEnum は DB から int64 を読み取って Enum 型に変換する
func ConvertIntToEnum(value interface{}) (int64, error) {
	i, ok := value.(int64)
	if !ok {
		return 0, fmt.Errorf("failed to scan enum: value is not int64: %v", value)
	}
	return i, nil
}

// ConvertEnumToInt は Enum 型を DB に保存する
func ConvertEnumToInt(val int) (driver.Value, error) {
	return int64(val), nil
}
