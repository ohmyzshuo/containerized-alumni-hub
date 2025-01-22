package pubExporter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Handler struct holds the service instance
type Handler struct {
	svc *Service
}

// NewHandler creates a new instance of Handler
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// ExportGraduatesHandler is a handler function for exporting graduates' information by convocation year
func (h *Handler) ExportGraduatesHandler(c *gin.Context) {
	// Extract convocation year from query params
	convocationYear := c.DefaultQuery("convocation_year", "")
	if convocationYear == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Convocation year is required"})
		return
	}

	// Convert to integer
	convocationYearInt, err := strconv.Atoi(convocationYear)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid convocation year format"})
		return
	}

	// Generate the graduate report
	file, err := h.svc.ExportGraduatesByYear(convocationYearInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error generating report: %v", err)})
		return
	}

	// Save the file to a temporary location or return directly to user
	fileName := fmt.Sprintf("graduates_report_%d.xlsx", convocationYearInt)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// Serve the file as a download
	if err := file.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing file to response"})
		return
	}

	log.Printf("Report successfully generated and sent for convocation year %d", convocationYearInt)
}
