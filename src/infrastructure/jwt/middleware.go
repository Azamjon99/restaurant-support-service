package jwt

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	headerKeyAuthorization = "Authorization"
)

type Middleware interface {
	RequireToken() gin.HandlerFunc
}

type middlewareImpl struct {
	svc Service
}

func newMiddlware(svc Service) Middleware {
	return &middlewareImpl{
		svc: svc,
	}
}

func (m *middlewareImpl) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		tokenString := c.GetHeader(headerKeyAuthorization)
		if tokenString == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid credentials"))
			return
		}

		token, err := m.svc.ParseToken(ctx, tokenString)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx = TokenInCtx(ctx, token)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
