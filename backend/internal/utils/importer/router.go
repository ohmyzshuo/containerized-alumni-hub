package importer

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	importer := r.Group("/import")

	importer.POST("/", handler.ImportAlumni)

}
