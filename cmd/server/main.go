
package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/yourname/feedbacksvc/internal/config"
    "github.com/yourname/feedbacksvc/internal/httpserver"
    "github.com/yourname/feedbacksvc/internal/pubsub/stdout"
    "github.com/yourname/feedbacksvc/internal/store/memory"
)

func main() {
    cfg := config.FromEnv()

    // Dependencies
    store := memory.New()
    publisher := stdout.New()

    r := httpserver.NewRouter(store, publisher, cfg)

    srv := &http.Server{
        Addr:              ":" + cfg.Port,
        Handler:           r,
        ReadTimeout:       10 * time.Second,
        WriteTimeout:      10 * time.Second,
        ReadHeaderTimeout: 5 * time.Second,
        IdleTimeout:       60 * time.Second,
    }

    log.Printf("feedbacksvc listening on :%s (env=%s)", cfg.Port, cfg.Env)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Println("server error:", err)
        os.Exit(1)
    }
}
