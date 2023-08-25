package utils

import (
	"fmt"
	"net/http"
)

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func HttpError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
