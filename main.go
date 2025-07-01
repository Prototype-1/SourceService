package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"
    "github.com/google/uuid"
    "github.com/Prototype-1/SourceService/model" // replace with your actual module name
)

var allUsers []model.UserProfile

func main() {
    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/users/changes", usersHandler)

    fmt.Println("SourceService running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var responseBatch []model.UserProfile

    // Generate 2-5 new users
    newCount := rand.Intn(4) + 2
    for i := 0; i < newCount; i++ {
        user := generateRandomUser()
        allUsers = append(allUsers, user)
        responseBatch = append(responseBatch, user)
    }

    // Update 2-5 existing users if any
    if len(allUsers) > 0 {
        updateCount := rand.Intn(4) + 2
        for i := 0; i < updateCount; i++ {
            index := rand.Intn(len(allUsers))
            allUsers[index].Status = randomStatus()
            allUsers[index].LastUpdatedAt = time.Now().Format(time.RFC3339)
            responseBatch = append(responseBatch, allUsers[index])
        }
    }

    json.NewEncoder(w).Encode(responseBatch)
}

func generateRandomUser() model.UserProfile {
    return model.UserProfile{
        ID:            uuid.New().String(),
        Name:          fmt.Sprintf("User%d", rand.Intn(1000)),
        Email:         fmt.Sprintf("user%d@example.com", rand.Intn(10000)),
        Mobile:        fmt.Sprintf("+91%010d", rand.Intn(1000000000)),
        Status:        randomStatus(),
        LastUpdatedAt: time.Now().Format(time.RFC3339),
    }
}

func randomStatus() string {
    statuses := []string{"active", "inactive", "pending"}
    return statuses[rand.Intn(len(statuses))]
}
