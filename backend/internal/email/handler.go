package email

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateEmail(c *gin.Context) {
	to := c.PostForm("to")
	from := os.Getenv("EMAIL_FROM")
	subject := c.PostForm("subject")
	body := c.PostForm("body")

	// Validate required fields
	if to == "" || subject == "" || body == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Missing required fields",
			"data":    nil,
		})
		return
	}

	// Handle file uploads
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Error retrieving form data",
			"data":    nil,
		})
		return
	}

	files := form.File["attachments"]
	var attachments []Attachment

	for _, file := range files {
		// Save the file to a specific directory
		filePath := filepath.Join("uploads/email", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Failed to save file",
				"data":    nil,
			})
			return
		}

		// Append to attachments
		attachments = append(attachments, Attachment{
			OriginalName:   file.Filename,
			AttachmentPath: filePath,
		})
	}

	email := Email{
		To:          to,
		From:        from,
		Subject:     subject,
		Body:        body,
		Attachments: attachments,
	}

	if err := h.svc.CreateEmail(&email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    email,
	})
}

func (h *Handler) GetEmails(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid page number",
			"data":    nil,
		})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid page size",
			"data":    nil,
		})
		return
	}

	emails, total, err := h.svc.GetEmails(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    emails,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
