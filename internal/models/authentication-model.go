package models

type AuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
