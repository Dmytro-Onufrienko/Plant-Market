package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/config"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/middleware"
)

func Register(r *gin.RouterGroup, pool *pgxpool.Pool, cfg *config.Config) {
	h := NewHandler(pool, cfg)
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)
	r.GET("/me", middleware.AuthRequired(cfg.JWTSecret), h.Me)
}
