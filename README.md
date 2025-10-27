
  # Feedback Service (Go + chi)

  Minimal feedback microservice designed for a 2‑hour tech test. 
  Exposes two endpoints, stores feedback using an abstract storage interface (in‑memory by default), and publishes a message to a Pub/Sub publisher (stdout simulation).

  ## Endpoints
  - `POST /feedback` — create feedback; persists and publishes
  - `GET /health` — returns `{ "status": "ok" }`

  ## Quick Start

  1. Clone this repo (or unzip the provided archive).
  2. Copy env file and edit if desired:
     ```bash
     cp .env.example .env
     ```
  3. Run the server:
     ```bash
     go run ./cmd/server
     ```
     Server listens on `:$PORT` (default `8080`).

  ### cURL
  ```bash
  curl -s -X POST http://localhost:8080/feedback \
-H "Content-Type: application/json" \
-d '{"name":"Alice","email":"alice@example.com","message":"Great service!"}'
  ```
  ```bash
  curl -s http://localhost:8080/health
  ```

  ## Project Structure

  ```text
  cmd/server/               # main entrypoint
  internal/config/          # env config
  internal/domain/          # domain models
  internal/http/            # router and HTTP setup
    handlers/               # HTTP handlers
  internal/store/           # storage interfaces
    memory/                 # in‑memory store impl
  internal/pubsub/          # publisher interfaces
    stdout/                 # stdout publisher impl
  pkg/response/             # response helpers
  ```

  ## Notes

  - **Storage:** In‑memory storage implements the `Store` interface. A Firestore or Cloud SQL impl can be added by implementing the interface and wiring it in `cmd/server/main.go` using build tags or env.
  - **Publisher:** The stdout publisher simulates GCP Pub/Sub and prints the published payload and topic.
  - **Validation:** Basic field validation (non‑empty; minimal email check).

  ## Docker (optional)
  ```dockerfile
  # see Dockerfile in repo
  ```
  Build & run:
  ```bash
  docker build -t feedbacksvc:dev .
  docker run --env-file .env -p 8080:8080 feedbacksvc:dev
  ```

  ## Environment
  ```env
  GCP_PROJECT_ID=test-project
  PUBSUB_TOPIC=feedback-topic
  GO_ENV=dev
  PORT=8080
  ```
