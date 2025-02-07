package models


// RegisterInput is used for user registration
type RegisterInput struct {
	Username string `json:"username" example:"testuser" default:"testuser"`
	Email    string `json:"email" example:"user@example.com" default:"user@example.com"`
	Password string `json:"password" example:"123456" default:"123456"`
}


type LoginInput struct {
	Email    string `json:"email" example:"test@example.com" default:"user@example.com"`
	Password string `json:"password" example:"password123" default:"123456"`
}