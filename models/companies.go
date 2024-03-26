package models

// Company es una estructura que representa una empresa.
type Company struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
	Location  string `json:"location"`
	Industry  string `json:"industry"`
	UserID    string `json:"user_id"`
}