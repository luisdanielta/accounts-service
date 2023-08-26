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

func HttpSendError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(e.Error()))
	Error(err)
}
