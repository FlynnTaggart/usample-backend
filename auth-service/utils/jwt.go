package utils

import (
	"auth-service/models"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

func NewJwtWrapper(secretKey string, issuer string, expirationHours int64) *JwtWrapper {
	return &JwtWrapper{
		SecretKey:       secretKey,
		Issuer:          issuer,
		ExpirationHours: expirationHours,
	}
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string
}

func (w JwtWrapper) GenerateToken(user *models.User) (signedToken string, err error) {
	claims := &UserClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{"default"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(w.ExpirationHours))),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w JwtWrapper) ValidateToken(signedToken string) (claims *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*UserClaims)

	if !ok {
		return nil, errors.New("jwt validate: couldn't parse claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("jwt validate: jwt is expired")
	}

	return claims, nil
}
