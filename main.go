package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/Prototype-1/SourceService/handler"
    "github.com/Prototype-1/SourceService/model"
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    var allUsers []model.UserProfile
    userHandler := handler.NewUserHandler(logger, &allUsers)

    mux := http.NewServeMux()
    mux.HandleFunc("/users/changes", userHandler.UsersHandler)

    srv := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    // Run server in goroutine in order to avoid codependency
    go func() {
        logger.Info("SourceService started on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("ListenAndServe failed", zap.Error(err))
        }
    }()

    // Wait for interrupt signal to gracefully shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    <-quit
    logger.Info("Shutdown signal received")

    // Create context with timeout to allow graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        logger.Fatal("Server forced to shutdown", zap.Error(err))
    }

    logger.Info("SourceService exiting gracefully")
    fmt.Println("SourceService exited")
}
