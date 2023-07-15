package middleware

import (
	"context"
	"net/http"

	"tippers-back/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type authString string

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		bearer := "Bearer "
		token := auth[len(bearer):]
		validate, err := service.JwtValidate(token)
		if err != nil && !validate.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		claims, ok := validate.Claims.(*jwt.MapClaims)
		if ok {
			c.Set("user_id", (*claims)["user_id"].(float64))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is not valid",
			})
		}

	}
}

func CtxValue(ctx context.Context) *jwt.MapClaims {
	raw, _ := ctx.Value(authString("auth")).(*jwt.MapClaims)
	return raw
}
