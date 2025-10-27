
package domain

import "time"

type Feedback struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
}
