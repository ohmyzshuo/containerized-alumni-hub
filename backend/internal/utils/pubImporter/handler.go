package pubImporter

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ImportPublications(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "file is required",
		})
		return
	}

	timestamp := time.Now().Format("20060102150405")
	fileName := timestamp + "_" + file.Filename
	filePath := filepath.Join("./uploads/import", fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "failed to save file",
		})
		return
	}

	importedCount, invalidCount, skippedRows, err := h.svc.ImportAlumuiPublicationsFromExcel(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Alumni Publications imported successfully",
		"data": gin.H{
			"imported_count": importedCount,
			"invalid_count":  invalidCount,
			"skipped_rows":   skippedRows,
		},
	})
}
