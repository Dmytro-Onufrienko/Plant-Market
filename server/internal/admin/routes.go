package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(admin *gin.RouterGroup, pool *pgxpool.Pool) {
	h := NewHandler(pool)
	admin.GET("/stats", h.Stats)
}
