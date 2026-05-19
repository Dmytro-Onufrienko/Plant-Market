package reviews

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

// GET /api/reviews/latest
func (h *Handler) Latest(c *gin.Context) {
	limit := httputil.IntQuery(c, "limit", 3)
	reviews, err := GetLatest(c.Request.Context(), h.pool, limit)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	httputil.RespondOK(c, gin.H{"reviews": reviews})
}

// POST /api/reviews
func (h *Handler) Create(c *gin.Context) {
	userID := c.GetString(middleware.CtxUserID)

	var req CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	exists, err := CheckExists(c.Request.Context(), h.pool, userID, req.PlantID)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if exists {
		httputil.RespondError(c, http.StatusConflict, "you have already reviewed this plant")
		return
	}

	if err := Create(c.Request.Context(), h.pool, userID, req.PlantID, req.Rating, req.Text); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not create review")
		return
	}
	httputil.RespondCreated(c, nil)
}

// GET /api/admin/reviews
func (h *Handler) AdminList(c *gin.Context) {
	reviews, err := GetAll(c.Request.Context(), h.pool)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if reviews == nil {
		reviews = []ReviewWithDetails{}
	}
	httputil.RespondOK(c, gin.H{"reviews": reviews})
}

// DELETE /api/admin/reviews/:id
func (h *Handler) AdminDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}

	if err := Delete(c.Request.Context(), h.pool, id); err != nil {
		if err.Error() == "not found" {
			httputil.RespondError(c, http.StatusNotFound, "review not found")
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "could not delete review")
		return
	}
	httputil.RespondOK(c, nil)
}
