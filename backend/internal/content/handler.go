package content

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jdkato/prose/v2"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// extractTags analyzes the input text and extracts keyword tags
func extractTags(text string) ([]string, error) {
	// Create a new document with the input text
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}

	// Use a map to store unique tags
	tagSet := make(map[string]struct{})

	// Extract named entities (like "Go" or "Python")
	for _, ent := range doc.Entities() {
		tagSet[strings.ToLower(ent.Text)] = struct{}{}
	}

	// Extract nouns and proper nouns as potential tags
	for _, tok := range doc.Tokens() {
		if tok.Tag == "NN" || tok.Tag == "NNP" { // NN: Noun, NNP: Proper Noun
			tagSet[strings.ToLower(tok.Text)] = struct{}{}
		}
	}

	// Convert the map keys to a slice
	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}
	log.Println("@@@@@@@@@@", tags)
	return tags, nil
}

type CreateContentRequest struct {
	Title            string    `form:"title"`
	Description      string    `form:"description"`
	Contact          string    `form:"contact"`
	CreatedBy        uint      `form:"created_by"`
	UpdatedBy        uint      `form:"updated_by"`
	IsHidden         bool      `form:"is_hidden"`
	IsPinned         bool      `form:"is_pinned"`
	CreatedAt        time.Time `form:"created_at"`
	UpdatedAt        time.Time `form:"updated_at"`
	Venue            string    `form:"venue"`
	FacultyID        uint      `form:"faculty_id"`
	StartTime        time.Time `form:"start_time"`
	EndTime          time.Time `form:"end_time"`
	ContentType      uint      `form:"content_type"`
	NumOfLikes       uint      `form:"num_of_likes"`
	ParticipantQuota uint      `form:"participant_quota"`
}

func (h *Handler) CreateContent(c *gin.Context) {
	var req CreateContentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	content := Content{
		Title:            req.Title,
		Description:      req.Description,
		Contact:          req.Contact,
		CreatedBy:        req.CreatedBy,
		UpdatedBy:        req.UpdatedBy,
		IsHidden:         req.IsHidden,
		IsPinned:         req.IsPinned,
		CreatedAt:        req.CreatedAt,
		UpdatedAt:        req.UpdatedAt,
		Venue:            req.Venue,
		FacultyID:        req.FacultyID,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		ContentType:      req.ContentType,
		NumOfLikes:       req.NumOfLikes,
		ParticipantQuota: req.ParticipantQuota,
	}

	// 设置默认用户 ID
	if content.CreatedBy == 0 {
		content.CreatedBy = 2 // 设置为一个有效的用户 ID
	}
	if content.UpdatedBy == 0 {
		content.UpdatedBy = 2 // 设置为一个有效的用户 ID
	}

	// 提取标签
	tags, err := extractTags(req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 处理文件上传
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid form data",
			"data":    nil,
		})
		return
	}

	files := form.File["attachments"]
	var attachments []Attachment
	for _, file := range files {
		// 保存文件到本地
		filePath := filepath.Join("uploads/content", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		// 添加附件信息
		attachments = append(attachments, Attachment{
			OriginalName:   file.Filename,
			AttachmentPath: filePath,
		})
	}

	if err := h.svc.CreateContent(&content, attachments, tags); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 将附件添加到内容对象中
	content.Attachments = attachments

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    content,
	})
}

func (h *Handler) GetContentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid content ID",
			"data":    nil,
		})
		return
	}

	content, attachments, err := h.svc.GetContentByID(uint(id))
	if err != nil {
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
		"data":    gin.H{"content": content, "attachments": attachments},
	})
}

func (h *Handler) ListDefault(c *gin.Context) {
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

	search := c.DefaultQuery("search", "")

	contents, total, err := h.svc.ListDefault(page, pageSize, search)
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
		"data":    contents,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
func (h *Handler) ListPinned(c *gin.Context) {
	contents, err := h.svc.ListPinned()
	if err != nil {
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
		"data":    contents,
	})
}

type updateContentRequest struct {
	Title            string    `form:"title"`
	Description      string    `form:"description"`
	Contact          string    `form:"contact"`
	CreatedBy        uint      `form:"created_by"`
	UpdatedBy        uint      `form:"updated_by"`
	IsHidden         bool      `form:"is_hidden"`
	IsPinned         bool      `form:"is_pinned"`
	Venue            string    `form:"venue"`
	FacultyID        uint      `form:"faculty_id"`
	StartTime        time.Time `form:"start_time"`
	EndTime          time.Time `form:"end_time"`
	ContentType      uint      `form:"content_type"`
	NumOfLikes       uint      `form:"num_of_likes"`
	ParticipantQuota uint      `form:"participant_quota"`
}

func (h *Handler) UpdateContent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid content ID",
			"data":    nil,
		})
		return
	}

	var req updateContentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid form data",
			"data":    nil,
		})
		return
	}

	// 处理文件上传
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid form data",
			"data":    nil,
		})
		return
	}

	files := form.File["attachments"]
	var attachments []Attachment
	for _, file := range files {
		// 保存文件到本地
		filePath := filepath.Join("uploads/content", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		// 添加附件信息
		attachments = append(attachments, Attachment{
			OriginalName:   file.Filename,
			AttachmentPath: filePath,
		})
	}

	// 创建 Content 实例
	content := Content{
		Title:            req.Title,
		Description:      req.Description,
		Contact:          req.Contact,
		CreatedBy:        req.CreatedBy,
		UpdatedBy:        req.UpdatedBy,
		IsHidden:         req.IsHidden,
		IsPinned:         req.IsPinned,
		Venue:            req.Venue,
		FacultyID:        req.FacultyID,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		ContentType:      req.ContentType,
		NumOfLikes:       req.NumOfLikes,
		ParticipantQuota: req.ParticipantQuota,
	}

	updatedContent, err := h.svc.UpdateContent(uint(id), &content, attachments)
	updatedContent.Attachments = attachments

	if err != nil {
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
		"data":    updatedContent,
	})
}

func (h *Handler) DeleteContent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid content ID",
			"data":    nil,
		})
		return
	}

	if err := h.svc.DeleteContent(uint(id)); err != nil {
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
		"data":    nil,
	})
}

type AlumniIDsRequest struct {
	AlumniIDs []uint `json:"alumni_ids"`
}

func (h *Handler) SendContentEmail(c *gin.Context) {
	contentID, err := strconv.Atoi(c.Query("content_id"))
	if err != nil || contentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid content ID",
			"data":    nil,
		})
		return
	}

	var request AlumniIDsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	if err := h.svc.SendContentEmail(uint(contentID), request.AlumniIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Emails sent successfully",
		"data":    nil,
	})
}

func (h *Handler) ListAllByPreferences(c *gin.Context) {
	alumniIDStr := c.Param("alumni_id")
	alumniID, err := strconv.ParseUint(alumniIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumni ID"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	contents, total, err := h.svc.ListAllByPreferences(uint(alumniID), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    contents,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

func (h *Handler) LikeContent(c *gin.Context) {
	alumniIDStr := c.Query("alumni_id")
	contentIDStr := c.Query("content_id")

	alumniID, err := strconv.ParseUint(alumniIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumni ID"})
		return
	}

	contentID, err := strconv.ParseUint(contentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content ID"})
		return
	}

	if err := h.svc.LikeContent(uint(alumniID), uint(contentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content liked successfully"})
}

func (h *Handler) DislikeContent(c *gin.Context) {
	alumniIDStr := c.Query("alumni_id")
	contentIDStr := c.Query("content_id")

	alumniID, err := strconv.ParseUint(alumniIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumni ID"})
		return
	}

	contentID, err := strconv.ParseUint(contentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content ID"})
		return
	}

	if err := h.svc.DislikeContent(uint(alumniID), uint(contentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Content disliked successfully"})
}

/// event management only

type ChangeStatusRequest struct {
	AlumniID  uint   `json:"alumni_id" binding:"required"`
	ContentID uint   `json:"content_id" binding:"required"`
	Comment   string `json:"comment"`
	ToStatus  uint   `json:"to_status"`
}

func (h *Handler) ChangeParticipantStatus(c *gin.Context) {
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := h.svc.ChangeAlumniStatusForEvent(req.AlumniID, req.ContentID, req.ToStatus, req.Comment); err != nil {
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
		"data":    nil,
	})
}

func (h *Handler) ChangeStatusByToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "token is required",
			"data":    nil,
		})
		return
	}

	claims, err := ParseAndValidateToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err = h.svc.ChangeAlumniStatusForEvent(claims.AlumniID, claims.ContentID, claims.Status, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlResponse))
}

const htmlResponse = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Response Recorded</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f9;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .container {
            background-color: #fff;
            padding: 20px 40px;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            text-align: center;
        }

        h2 {
            color: #333;
        }

        p {
            color: #666;
            line-height: 1.6;
        }

        .signature {
            margin-top: 20px;
            font-style: italic;
            color: #888;
        }
    </style>
</head>
<body>
<div class="container">
    <h2>Dear Alumni,</h2>
    <p>Your response has been recorded.</p>
    <p>Thank you.</p>
    <div class="signature">AlumniHub @ University of Malaya</div>
</div>
</body>
</html>
`

func generateEmailLinks(alumniID uint, contentID uint) (string, string, string, error) {
	baseURL := os.Getenv("BASE_URL")
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}
	notInterestedToken, err := GenerateSignedToken(alumniID, contentID, 1)
	if err != nil {
		return "", "", "", err
	}
	notInterestedLink := baseURL + "contents/respond?token=" + notInterestedToken

	interestedButCannotJoinToken, err := GenerateSignedToken(alumniID, contentID, 2)
	if err != nil {
		return "", "", "", err
	}
	interestedButCannotJoinLink := baseURL + "contents/respond?token=" + interestedButCannotJoinToken

	registeredToken, err := GenerateSignedToken(alumniID, contentID, 3)
	if err != nil {
		return "", "", "", err
	}
	registeredLink := baseURL + "contents/respond?token=" + registeredToken
	log.Printf("Generated links for alumni %d: not interested: %s, interested: %s, register: %s",
		alumniID, notInterestedLink, interestedButCannotJoinLink, registeredLink)
	return notInterestedLink, interestedButCannotJoinLink, registeredLink, nil
}

// RegisterForEvent change status to 3
func (h *Handler) RegisterForEvent(c *gin.Context) {
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.svc.ChangeAlumniStatusForEvent(req.AlumniID, req.ContentID, 3, req.Comment); err != nil {
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
		"data":    nil,
	})
}

// CancelEventRegistration change status to 2
func (h *Handler) CancelEventRegistration(c *gin.Context) {
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.svc.ChangeAlumniStatusForEvent(req.AlumniID, req.ContentID, 2, req.Comment); err != nil {
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
		"data":    nil,
	})
}

// ShowNoInterestInEvent change status to 1
func (h *Handler) ShowNoInterestInEvent(c *gin.Context) {
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.svc.ChangeAlumniStatusForEvent(req.AlumniID, req.ContentID, 1, req.Comment); err != nil {
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
		"data":    nil,
	})
}

// AttendEvent change status to 4
func (h *Handler) AttendEvent(c *gin.Context) {
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.svc.ChangeAlumniStatusForEvent(req.AlumniID, req.ContentID, 4, req.Comment); err != nil {
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
		"data":    nil,
	})
}

// RecordEventAbsence change status to 5
func (h *Handler) RecordEventAbsence(c *gin.Context) {
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.svc.ChangeAlumniStatusForEvent(req.AlumniID, req.ContentID, 5, req.Comment); err != nil {
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
		"data":    nil,
	})
}

func (h *Handler) GetEventParticipants(c *gin.Context) {
	contentIDStr := c.Query("content_id")
	statusStr := c.Query("status")

	if contentIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "content_id is required",
			"data":    nil,
		})
		return
	}

	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid content_id",
			"data":    nil,
		})
		return
	}

	var status *int
	if statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "invalid status",
				"data":    nil,
			})
			return
		}
		status = &statusInt
	}

	participants, err := h.svc.GetEventParticipants(contentID, status)
	if err != nil {
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
		"data":    participants,
	})
}
func (h *Handler) GetAlumnusParticipation(c *gin.Context) {
	// 获取 query 参数
	alumniIDStr := c.Query("alumni_id")
	statusStr := c.Query("status")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	if alumniIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "alumni_id is required",
			"data":    nil,
		})
		return
	}

	alumniID, err := strconv.Atoi(alumniIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid alumni_id",
			"data":    nil,
		})
		return
	}

	var status *int
	if statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "invalid status",
				"data":    nil,
			})
			return
		}
		status = &statusInt
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid page number",
			"data":    nil,
		})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid page size",
			"data":    nil,
		})
		return
	}

	contents, total, err := h.svc.GetAlumnusParticipation(alumniID, status, page, pageSize)
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
		"data":    contents,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

type NewsletterRequest struct {
	ContentIDs []uint `json:"content_ids"`
	AlumniIDs  []uint `json:"alumni_ids"`
}

func (h *Handler) SendNewsletter(c *gin.Context) {
	var req NewsletterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	if len(req.ContentIDs) == 0 || len(req.AlumniIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Content IDs and Alumni IDs cannot be empty",
			"data":    nil,
		})
		return
	}

	if err := h.svc.SendNewsletter(req.ContentIDs, req.AlumniIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Newsletter sent successfully",
		"data":    nil,
	})
}
