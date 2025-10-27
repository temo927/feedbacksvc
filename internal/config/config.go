package config

import "os"

type Config struct {
    ProjectID string
    PubsubTopic string
    Env string
    Port string
}

func getEnv(key, def string) string {
    v := os.Getenv(key)
    if v == "" {
        return def
    }
    return v
}

func FromEnv() Config {
    return Config{
        ProjectID:  getEnv("GCP_PROJECT_ID", "test-project"),
        PubsubTopic:getEnv("PUBSUB_TOPIC", "feedback-topic"),
        Env:        getEnv("GO_ENV", "dev"),
        Port:       getEnv("PORT", "8080"),
    }
}
