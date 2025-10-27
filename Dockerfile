# syntax=docker/dockerfile:1

# Use a Go >= 1.24 toolchain (1.25 is fine)
FROM golang:1.25 AS build
WORKDIR /app

# Copy mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags "-s -w" -o feedbacksvc ./cmd/server

# Small, non-root runtime image
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=build /app/feedbacksvc /feedbacksvc
ENV PORT=8080
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/feedbacksvc"]
