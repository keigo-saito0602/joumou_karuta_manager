// domain/errors.go
package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrInvalidArgument     = errors.New("invalid argument")
	ErrNotFound            = errors.New("data not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrDataResourceAccess  = errors.New("failed to access data resource")
	ErrInternalServer      = errors.New("internal server error")
	ErrPermissionDenied    = errors.New("permission denied")
	ErrUnimplemented       = errors.New("unimplemented")
	ErrDuplicateEntry      = fmt.Errorf("%w: duplicate entry", ErrInvalidArgument)
	ErrUserNotFound        = fmt.Errorf("%w: user not found", ErrNotFound)
	ErrUserAlreadyExists   = fmt.Errorf("%w: user already exists", ErrAlreadyExists)
	ErrInvalidUserName     = fmt.Errorf("%w: invalid user name", ErrInvalidArgument)
	ErrInvalidAgeParameter = fmt.Errorf("%w: age must be >= 0", ErrInvalidArgument)
)

func ErrorHandler(err error) error {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		switch mysqlErr.Number {
		case 1062:
			return ErrDuplicateEntry
		case 1643:
			return ErrNotFound
		}
	}
	return ErrDataResourceAccess
}

func ErrorToHTTPStatus(err error) int {
	switch {
	case errors.Is(err, ErrInvalidArgument):
		return http.StatusBadRequest
	case errors.Is(err, context.DeadlineExceeded):
		return http.StatusRequestTimeout
	case errors.Is(err, ErrNotFound),
		errors.Is(err, gorm.ErrRecordNotFound),
		errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound
	case errors.Is(err, ErrAlreadyExists):
		return http.StatusConflict
	case errors.Is(err, ErrPermissionDenied):
		return http.StatusForbidden
	case errors.Is(err, ErrDataResourceAccess):
		return http.StatusInternalServerError
	case errors.Is(err, ErrUnimplemented):
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}
}