package utils

import (
	"net/http"
)

func ValidateItemStatus(response *http.Response) bool {
	valid := true

	switch response.StatusCode {
	case 404:
		valid = false
	default:
		valid = true
	}

	return valid
}

func ValidateMatchStatus(response *http.Response) bool {
	valid := true

	switch response.StatusCode {
	case 422:
		valid = false
	default:
		valid = true
	}

	return valid
}
