package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/Azamjon99/restaurant-support-service/src/infrastructure/rand"
)

type Service interface {
	CreateToken(ctx context.Context, subject string) (string, error)
	ParseToken(ctx context.Context, token string) (*Token, error)
	Middleware() Middleware
}

type svcImpl struct {
	secret     string
	expInSec   int
	middleware Middleware
}

func NewService(secret string, expInSec int) Service {
	s := &svcImpl{
		secret:   secret,
		expInSec: expInSec,
	}

	s.middleware = newMiddlware(s)
	return s
}

func (s *svcImpl) CreateToken(ctx context.Context, subject string) (string, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Duration(s.expInSec) * time.Second)

	claims := jwt.RegisteredClaims{
		ID:        rand.UUID(),
		Subject:   subject, // eg. eaterId
		Issuer:    "food.uz",
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *svcImpl) ParseToken(ctx context.Context, tokenValue string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenValue, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, token.Claims.Valid()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	} else {
		fmt.Println(err)
	}

	return &Token{
		UserID: claims["sub"].(string),
	}, nil
}

func (s *svcImpl) Middleware() Middleware {
	return s.middleware
}
