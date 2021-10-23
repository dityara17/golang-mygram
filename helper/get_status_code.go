package helper

import (
	"net/http"

	"github.com/arfan21/golang-mygram/exception"
)

func GetStatusCode(err error) int {
	if isValidationError(err) {
		return http.StatusBadRequest
	}

	// error authentication
	// if isAuthenticationError(err) {
	// 	return http.StatusUnauthorized
	// }
	// error authorization

	return http.StatusInternalServerError
}

func isValidationError(err error) bool {
	_, ok := err.(exception.ErrorValidation)
	if !ok {
		return false
	}

	return true
}
