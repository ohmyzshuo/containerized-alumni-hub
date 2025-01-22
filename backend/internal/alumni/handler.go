package alumni

import (
	"alumni_hub/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	svc ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{svc: service}
}

func (h *Handler) GetAlumnus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid Alumni ID",
		})
		return
	}

	alumnus, err := h.svc.GetAlumnus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error fetching alumnus",
		})
		return
	}

	if alumnus == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Alumnus not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    alumnus,
	})
}

func (h *Handler) GetAlumni(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	if err != nil || pageSize < 1 {
		pageSize = 15
	}

	searchQuery := c.DefaultQuery("search", "")

	alumni, total, err := h.svc.GetAlumni(page, pageSize, searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching alumni",
		})
		return
	}

	for i := range alumni {
		alumni[i].Password = ""
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    alumni,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

func (h *Handler) CreateAlumni(c *gin.Context) {
	var alumni Alumni
	if err := c.ShouldBindJSON(&alumni); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	createdAlumni, err := h.svc.CreateAlumni(&alumni)
	if err != nil {
		if err.Error() == "duplicated matric_no" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    230,
				"data":    nil,
				"message": "Duplicated matric_no",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    nil,
				"message": err.Error(),
			})
		}
		return
	}
	createdAlumni.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"data":    createdAlumni,
		"message": "Alumni created successfully",
	})
}

func (h *Handler) UpdateAlumni(c *gin.Context) {
	var updatedAlumni Alumni
	if err := c.ShouldBindJSON(&updatedAlumni); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid alumni ID",
		})
		return
	}

	updated, err := h.svc.UpdateAlumni(uint(id), &updatedAlumni)
	if err != nil {
		if err.Error() == "duplicated matric_no" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    230,
				"data":    nil,
				"message": "Duplicated matric_no",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    nil,
				"message": err.Error(),
			})
		}
		return
	}

	updated.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    updated,
		"message": "Alumni updated successfully",
	})
}

func (h *Handler) DeleteAlumni(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid alumni ID",
		})
		return
	}

	if err := h.svc.DeleteAlumni(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error hiding alumni",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "Alumni hidden successfully",
	})
}

func (h *Handler) CheckAlumniExistence(c *gin.Context) {
	matricNo := c.Query("matric_no")

	alumni, exists, err := h.svc.CheckAlumniExistence(matricNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error checking alumni existence",
		})
		return
	}
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    alumni,
			"message": "The matric number exists in the database, please log in with the same username",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "The matric number doesn’t exist in the database, please contact staff",
	})
}

func (h *Handler) GetAlumniByEmail(c *gin.Context) {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	email := requestData.Email

	alu, err := h.svc.GetAlumniByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if alu == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alumni not found"})
		return
	}
	alu.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    alu,
		"code":    http.StatusOK,
	})
}

func (h *Handler) GetAlumniByMatricNo(c *gin.Context) {
	var requestData struct {
		MatricNo string `json:"matric_no"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	MatricNo := requestData.MatricNo

	alu, err := h.svc.GetAlumniByMatricNo(MatricNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if alu == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alumni not found"})
		return
	}
	alu.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    alu,
		"code":    http.StatusOK,
	})
}
func (h *Handler) GetAlumniByToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "UserToken is required",
			"data":    nil,
		})
		return
	}
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	user, err := h.svc.GetAlumniByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"code": http.StatusOK,
	})
}

func (h *Handler) ResetAlumnusPassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid alumnus ID",
		})
		return
	}

	alumnus, err := h.svc.GetAlumnusByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(alumnus.MatricNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Failed to hash password",
		})
		return
	}

	updatedAlumnus := &Alumni{
		Password: hashedPassword,
	}

	updated, err := h.svc.UpdateAlumni(uint(id), updatedAlumnus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	updated.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    updated,
		"message": "Password reset successfully",
	})
}

func (h *Handler) ChangeAlumnusPassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid alumnus ID",
		})
		return
	}

	var requestBody struct {
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid request body",
		})
		return
	}

	if requestBody.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Password cannot be empty",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Failed to hash password",
		})
		return
	}

	updatedAlumnus := &Alumni{
		Password: hashedPassword,
	}

	updated, err := h.svc.UpdateAlumni(uint(id), updatedAlumnus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	updated.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    updated,
		"message": "Password changed successfully",
	})
}

func (h *Handler) SendUpdateReminders(c *gin.Context) {
	err := h.svc.SendUpdateReminders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to send update reminders: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Update reminders sent successfully",
	})
}
