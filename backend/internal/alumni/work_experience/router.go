package work_experience

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	work := r.Group("/alumni/works")

	work.POST("/", h.CreateWorkExperience)
	work.GET("/", h.GetWorkExperiencesByAlumniID)
	work.PATCH("/:id", h.UpdateWorkExperience)
	work.DELETE("/:id", h.DeleteWorkExperience)
	work.GET("/get", h.GetWorksByToken)

}
