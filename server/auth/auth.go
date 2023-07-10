package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt"
)

type Auth struct {
	AuthKey               rsa.PrivateKey
	JWTExpirationDuration time.Duration
}

type AuthConfig struct {
	AuthKeyLocation       string        `json:"authKeyLocation"`
	JWTExpirationDuration time.Duration `json:"jwtExpirationDuration"`
}

func loadKey(keyLocation string) (*rsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
}

func NewAuth(config AuthConfig) (*Auth, error) {
	authKey, err := loadKey(config.AuthKeyLocation)
	if err != nil {
		return nil, err
	}

	return &Auth{
		AuthKey:               *authKey,
		JWTExpirationDuration: config.JWTExpirationDuration,
	}, nil
}

func (auth Auth) GenerateToken(username string) (string, error) {
	claimsMap := jwt.MapClaims(map[string]interface{}{
		"username": username,
	})

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claimsMap)
	tokenString, err := token.SignedString(&auth.AuthKey)
	if err != nil {
		return "", fmt.Errorf("Error generating token: %w", err)
	}

	return tokenString, nil
}

func (auth Auth) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return &auth.AuthKey.PublicKey, nil
	})
	if err != nil {
		return "", errors.New("Error parsing token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("Invalid token")
	}

	return username, nil
}
