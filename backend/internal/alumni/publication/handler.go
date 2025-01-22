package publication

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

func (h *Handler) CreatePublication(c *gin.Context) {
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

	var newPublication Publication
	if err := c.ShouldBindJSON(&newPublication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newPublication.AlumniID = uint(alumniID)
	if err := h.svc.CreatePublication(&newPublication); err != nil {
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
		"data":    newPublication,
	})
}

func (h *Handler) GetPublicationsByAlumniID(c *gin.Context) {
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

	publications, err := h.svc.GetPublicationsByAlumniID(uint(alumniID))
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
		"data":    publications,
	})
}

func (h *Handler) UpdatePublication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid update ID",
			"data":    nil,
		})
		return
	}

	var updatedPublication Publication
	if err := c.ShouldBindJSON(&updatedPublication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	publication, err := h.svc.UpdatePublication(uint(id), &updatedPublication)
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
		"data":    publication,
	})
}

func (h *Handler) DeletePublication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid delete ID",
			"data":    nil,
		})
		return
	}

	if err := h.svc.DeletePublication(uint(id)); err != nil {
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

func (h *Handler) GetPublicationsByToken(c *gin.Context) {
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

	honors, err := h.svc.GetPublicationsByToken(token)
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

func (h *Handler) StatisticSearch(c *gin.Context) {
	// Parse optional query parameters
	acceptedDateStart := c.Query("accepted_date_start")
	acceptedDateEnd := c.Query("accepted_date_end")
	publicationTypes := c.QueryArray("publication_type")
	statuses := c.QueryArray("status")
	quartiles := c.QueryArray("quartile")
	search := c.Query("search")

	// Parse pagination parameters
	pageSizeStr := c.Query("page_size")
	currentPageStr := c.Query("current_page")

	// Set default values for pagination if not provided
	pageSize := 10
	currentPage := 1

	// Convert pagination parameters to integers
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil {
			pageSize = ps
		}
	}
	if currentPageStr != "" {
		if cp, err := strconv.Atoi(currentPageStr); err == nil {
			currentPage = cp
		}
	}

	// Call the service to get statistics
	publications, meta, statistics, err := h.svc.GetStatistics(acceptedDateStart, acceptedDateEnd, publicationTypes, statuses, quartiles, search, pageSize, currentPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// Return the results
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    publications,
		"meta": gin.H{
			"pagination": meta.Pagination,
			"statistics": statistics,
		},
	})
}
