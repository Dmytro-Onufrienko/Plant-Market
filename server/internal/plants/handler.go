package plants

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/reviews"
)

type Handler struct {
	pool *pgxpool.Pool
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

// GET /api/plants
func (h *Handler) List(c *gin.Context) {
	f := Filter{
		CategorySlug: c.Query("category"),
		Page:         httputil.IntQuery(c, "page", 1),
		Limit:        httputil.IntQuery(c, "limit", 20),
	}
	if v := c.Query("featured"); v == "true" {
		t := true
		f.Featured = &t
	}

	result, total, err := GetAll(c.Request.Context(), h.pool, f)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if result == nil {
		result = []Plant{}
	}

	httputil.RespondOK(c, gin.H{
		"plants": result,
		"total":  total,
		"page":   f.Page,
	})
}

// GET /api/plants/by-id/:id
func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}
	plant, err := GetByID(c.Request.Context(), h.pool, id)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if plant == nil {
		httputil.RespondError(c, http.StatusNotFound, "plant not found")
		return
	}
	httputil.RespondOK(c, gin.H{"plant": plant})
}

// GET /api/plants/:slug
func (h *Handler) GetBySlug(c *gin.Context) {
	plant, err := GetBySlug(c.Request.Context(), h.pool, c.Param("slug"))
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if plant == nil {
		httputil.RespondError(c, http.StatusNotFound, "plant not found")
		return
	}

	plantReviews, err := reviews.GetByPlant(c.Request.Context(), h.pool, plant.ID)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if plantReviews == nil {
		plantReviews = []reviews.ReviewWithUser{}
	}

	var avgRating float64
	if len(plantReviews) > 0 {
		sum := 0
		for _, r := range plantReviews {
			sum += r.Rating
		}
		avgRating = float64(sum) / float64(len(plantReviews))
	}

	httputil.RespondOK(c, gin.H{
		"plant":      plant,
		"reviews":    plantReviews,
		"avg_rating": avgRating,
	})
}

// GET /api/categories
func (h *Handler) ListCategories(c *gin.Context) {
	cats, err := GetAllCategories(c.Request.Context(), h.pool)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if cats == nil {
		cats = []Category{}
	}
	httputil.RespondOK(c, gin.H{"categories": cats})
}

// ─── Admin ────────────────────────────────────────────────────────────────────

// GET /api/admin/plants
func (h *Handler) AdminList(c *gin.Context) {
	result, _, err := GetAll(c.Request.Context(), h.pool, Filter{Page: 1, Limit: 1000})
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if result == nil {
		result = []Plant{}
	}
	httputil.RespondOK(c, gin.H{"plants": result})
}

// POST /api/admin/plants
func (h *Handler) AdminCreate(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	plant, err := Create(c.Request.Context(), h.pool, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not create plant")
		return
	}
	httputil.RespondCreated(c, gin.H{"plant": plant})
}

// PUT /api/admin/plants/:id
func (h *Handler) AdminUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}

	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	plant, err := Update(c.Request.Context(), h.pool, id, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not update plant")
		return
	}
	if plant == nil {
		httputil.RespondError(c, http.StatusNotFound, "plant not found")
		return
	}
	httputil.RespondOK(c, gin.H{"plant": plant})
}

// DELETE /api/admin/plants/:id
func (h *Handler) AdminDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}

	plant, err := GetByID(c.Request.Context(), h.pool, id)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if plant == nil {
		httputil.RespondError(c, http.StatusNotFound, "plant not found")
		return
	}

	if err := Delete(c.Request.Context(), h.pool, id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not delete plant")
		return
	}
	httputil.RespondOK(c, nil)
}
