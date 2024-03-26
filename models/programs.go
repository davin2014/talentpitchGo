package models

import "time"

// Program es una estructura que representa un programa de formación.
type Program struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UserID      string    `json:"user_id"`
}