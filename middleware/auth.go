package middleware

import (
	"context"

	"tippers-back/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type authString string

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Next()
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := service.JwtValidate(auth)
		if err != nil || !validate.Valid {
			c.Next()
			return
		}

		customClaim, _ := validate.Claims.(*jwt.MapClaims)

		ctx := context.WithValue(c.Request.Context(), authString("auth"), customClaim)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}

func CtxValue(ctx context.Context) *jwt.MapClaims {
	raw, _ := ctx.Value(authString("auth")).(*jwt.MapClaims)
	return raw
}
