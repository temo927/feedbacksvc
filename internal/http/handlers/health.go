
package handlers

import (
    "net/http"

    "github.com/yourname/feedbacksvc/pkg/response"
)

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
    response.JSON(w, http.StatusOK, map[string]string{"status":"ok"})
}
