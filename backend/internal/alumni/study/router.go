package study

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	study := r.Group("/alumni/studies")

	study.POST("/", h.CreateStudy)
	study.GET("/", h.GetStudiesByAlumniID)
	study.PATCH("/:id", h.UpdateStudy)
	study.DELETE("/:id", h.DeleteStudy)
	study.GET("/get", h.GetStudiesByToken)
}
