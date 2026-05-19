package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(authed *gin.RouterGroup, pool *pgxpool.Pool) {
	h := NewHandler(pool)
	authed.PUT("/users/email", h.UpdateEmail)
	authed.PUT("/users/password", h.UpdatePassword)
}
