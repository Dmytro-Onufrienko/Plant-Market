package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/services"
)

const (
	CookieName  = "monstera_token"
	CtxUserID   = "userID"
	CtxUserRole = "userRole"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(CookieName)
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": nil, "error": "unauthorized"})
			return
		}

		claims, err := services.ValidateToken(token, jwtSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": nil, "error": "unauthorized"})
			return
		}

		c.Set(CtxUserID, claims.UserID)
		c.Set(CtxUserRole, claims.Role)
		c.Next()
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get(CtxUserRole)
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"data": nil, "error": "forbidden"})
			return
		}
		c.Next()
	}
}
