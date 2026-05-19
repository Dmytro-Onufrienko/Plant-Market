package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
)

type Handler struct {
	pool *pgxpool.Pool
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

// GET /api/blog
func (h *Handler) List(c *gin.Context) {
	posts, err := GetPublished(c.Request.Context(), h.pool)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if posts == nil {
		posts = []Post{}
	}
	httputil.RespondOK(c, gin.H{"posts": posts})
}

// GET /api/blog/:slug
func (h *Handler) GetBySlug(c *gin.Context) {
	post, err := GetPublishedBySlug(c.Request.Context(), h.pool, c.Param("slug"))
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if post == nil {
		httputil.RespondError(c, http.StatusNotFound, "post not found")
		return
	}
	httputil.RespondOK(c, gin.H{"post": post})
}

// ─── Admin ────────────────────────────────────────────────────────────────────

// GET /api/admin/blog
func (h *Handler) AdminList(c *gin.Context) {
	posts, err := GetAll(c.Request.Context(), h.pool)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if posts == nil {
		posts = []Post{}
	}
	httputil.RespondOK(c, gin.H{"posts": posts})
}

// POST /api/admin/blog
func (h *Handler) AdminCreate(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	post, err := Create(c.Request.Context(), h.pool, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not create post")
		return
	}
	httputil.RespondCreated(c, gin.H{"post": post})
}

// PUT /api/admin/blog/:id
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

	post, err := Update(c.Request.Context(), h.pool, id, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not update post")
		return
	}
	if post == nil {
		httputil.RespondError(c, http.StatusNotFound, "post not found")
		return
	}
	httputil.RespondOK(c, gin.H{"post": post})
}

// DELETE /api/admin/blog/:id
func (h *Handler) AdminDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.RespondValidationError(c, "invalid id")
		return
	}

	if err := Delete(c.Request.Context(), h.pool, id); err != nil {
		if err.Error() == "not found" {
			httputil.RespondError(c, http.StatusNotFound, "post not found")
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "could not delete post")
		return
	}
	httputil.RespondOK(c, nil)
}
