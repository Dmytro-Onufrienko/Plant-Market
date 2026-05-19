package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/middleware"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/services"
)

type Handler struct {
	pool *pgxpool.Pool
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

// PUT /api/users/email
func (h *Handler) UpdateEmail(c *gin.Context) {
	userID := c.GetString(middleware.CtxUserID)

	var req UpdateEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	user, err := GetByID(c.Request.Context(), h.pool, userID)
	if err != nil || user == nil {
		httputil.RespondError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if !services.CheckPassword(user.PasswordHash, req.CurrentPassword) {
		httputil.RespondError(c, http.StatusForbidden, "incorrect password")
		return
	}

	existing, err := GetByEmail(c.Request.Context(), h.pool, req.NewEmail)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if existing != nil {
		httputil.RespondError(c, http.StatusConflict, "email already in use")
		return
	}

	if err := UpdateEmail(c.Request.Context(), h.pool, userID, req.NewEmail); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not update email")
		return
	}
	httputil.RespondOK(c, nil)
}

// PUT /api/users/password
func (h *Handler) UpdatePassword(c *gin.Context) {
	userID := c.GetString(middleware.CtxUserID)

	var req UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		httputil.RespondValidationError(c, "passwords do not match")
		return
	}

	user, err := GetByID(c.Request.Context(), h.pool, userID)
	if err != nil || user == nil {
		httputil.RespondError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	if !services.CheckPassword(user.PasswordHash, req.CurrentPassword) {
		httputil.RespondError(c, http.StatusForbidden, "incorrect password")
		return
	}

	newHash, err := services.HashPassword(req.NewPassword)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not hash password")
		return
	}

	if err := UpdatePassword(c.Request.Context(), h.pool, userID, newHash); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not update password")
		return
	}
	httputil.RespondOK(c, nil)
}
