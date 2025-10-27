package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/temo927/feedbacksvc/internal/config"
	httpserver "github.com/temo927/feedbacksvc/internal/http"
	"github.com/temo927/feedbacksvc/internal/pubsub/stdout"
	"github.com/temo927/feedbacksvc/internal/store"
	firestorestore "github.com/temo927/feedbacksvc/internal/store/firestore"
)

func mustFirestoreStore(ctx context.Context, cfg config.Config) store.Store {
	collection := os.Getenv("FIRESTORE_COLLECTION")
	st, err := firestorestore.New(ctx, cfg.ProjectID, collection)
	if err != nil {
		log.Fatalf("failed to init Firestore store: %v", err)
	}
	return st
}

func main() {
	cfg := config.FromEnv()

	// Build deps
	ctx := context.Background()
	st := mustFirestoreStore(ctx, cfg)
	publisher := stdout.New()

	r := httpserver.NewRouter(st, publisher, cfg)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("feedbacksvc listening on :%s (env=%s, store=firestore, project=%s, collection=%s)",
		cfg.Port, cfg.Env, cfg.ProjectID, os.Getenv("FIRESTORE_COLLECTION"))

	// Run server in background
	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	// Graceful shutdown on SIGINT/SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stop:
		log.Printf("shutdown signal: %s", sig)
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}

	shCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = srv.Shutdown(shCtx)

	// Close Firestore client if available
	if c, ok := st.(interface{ Close() error }); ok {
		_ = c.Close()
	}

	log.Println("server stopped gracefully")
}
