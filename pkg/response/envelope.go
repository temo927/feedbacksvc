package response

import (
    "encoding/json"
    "net/http"
)

type Envelope map[string]any

func JSON(w http.ResponseWriter, status int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, msg string) {
    JSON(w, status, Envelope{"error": msg})
}
