package content

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	content := r.Group("/contents")
	content.GET("/like", h.LikeContent)
	content.GET("/dislike", h.DislikeContent)
	content.POST("/", h.CreateContent)
	content.POST("/email", h.SendContentEmail)
	content.GET("/:id", h.GetContentByID)
	content.GET("/default", h.ListDefault)
	content.GET("/pinned", h.ListPinned)
	content.PATCH("/:id", h.UpdateContent)
	content.DELETE("/:id", h.DeleteContent)
	content.GET("/ai", h.ListAllByPreferences)

	// event management
	content.POST("/attend", h.AttendEvent)
	content.POST("/cancel", h.CancelEventRegistration)
	content.POST("/absent", h.RecordEventAbsence)
	content.POST("/register", h.RegisterForEvent)
	content.POST("/uninterested", h.ShowNoInterestInEvent)
	content.POST("/change", h.ChangeParticipantStatus)
	content.GET("/participants", h.GetEventParticipants)
	content.GET("/participation", h.GetAlumnusParticipation)
	content.POST("/newsletter", h.SendNewsletter)
	content.GET("/respond", h.ChangeStatusByToken)
}
