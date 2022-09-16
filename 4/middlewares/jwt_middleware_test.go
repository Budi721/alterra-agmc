package middlewares

import (
	"github.com/Budi721/alterra-agmc/v2/constants"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "sukses create token sesuai id", args: struct{ id uint }{id: uint(1)}, want: "1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.id)

			if assert.NoError(t, err) {
				token, _ := jwt.Parse(got, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						t.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					return []byte(constants.SecretJwt), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					assert.Same(t, claims["jti"], tt.want)
				}
			}
		})
	}
}
