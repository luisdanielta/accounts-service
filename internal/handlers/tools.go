package handlers

import "net/http"

func ValidateRequiredFields(r *http.Request, fields []string) bool {
	for _, field := range fields {
		if r.FormValue(field) == "" {
			return false
		}
	}
	return true
}
