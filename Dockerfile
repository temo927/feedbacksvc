
# syntax=docker/dockerfile:1
FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o feedbacksvc ./cmd/server

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /app/feedbacksvc /feedbacksvc
ENV PORT=8080
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/feedbacksvc"]
