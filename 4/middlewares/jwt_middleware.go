package middlewares

import (
	"strconv"
	"time"

	"github.com/Budi721/alterra-agmc/v2/constants"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

// GenerateToken function to create new token
func GenerateToken(id uint) (string, error) {
	// Set custom claims
	claims := &jwt.StandardClaims{
		Id:        strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(constants.SecretJwt))
}

// JWTMiddleware middleware to validate user to the response.
var JWTMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(constants.SecretJwt),
})
