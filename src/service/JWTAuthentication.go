package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"vietha/src/entity"
)

type AuthCustomClaims struct {
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	UserID int      `json:"userId"`
	Roles  []string `json:"role"`
	jwt.StandardClaims
}
type JWTService interface {
	GenerateAccessToken(user *entity.User) string
	GenerateRefreshToken(user *entity.User) string
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
func (service *jwtServices) GenerateAccessToken(user *entity.User) string {
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Name
	}
	claims := &AuthCustomClaims{
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.Id,
		Roles:  roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (service *jwtServices) GenerateRefreshToken(user *entity.User) string {
	claims := &AuthCustomClaims{
		UserID: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 7 * 24).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Name,
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
