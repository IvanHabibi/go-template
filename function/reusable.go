package function

import (
	"encoding/json"
	"net/http"
)

// BaseRespons base response
type BaseRespons struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// SendResponse ..
func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) string {
	res := new(BaseRespons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res.StatusCode = statusCode
	res.Message = message
	res.Data = data
	json.NewEncoder(w).Encode(res)
	return message
}
