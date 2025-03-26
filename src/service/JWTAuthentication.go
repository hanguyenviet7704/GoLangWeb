package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AuthCustomClaims struct {
	Email  string `json:"email"`
	UserID int    `json:"userId"`
	jwt.StandardClaims
}
type JWTService interface {
	GenerateToken(email string, UserID int, expiration time.Duration) string
	ValidateToken(token string) (*AuthCustomClaims, error)
}
type jwtServices struct {
	secretKey string
	issuer    string
}

func GetSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: GetSecretKey(),
		issuer:    "VietHa",
	}
}
func (service *jwtServices) GenerateToken(email string, UserID int, expiration time.Duration) string {
	claims := &AuthCustomClaims{
		Email:  email,
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (service *jwtServices) ValidateToken(encodedToken string) (*AuthCustomClaims, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token algorithm: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("Token is expired")
		}
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}
