package auth

import (
	"alumni_hub/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB      *gorm.DB
	Service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) Login(c *gin.Context) {
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // "alumni" or "staff"
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Invalid request",
		})
		return
	}

	userID, role, err := h.Service.AuthenticateUser(login.Username, login.Password, login.Role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": err.Error(),
		})
		return
	}

	token, err := h.Service.GenerateToken(userID, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Failed to generate token",
		})
		return
	}

	// Store the token in the database
	if err := h.Service.StoreToken(userID, login.Role, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Failed to store token!" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"token": token,
		},
		"message": "Success",
	})
}

func (h *Handler) sendOTP(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Email is required",
		})
		return
	}

	otp := utils.GenerateOTP(email)
	err := utils.SendEmail(email, "AlumniHub OTP Code", "Your OTP code is: "+otp, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Failed to send email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OTP sent successfully",
	})
}

func (h *Handler) verifyOTP(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")

	if email == "" || otp == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Email and OTP are required",
		})
		return
	}

	if utils.VerifyOTP(email, otp) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OTP verified successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    233,
			"message": "Failed to verify OTP",
		})
	}
}
