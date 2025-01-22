package honor

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	honor := r.Group("/alumni/honors")

	honor.POST("/", h.CreateHonor)
	honor.GET("/", h.GetHonorByAlumniID)
	honor.PATCH("/:id", h.UpdateHonor)
	honor.DELETE("/:id", h.DeleteHonor)
	honor.GET("/get", h.GetHonorsByToken)
}
