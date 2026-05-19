package upload

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/image/webp"

	"github.com/dmytro-onufrienko/monstera-shop/server/internal/config"
	"github.com/dmytro-onufrienko/monstera-shop/server/internal/httputil"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

// POST /api/admin/upload
func (h *Handler) Upload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		httputil.RespondValidationError(c, "image file is required")
		return
	}

	src, err := file.Open()
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not open file")
		return
	}
	defer src.Close()

	img, err := imaging.Decode(src)
	if err != nil {
		httputil.RespondValidationError(c, "invalid image file")
		return
	}

	resized := imaging.Fit(img, 800, 800, imaging.Lanczos)

	filename := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
	savePath := filepath.Join(h.cfg.UploadsDir, "plants", filename)

	if err := imaging.Save(resized, savePath); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "could not save image")
		return
	}

	httputil.RespondCreated(c, gin.H{"url": "/uploads/plants/" + filename})
}
