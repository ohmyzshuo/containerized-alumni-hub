package staff

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	r.Use(func(c *gin.Context) {
		c.Next()

		if c.Writer.Status() == 200 {
			if response, ok := c.Get("response"); ok {
				if user, ok := response.(Staff); ok {
					user.Password = ""
					c.JSON(200, user)
				}
			}
		}
	})
	r.GET("/staff", handler.GetStaffs)
	r.POST("/staff", handler.CreateStaff)
	r.PATCH("/staff/:id", handler.UpdateStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)
	r.GET("/staff/me", handler.GetStaffByToken)
	r.POST("/staff/reset_password/:id", handler.ResetStaffPassword)
	r.POST("/staff/change_password/:id", handler.ChangeStaffPassword)
}
