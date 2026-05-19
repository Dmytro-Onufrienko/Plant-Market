package upload

import (
	"github.com/gin-gonic/gin"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/config"
)

func Register(admin *gin.RouterGroup, cfg *config.Config) {
	h := NewHandler(cfg)
	admin.POST("/upload", h.Upload)
}
