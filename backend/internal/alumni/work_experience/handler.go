package work_experience

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateWorkExperience(c *gin.Context) {
	alumniIDStr := c.Query("alumni_id")
	alumniID, err := strconv.Atoi(alumniIDStr)
	if err != nil || alumniID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid alumni_id",
			"data":    nil,
		})
		return
	}

	var newWorkExperience WorkExperience
	if err := c.ShouldBindJSON(&newWorkExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newWorkExperience.AlumniID = uint(alumniID)
	if err := h.svc.CreateWorkExperience(&newWorkExperience); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Success",
		"data":    newWorkExperience,
	})
}

func (h *Handler) GetWorkExperiencesByAlumniID(c *gin.Context) {
	alumniIDStr := c.Query("alumni_id")
	alumniID, err := strconv.Atoi(alumniIDStr)
	if err != nil || alumniID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid alumni_id",
			"data":    nil,
		})
		return
	}

	workExperiences, err := h.svc.GetWorkExperiencesByAlumniID(uint(alumniID))
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
		"data":    workExperiences,
	})
}

func (h *Handler) UpdateWorkExperience(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid update ID",
			"data":    nil,
		})
		return
	}

	var updatedWorkExperience WorkExperience
	if err := c.ShouldBindJSON(&updatedWorkExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	workExperience, err := h.svc.UpdateWorkExperience(uint(id), &updatedWorkExperience)
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
		"data":    workExperience,
	})
}

func (h *Handler) DeleteWorkExperience(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid delete ID",
			"data":    nil,
		})
		return
	}

	if err := h.svc.DeleteWorkExperience(uint(id)); err != nil {
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

func (h *Handler) GetWorksByToken(c *gin.Context) {
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

	works, err := h.svc.GetWorksByToken(token)
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
		"data":    works,
	})
}
