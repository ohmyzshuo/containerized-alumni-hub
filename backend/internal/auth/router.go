package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	r.POST("/login", handler.Login)
	r.POST("/send-otp", handler.sendOTP)
	r.POST("/verify-otp", handler.verifyOTP)
}
