# 🧠 Feedback Service (Go + Chi + Firestore Emulator)

A minimal feedback microservice built in Go for a 2-hour technical challenge.  
It exposes two endpoints, stores feedback in **Google Firestore (via emulator)**,  
and simulates Pub/Sub publishing via **stdout logs**.

---

## ⚙️ Features
- **`POST /feedback`** — Creates feedback; stores it in Firestore emulator and logs a Pub/Sub-like message.  
- **`GET /health`** — Returns `{ "status": "ok" }`.  
- Uses **Firestore Emulator** (no GCP credentials or billing required).  
- Simulated **Pub/Sub** using console logs (`fmt.Printf`).  
- Fully Dockerized setup — single command to start everything.  

---

## 🚀 Quick Start (Docker Compose)

### 1️⃣ Build and start services
```bash
docker compose up --build
```

This will:
- Start the **Firestore Emulator** (Java 21 inside a container) on **localhost:8081**  
- Start the **Feedback API** on **localhost:8080**

---

## 🧪 Test the Endpoints

### ✅ Health Check
```bash
curl -s http://localhost:8080/health
```
**Expected:**
```json
{"status":"ok"}
```

---

### 💬 Create Feedback
```bash
curl -s -X POST http://localhost:8080/feedback   -H "Content-Type: application/json"   -d '{"name":"Alice","email":"alice@example.com","message":"Great service!"}'
```
**Expected response:**
```json
{
  "data": {
    "id": "4b7e8c...",
    "name": "Alice",
    "email": "alice@example.com",
    "message": "Great service!",
    "created_at": "2025-10-27T14:55:10Z"
  }
}
```

---

## 🔍 Check Firestore Data (via Emulator REST API)

The Firestore emulator runs on **localhost:8081**.  
You can query stored data with:

```bash
curl -s -X POST "http://localhost:8081/v1/projects/test-project/databases/(default)/documents:runQuery"   -H "Content-Type: application/json"   -d '{"structuredQuery":{"from":[{"collectionId":"feedback"}],"limit":10}}' | jq
```
> If you don’t have `jq`, just remove the pipe.

**Example output:**
```json
{
  "id": { "stringValue": "4b7e8c..." },
  "name": { "stringValue": "Alice" },
  "email": { "stringValue": "alice@example.com" },
  "message": { "stringValue": "Great service!" }
}
```

---

## 🪵 View “Pub/Sub” Logs

Pub/Sub messages are simulated and printed directly to the application logs.  
You can follow them in real time:

```bash
docker compose logs -f app
```

**Example:**
```
[pubsub:feedback-topic] 2025-10-27T14:55:10Z {"id":"4b7e8c...","name":"Alice","email":"alice@example.com","message":"Great service!","created_at":"..."}
```

---

## 🧰 Environment Variables (`.env`)

```env
GO_ENV=dev
PORT=8080

# Firestore emulator configuration
GCP_PROJECT_ID=test-project
FIRESTORE_COLLECTION=feedback
# The app automatically connects to the emulator via docker-compose
# FIRESTORE_EMULATOR_HOST=firestore-emulator:8080

# Pub/Sub simulation
PUBSUB_TOPIC=feedback-topic
```

---

## 🧱 Project Structure

```text
cmd/server/                # main entrypoint
internal/config/           # environment config
internal/domain/           # domain models (Feedback)
internal/http/             # router & middleware setup
  handlers/                # HTTP handlers
internal/store/            # storage interfaces
  firestore/               # Firestore implementation
internal/pubsub/           # publisher interfaces
  stdout/                  # stdout-based publisher (Pub/Sub simulation)
pkg/response/              # JSON response helpers
```

---

## 🧹 Useful Commands

Stop and remove all containers and emulator data:
```bash
docker compose down -v
```

Rebuild only the app:
```bash
docker compose build app
```

Follow app logs only:
```bash
docker compose logs -f app
```

Restart everything cleanly:
```bash
docker compose down -v && docker compose up --build
```

---

## ✅ Summary

| Component | Purpose | Host Port |
|------------|----------|-----------|
| **App (Go)** | REST API (`/feedback`, `/health`) | `8080` |
| **Firestore Emulator** | Local Firestore DB | `8081` |
| **Pub/Sub Simulation** | Console log messages | N/A |

---

**Author:** Temur Rekhviashvili  
**Language:** Go 1.25  
**Database:** Firestore (via emulator)  
**Pub/Sub:** Simulated via stdout logs  
**Runtime:** Docker + Docker Compose  
**Purpose:** Mini Technical Test — Senior Golang Developer (GCP)