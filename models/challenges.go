package models



type Challenge struct {
    Id          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Difficulty  int       `json:"difficulty"`
    UserID      string    `json:"user_id"`
}