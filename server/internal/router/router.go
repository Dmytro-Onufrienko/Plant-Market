package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/admin"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/auth"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/blog"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/config"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/middleware"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/orders"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/plants"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/reviews"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/upload"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/users"
)

func New(cfg *config.Config, pool *pgxpool.Pool) *gin.Engine {
	if cfg.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.Logger(cfg.IsProd()))
	r.Use(middleware.CORS(cfg.AllowedOrigins))

	r.NoRoute(func(c *gin.Context) { httputil.NotFound(c) })

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Static("/uploads", cfg.UploadsDir)

	// ── Auth ──────────────────────────────────────────────────
	auth.Register(r.Group("/api/auth"), pool, cfg)

	// ── Public API ────────────────────────────────────────────
	public := r.Group("/api")

	// ── Authenticated API ─────────────────────────────────────
	authed := r.Group("/api")
	authed.Use(middleware.AuthRequired(cfg.JWTSecret))

	// ── Admin API ─────────────────────────────────────────────
	adminGroup := r.Group("/api/admin")
	adminGroup.Use(middleware.AuthRequired(cfg.JWTSecret), middleware.AdminRequired())

	plants.Register(public, adminGroup, pool)
	blog.Register(public, adminGroup, pool)
	reviews.Register(public, authed, adminGroup, pool)
	orders.Register(authed, adminGroup, pool)
	users.Register(authed, pool)
	admin.Register(adminGroup, pool)
	upload.Register(adminGroup, cfg)

	return r
}
