package study

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

func (h *Handler) CreateStudy(c *gin.Context) {
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

	var newStudy Study
	if err := c.ShouldBindJSON(&newStudy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newStudy.AlumniID = uint(alumniID)
	if err := h.svc.CreateStudy(&newStudy); err != nil {
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
		"data":    newStudy,
	})
}

func (h *Handler) GetStudiesByAlumniID(c *gin.Context) {
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

	studies, err := h.svc.GetStudiesByAlumniID(uint(alumniID))
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
		"data":    studies,
	})
}

func (h *Handler) UpdateStudy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid update ID",
			"data":    nil,
		})
		return
	}

	var updatedStudy Study
	if err := c.ShouldBindJSON(&updatedStudy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	study, err := h.svc.UpdateStudy(uint(id), &updatedStudy)
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
		"data":    study,
	})
}

func (h *Handler) DeleteStudy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid delete ID",
			"data":    nil,
		})
		return
	}

	if err := h.svc.DeleteStudy(uint(id)); err != nil {
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

func (h *Handler) GetStudiesByToken(c *gin.Context) {
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

	studies, err := h.svc.GetStudiesByToken(token)
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
		"data":    studies,
	})
}
