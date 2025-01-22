package publication

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	publication := r.Group("/alumni/publications")

	publication.POST("/", h.CreatePublication)
	publication.GET("/", h.GetPublicationsByAlumniID)
	publication.PATCH("/:id", h.UpdatePublication)
	publication.DELETE("/:id", h.DeletePublication)
	publication.GET("/get", h.GetPublicationsByToken)
	publication.GET("/statistics", h.StatisticSearch)
}
