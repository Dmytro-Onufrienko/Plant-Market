package orders

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(authed, admin *gin.RouterGroup, pool *pgxpool.Pool) {
	h := NewHandler(pool)
	authed.POST("/orders", h.Checkout)
	authed.GET("/orders", h.ListMine)
	authed.GET("/orders/:id", h.GetMine)
	admin.GET("/orders", h.AdminList)
	admin.PUT("/orders/:id/status", h.AdminUpdateStatus)
}
