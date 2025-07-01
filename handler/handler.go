package handler

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
	"fmt"
    "github.com/google/uuid"
    "go.uber.org/zap"
     "github.com/Prototype-1/SourceService/model" 
)

type UserHandler struct {
    Logger   *zap.Logger
    AllUsers *[]model.UserProfile
}

func NewUserHandler(logger *zap.Logger, users *[]model.UserProfile) *UserHandler {
    return &UserHandler{
        Logger:   logger,
        AllUsers: users,
    }
}

func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var responseBatch []model.UserProfile

    // Generate 2-5 new users
    newCount := rand.Intn(4) + 2
    for i := 0; i < newCount; i++ {
        user := h.generateRandomUser()
        *h.AllUsers = append(*h.AllUsers, user)
        responseBatch = append(responseBatch, user)

        h.Logger.Info("Generated new user", zap.String("id", user.ID), zap.String("email", user.Email))
    }

    // Update 2-5 existing users
    if len(*h.AllUsers) > 0 {
        updateCount := rand.Intn(4) + 2
        for i := 0; i < updateCount; i++ {
            index := rand.Intn(len(*h.AllUsers))
            (*h.AllUsers)[index].Status = randomStatus()
            (*h.AllUsers)[index].LastUpdatedAt = time.Now().Format(time.RFC3339)
            responseBatch = append(responseBatch, (*h.AllUsers)[index])

            h.Logger.Info("Updated existing user", zap.String("id", (*h.AllUsers)[index].ID), zap.String("status", (*h.AllUsers)[index].Status))
        }
    }

    json.NewEncoder(w).Encode(responseBatch)
}

func (h *UserHandler) generateRandomUser() model.UserProfile {
    return model.UserProfile{
        ID:            uuid.New().String(),
        Name:          randomName(),
        Email:         randomEmail(),
        Mobile:        randomMobile(),
        Status:        randomStatus(),
        LastUpdatedAt: time.Now().Format(time.RFC3339),
    }
}

func randomName() string {
    return "User" + randomNumber(1000)
}

func randomEmail() string {
    return "user" + randomNumber(10000) + "@example.com"
}

func randomMobile() string {
    return "+91" + randomNumber(1000000000)
}

func randomStatus() string {
    statuses := []string{"active", "inactive", "pending"}
    return statuses[rand.Intn(len(statuses))]
}

func randomNumber(max int) string {
    return fmt.Sprintf("%d", rand.Intn(max))
}
