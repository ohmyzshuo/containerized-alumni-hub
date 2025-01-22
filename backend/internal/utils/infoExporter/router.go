package infoExporter

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the publications exporter
func RegisterRoutes(r *gin.Engine, handler *Handler) {
	// Route to export publications based on the convocation year
	r.GET("/export-info", handler.ExportInfoHandler)
}
