package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Budi721/alterra-agmc/v6/internal/dto"
	"github.com/Budi721/alterra-agmc/v6/pkg/util"
	"github.com/golang-jwt/jwt/v4"
)

var (
	JwtSecret        = []byte(util.Getenv("JWT_SECRET", "testsecret"))
	JwtExp           = time.Duration(1) * time.Hour
	JwtSigningMethod = jwt.SigningMethodHS256
)

func getTokenString(authHeader string) (*string, error) {
	var token string
	if strings.Contains(authHeader, "Bearer") {
		token = strings.Replace(authHeader, "Bearer ", "", -1)
		return &token, nil
	}
	return nil, fmt.Errorf("authorization not found")
}

func CreateJWTClaims(email string, userID uint) dto.JWTClaims {
	return dto.JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JwtExp)),
		},
	}
}

func CreateJWTToken(claims dto.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(JwtSigningMethod, claims)
	return token.SignedString([]byte(JwtSecret))
}

func ParseJWTToken(authHeader string) (*dto.JWTClaims, error) {
	tokenString, err := getTokenString(authHeader)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != JwtSigningMethod {
			return nil, fmt.Errorf("invalid signing method")
		}
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimsStr, err := json.Marshal(claims)
		if err != nil {
			return nil, fmt.Errorf("error when marshalling token")
		}

		var customClaims dto.JWTClaims
		if err := json.Unmarshal(claimsStr, &customClaims); err != nil {
			return nil, fmt.Errorf("error when unmarshalling token")
		}

		return &customClaims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
