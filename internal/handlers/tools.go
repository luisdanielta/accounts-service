package handlers

import "net/http"

// ValidateRequiredFields checks if the required fields are present in the request
func ValidateRequiredFields(r *http.Request, fields []string) bool {
	for _, field := range fields {
		if r.FormValue(field) == "" {
			return false
		}
	}
	return true
}
