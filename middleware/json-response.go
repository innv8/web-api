package middleware

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, code int, response interface{}) {
	responseBytes, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(responseBytes)
}
