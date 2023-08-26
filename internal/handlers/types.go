package handlers

import (
	"as/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Response structure for API responses
type Response struct {
	Status      int         `json:"status"`
	Msg         string      `json:"msg"`
	Data        interface{} `json:"data"`
	Application string
}

// Marshal returns JSON-encoded response
func (r *Response) Marshal() []byte {
	jsonData := map[string]interface{}{
		"status": r.Status,
		"msg":    r.Msg,
		"data":   r.Data,
	}

	res, err := json.Marshal(jsonData)
	utils.Error(err)
	return res
}

// Send sends the response to the client
func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", fmt.Sprintf("application/%s", r.Application))
	w.WriteHeader(r.Status)
	_, err := w.Write(r.Marshal())
	utils.Error(err)
}
