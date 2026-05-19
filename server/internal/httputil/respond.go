package httputil

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RespondOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"data": data, "error": nil})
}

func RespondCreated(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{"data": data, "error": nil})
}

func RespondError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{"data": nil, "error": message})
}

func RespondValidationError(c *gin.Context, message string) {
	RespondError(c, http.StatusBadRequest, message)
}

func NotFound(c *gin.Context) {
	RespondError(c, http.StatusNotFound, "not found")
}

func IntQuery(c *gin.Context, key string, def int) int {
	v, err := strconv.Atoi(c.Query(key))
	if err != nil || v <= 0 {
		return def
	}
	return v
}
