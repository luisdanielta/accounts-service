package utils

import (
	"net/http"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func HttpSendError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(e.Error()))
	Error(err)
}
