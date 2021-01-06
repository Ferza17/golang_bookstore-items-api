package http_utils

import (
	"encoding/json"
	restErr "github.com/Ferza17/golang_bookstore-items-api/utils/errors_utils"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err restErr.RestError) {
	RespondJSON(w, err.Status, err)
}