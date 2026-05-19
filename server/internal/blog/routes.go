package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(public, admin *gin.RouterGroup, pool *pgxpool.Pool) {
	h := NewHandler(pool)
	public.GET("/blog", h.List)
	public.GET("/blog/:slug", h.GetBySlug)
	admin.GET("/blog", h.AdminList)
	admin.POST("/blog", h.AdminCreate)
	admin.PUT("/blog/:id", h.AdminUpdate)
	admin.DELETE("/blog/:id", h.AdminDelete)
}
