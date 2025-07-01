package main

import (
    "fmt"
    "log"
    //"math/rand"
    "net/http"
    //"time"
    "github.com/Prototype-1/SourceService/handler"
    "github.com/Prototype-1/SourceService/model"
    "go.uber.org/zap"
)

// func main() {
//     rand.Seed(time.Now().UnixNano())

//     logger, _ := zap.NewProduction()
//     defer logger.Sync()

//     var allUsers []model.UserProfile

//     userHandler := handler.NewUserHandler(logger, &allUsers)

//     http.HandleFunc("/users/changes", userHandler.UsersHandler)

//     fmt.Println("SourceService running on :8080")
//     logger.Info("SourceService started on :8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }
func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    var allUsers []model.UserProfile

    userHandler := handler.NewUserHandler(logger, &allUsers)

    http.HandleFunc("/users/changes", userHandler.UsersHandler)

    fmt.Println("SourceService running on :8080")
    logger.Info("SourceService started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}