package handlers

import "context"

const (
	user_token = "user_token"
	admin_token = "admin_token"
)

type AuthHandler struct {
}

func CreateAuthHandler() *AuthHandler{
	return &AuthHandler{}
}

func (handler *AuthHandler) Login(ctx context.Context, form LoginForm) (string, error) {
	return "", nil
}