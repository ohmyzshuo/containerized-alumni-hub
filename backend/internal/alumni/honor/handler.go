package honor

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

func (h *Handler) CreateHonor(c *gin.Context) {
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

	var newHonor Honor
	if err := c.ShouldBindJSON(&newHonor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newHonor.AlumniID = uint(alumniID)
	if err := h.svc.CreateHonor(&newHonor); err != nil {
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
		"data":    newHonor,
	})
}

func (h *Handler) GetHonorByAlumniID(c *gin.Context) {
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

	honors, err := h.svc.GetHonorByAlumniID(uint(alumniID))
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
		"data":    honors,
	})
}

func (h *Handler) UpdateHonor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid update ID",
			"data":    nil,
		})
		return
	}

	var updatedHonor Honor
	if err := c.ShouldBindJSON(&updatedHonor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	honor, err := h.svc.UpdateHonor(uint(id), &updatedHonor)
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
		"data":    honor,
	})
}

func (h *Handler) DeleteHonor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid delete ID",
			"data":    nil,
		})
		return
	}

	if err := h.svc.DeleteHonor(uint(id)); err != nil {
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

func (h *Handler) GetHonorsByToken(c *gin.Context) {
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

	honors, err := h.svc.GetHonorsByToken(token)
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
		"data":    honors,
	})
}