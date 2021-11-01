package helper

import (
	"errors"
	"net/http"

	"github.com/arfan21/golang-mygram/exception"
	"gorm.io/gorm"
)

func GetStatusCode(err error) int {
	if isValidationError(err) {
		return http.StatusBadRequest
	}

	if isValidationAuthentication(err) {
		return http.StatusUnauthorized
	}

	if isValidationAuthorization(err) {
		return http.StatusForbidden
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, gorm.ErrMissingWhereClause) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func isValidationError(err error) bool {
	_, ok := err.(exception.ErrorValidation)
	return ok
}

func isValidationAuthentication(err error) bool {
	_, ok := err.(exception.ErrorAuthentication)
	return ok
}

func isValidationAuthorization(err error) bool {
	_, ok := err.(exception.ErrorAuthorization)
	return ok
}
