package model

type UserProfile struct {
    ID            string `json:"id"`
    Name          string `json:"name"`
    Email         string `json:"email"`
    Mobile        string `json:"mobile"`
    Status        string `json:"status"`
    LastUpdatedAt string `json:"last_updated_at"`
}
