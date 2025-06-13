package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var AcessSecret = []byte(os.Getenv("JWT_SECRET_ACCESS"))
var RefreshSecret = []byte(os.Getenv("JWT_SECRET_REFRESH"))

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// acesstoken generation
func GenerateAcessToken(userID, role string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AcessSecret)
}

// refresh token genteration
func GenerateRefreshToken(userID, role string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(RefreshSecret)
}

// function for validation acess token and refresh token
func ValidateToken(tokenStr string, isAcessToken bool) (*CustomClaims, error) {
	secret := AcessSecret
	if !isAcessToken {
		secret = RefreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err

	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
