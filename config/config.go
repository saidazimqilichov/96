package config

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
)

type Config struct {
    Port     string
    CertFile string
    KeyFile  string
}

func Load() Config {
    return Config{
        Port:     getEnv("PORT", "9000"),
        CertFile: getEnv("CERT_FILE", "localhost.pem"),
        KeyFile:  getEnv("KEY_FILE", "localhost-key.pem"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

func GracefulShutdown(srv *http.Server) {
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop

    log.Println("Shutting down server...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server exited")
}