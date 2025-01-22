package staff

import (
	"alumni_hub/internal/utils"
	customErrors "alumni_hub/internal/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Handler struct {
	svc *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{svc: service}
}

func (h *Handler) GetStaffs(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "15"))
	if err != nil || pageSize < 1 {
		pageSize = 15
	}

	searchQuery := c.DefaultQuery("search", "")

	filters := map[string]string{
		"is_super_admin": c.DefaultQuery("is_super_admin", ""),
		"faculty_id":     c.DefaultQuery("faculty_id", ""),
	}

	staff, total, err := h.svc.GetStaffs(page, pageSize, searchQuery, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching staff",
		})
		return
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	for i := range staff {
		staff[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    staff,
		"meta": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

func (h *Handler) CreateStaff(c *gin.Context) {
	var staff Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	// Email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(staff.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    customErrors.EmailFormatIncorrectError,
			"data":    nil,
			"message": "Email format is incorrect",
		})
		return
	}

	createdStaff, err := h.svc.CreateStaff(&staff)
	if err != nil {
		if err.Error() == "duplicated username" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    230,
				"data":    nil,
				"message": "Duplicated username",
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
	createdStaff.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"data":    createdStaff,
		"message": "Staff created successfully",
	})
}

func (h *Handler) UpdateStaff(c *gin.Context) {
	var updatedStaff Staff
	if err := c.ShouldBindJSON(&updatedStaff); err != nil {
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
			"message": "Invalid staff ID",
		})
		return
	}

	updated, err := h.svc.UpdateStaff(uint(id), &updatedStaff)
	if err != nil {
		if err.Error() == "duplicated username" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    230,
				"data":    nil,
				"message": "Duplicated username",
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
		"message": "Staff updated successfully",
	})
}

func (h *Handler) DeleteStaff(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid staff ID",
		})
		return
	}

	if err := h.svc.DeleteStaff(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Error hiding staff",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    nil,
		"message": "Staff hidden successfully",
	})
}
func (h *Handler) GetStaffByToken(c *gin.Context) {
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

	user, err := h.svc.GetStaffByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": user})
}
func (h *Handler) ResetStaffPassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid staff ID",
		})
		return
	}

	staff, err := h.svc.GetStaffByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(staff.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "Failed to hash password",
		})
		return
	}

	updatedStaff := &Staff{
		Password: hashedPassword,
	}

	updated, err := h.svc.UpdateStaff(uint(id), updatedStaff)
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

func (h *Handler) ChangeStaffPassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid staff ID",
		})
		return
	}

	var requestBody struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"data":    nil,
			"message": "Invalid request body",
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

	updatedStaff := &Staff{
		Password: hashedPassword,
	}

	updated, err := h.svc.UpdateStaff(uint(id), updatedStaff)
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
