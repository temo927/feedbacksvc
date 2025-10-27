
package handlers

import (
    "net/http"

    "github.com/temo927/feedbacksvc/pkg/response"
)

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
    response.JSON(w, http.StatusOK, map[string]string{"status":"ok"})
}
