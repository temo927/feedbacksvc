package handlers

import (
   "context"
   "crypto/rand"
   "encoding/hex"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
   "time"

   "github.com/temo927/feedbacksvc/internal/config"
   "github.com/temo927/feedbacksvc/internal/domain"
   "github.com/temo927/feedbacksvc/internal/pubsub"
   "github.com/temo927/feedbacksvc/internal/store"
   "github.com/temo927/feedbacksvc/pkg/response"
)

type Handlers struct {
   store store.Store
   pub   pubsub.Publisher
   cfg   config.Config
}

func New(st store.Store, pub pubsub.Publisher, cfg config.Config) *Handlers {
   return &Handlers{store: st, pub: pub, cfg: cfg}
}

type createFeedbackRequest struct {
   Name    string `json:"name"`
   Email   string `json:"email"`
   Message string `json:"message"`
}

func (h *Handlers) CreateFeedback(w http.ResponseWriter, r *http.Request) {
   var req createFeedbackRequest
   if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
       response.Error(w, http.StatusBadRequest, "invalid JSON body")
       return
   }

   if err := validate(req); err != nil {
       response.Error(w, http.StatusBadRequest, err.Error())
       return
   }

   id := randID()
   fb := domain.Feedback{
       ID:        id,
       Name:      strings.TrimSpace(req.Name),
       Email:     strings.TrimSpace(req.Email),
       Message:   strings.TrimSpace(req.Message),
       CreatedAt: time.Now().UTC(),
   }

   if err := h.store.Save(r.Context(), fb); err != nil {
       response.Error(w, http.StatusInternalServerError, "failed to persist feedback")
       return
   }

   // Fire and forget publish (still check context)
   _ = h.pub.Publish(r.Context(), h.cfg.PubsubTopic, map[string]any{
       "id": fb.ID, "name": fb.Name, "email": fb.Email, "message": fb.Message, "created_at": fb.CreatedAt,
   })

   response.JSON(w, http.StatusCreated, response.Envelope{
       "data": fb,
   })
}

func validate(req createFeedbackRequest) error {
   if strings.TrimSpace(req.Name) == "" {
       return errors.New("name is required")
   }
   if strings.TrimSpace(req.Email) == "" || !strings.Contains(req.Email, "@") {
       return errors.New("valid email is required")
   }
   if strings.TrimSpace(req.Message) == "" {
       return errors.New("message is required")
   }
   return nil
}

func randID() string {
   var b [16]byte
   _, _ = rand.Read(b[:])
   return hex.EncodeToString(b[:])
}

// helper to satisfy context usage in lints
var _ = context.TODO

