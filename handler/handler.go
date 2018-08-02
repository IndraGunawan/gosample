package handler

import (
	"encoding/json"
	"net/http"
)

// JSONResponse is helper to write JSON format in response
func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal(data)
	w.Write(body)
}
