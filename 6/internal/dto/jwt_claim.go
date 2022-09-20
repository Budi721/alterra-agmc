package dto

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID     uint   `json:"user_id"`
	Email      string `json:"email"`
	RoleID     uint   `json:"role_id"`
	DivisionID uint   `json:"division_id"`
	jwt.RegisteredClaims
}
