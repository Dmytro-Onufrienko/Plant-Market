package reviews

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(public, authed, admin *gin.RouterGroup, pool *pgxpool.Pool) {
	h := NewHandler(pool)
	public.GET("/reviews/latest", h.Latest)
	authed.POST("/reviews", h.Create)
	admin.GET("/reviews", h.AdminList)
	admin.DELETE("/reviews/:id", h.AdminDelete)
}
