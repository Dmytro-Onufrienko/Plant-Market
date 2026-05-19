package orders

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/middleware"
)

type Handler struct {
	pool *pgxpool.Pool
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

// POST /api/orders
func (h *Handler) Checkout(c *gin.Context) {
	userID := c.GetString(middleware.CtxUserID)

	var req CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	order, err := Create(c.Request.Context(), h.pool, userID, req.Items, req.Shipping)
	if err != nil {
		switch err.Error() {
		case "not found":
			httputil.RespondError(c, http.StatusNotFound, err.Error())
		default:
			httputil.RespondError(c, http.StatusBadRequest, err.Error())
		}
		return
	}
	httputil.RespondCreated(c, gin.H{"order": order})
}

// GET /api/orders
func (h *Handler) ListMine(c *gin.Context) {
	userID := c.GetString(middleware.CtxUserID)

	result, err := GetByUser(c.Request.Context(), h.pool, userID)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if result == nil {
		result = []Order{}
	}
	httputil.RespondOK(c, gin.H{"orders": result})
}

// GET /api/orders/:id
func (h *Handler) GetMine(c *gin.Context) {
	userID := c.GetString(middleware.CtxUserID)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}

	order, err := GetByID(c.Request.Context(), h.pool, id, userID)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if order == nil {
		httputil.RespondError(c, http.StatusNotFound, "order not found")
		return
	}
	httputil.RespondOK(c, gin.H{"order": order})
}

// ─── Admin ────────────────────────────────────────────────────────────────────

// GET /api/admin/orders
func (h *Handler) AdminList(c *gin.Context) {
	result, err := GetAll(c.Request.Context(), h.pool)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if result == nil {
		result = []OrderWithUser{}
	}
	httputil.RespondOK(c, gin.H{"orders": result})
}

// PUT /api/admin/orders/:id/status
func (h *Handler) AdminUpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	if err := UpdateStatus(c.Request.Context(), h.pool, id, req.Status); err != nil {
		if err.Error() == "not found" {
			httputil.RespondError(c, http.StatusNotFound, "order not found")
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "could not update status")
		return
	}
	httputil.RespondOK(c, nil)
}
