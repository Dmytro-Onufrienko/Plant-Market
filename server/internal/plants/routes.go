package plants

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(public, admin *gin.RouterGroup, pool *pgxpool.Pool) {
	h := NewHandler(pool)
	public.GET("/plants", h.List)
	public.GET("/plants/by-id/:id", h.GetByID)
	public.GET("/plants/:slug", h.GetBySlug)
	public.GET("/categories", h.ListCategories)
	admin.GET("/plants", h.AdminList)
	admin.POST("/plants", h.AdminCreate)
	admin.PUT("/plants/:id", h.AdminUpdate)
	admin.DELETE("/plants/:id", h.AdminDelete)
}
