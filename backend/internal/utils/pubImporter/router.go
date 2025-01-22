package pubImporter

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	pubImporter := r.Group("/pub-import")

	pubImporter.POST("/", handler.ImportPublications)

}
