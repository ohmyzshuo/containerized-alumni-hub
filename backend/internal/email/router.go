package email

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	email := r.Group("/emails")
	email.GET("/", h.GetEmails)
	email.POST("/", h.CreateEmail)

}
