package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/config"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/middleware"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/services"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/users"
)

type Handler struct {
	pool *pgxpool.Pool
	cfg  *config.Config
}

func NewHandler(pool *pgxpool.Pool, cfg *config.Config) *Handler {
	return &Handler{pool: pool, cfg: cfg}
}

// POST /api/auth/register
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	existing, err := users.GetByEmail(c.Request.Context(), h.pool, req.Email)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if existing != nil {
		httputil.RespondError(c, http.StatusConflict, "email already in use")
		return
	}

	hash, err := services.HashPassword(req.Password)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not hash password")
		return
	}

	user, err := users.Create(c.Request.Context(), h.pool, uuid.NewString(), req.Email, hash, req.Name)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not create user")
		return
	}

	h.setAuthCookie(c, user.ID, user.Role)
	httputil.RespondCreated(c, users.ToPublic(user))
}

// POST /api/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondValidationError(c, err.Error())
		return
	}

	user, err := users.GetByEmail(c.Request.Context(), h.pool, req.Email)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "database error")
		return
	}
	if user == nil || !services.CheckPassword(user.PasswordHash, req.Password) {
		httputil.RespondError(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	h.setAuthCookie(c, user.ID, user.Role)
	httputil.RespondOK(c, gin.H{"user": users.ToPublic(user)})
}

// POST /api/auth/logout
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie(middleware.CookieName, "", -1, "/", "", h.cfg.IsProd(), true)
	httputil.RespondOK(c, nil)
}

// GET /api/auth/me
func (h *Handler) Me(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxUserID)

	user, err := users.GetByID(c.Request.Context(), h.pool, userID.(string))
	if err != nil || user == nil {
		httputil.RespondError(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	httputil.RespondOK(c, users.ToPublic(user))
}

func (h *Handler) setAuthCookie(c *gin.Context, userID, role string) {
	token, _ := services.GenerateToken(userID, role, h.cfg.JWTSecret, h.cfg.JWTExpiry)
	maxAge := int(h.cfg.JWTExpiry / time.Second)
	c.SetCookie(middleware.CookieName, token, maxAge, "/", "", h.cfg.IsProd(), true)
}
