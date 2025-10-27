
package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/temo927/feedbacksvc/internal/config"
    "github.com/temo927/feedbacksvc/internal/http"
    "github.com/temo927/feedbacksvc/internal/pubsub/stdout"
    "github.com/temo927/feedbacksvc/internal/store/memory"
)

func TestCreateFeedback(t *testing.T) {
    cfg := config.FromEnv()
    st := memory.New()
    pub := stdout.New()
    r := httpserver.NewRouter(st, pub, cfg)

    body := map[string]string{
        "name": "Alice",
        "email": "alice@example.com",
        "message": "Great service!",
    }
    b, _ := json.Marshal(body)

    req := httptest.NewRequest(http.MethodPost, "/feedback", bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    if w.Code != http.StatusCreated {
        t.Fatalf("expected 201 Created, got %d, body=%s", w.Code, w.Body.String())
    }
}
