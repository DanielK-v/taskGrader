package models

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func NewLoginRequest(email, password string) *LoginRequest {
    return &LoginRequest{
        Email:    email,
        Password: password,
    }
}
