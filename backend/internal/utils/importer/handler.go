package importer

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

func (h *Handler) ImportAlumni(c *gin.Context) {
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
	readCount, validCount, exsitNo, sameNo, needUpdateNo, updatedNo, newNo, createNo, fileIssue, failedQueryRows, failedUpdateRows, failedCreateRows, skippedRows, headIssue, err := h.svc.ImportAlumniStudiesFromExcel(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Alumni data imported successfully",
		"data": gin.H{
			"read_count":         readCount,
			"valid_count":        validCount,
			"exsit_no":           exsitNo,
			"same_no":            sameNo,
			"need_update_no":     needUpdateNo,
			"updated_no":         updatedNo,
			"new_no":             newNo,
			"create_no":          createNo,
			"file_issue":         fileIssue,
			"failed_query_rows":  failedQueryRows,
			"failed_update_rows": failedUpdateRows,
			"failed_create_rows": failedCreateRows,
			"skipped_rows":       skippedRows,
			"head_issue":         headIssue,
		},
	})
}
