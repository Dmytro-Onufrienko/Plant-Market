package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/orders"
)

type Handler struct {
	pool *pgxpool.Pool
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

// GET /api/admin/stats
func (h *Handler) Stats(c *gin.Context) {
	ctx := c.Request.Context()

	var plantsCount, ordersCount, reviewsCount, usersCount int

	if err := h.pool.QueryRow(ctx, `SELECT COUNT(*) FROM plants`).Scan(&plantsCount); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if err := h.pool.QueryRow(ctx, `SELECT COUNT(*) FROM orders`).Scan(&ordersCount); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if err := h.pool.QueryRow(ctx, `SELECT COUNT(*) FROM reviews`).Scan(&reviewsCount); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if err := h.pool.QueryRow(ctx, `SELECT COUNT(*) FROM users`).Scan(&usersCount); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}

	recentOrders, err := orders.GetRecent(ctx, h.pool, 5)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}

	httputil.RespondOK(c, gin.H{
		"plants":        plantsCount,
		"orders":        ordersCount,
		"reviews":       reviewsCount,
		"users":         usersCount,
		"recent_orders": recentOrders,
	})
}
