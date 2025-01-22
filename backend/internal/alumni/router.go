package alumni

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	alumni := r.Group("/alumni")
	r.Use(func(c *gin.Context) {
		c.Next()

		if c.Writer.Status() == 200 {
			if response, ok := c.Get("response"); ok {
				if user, ok := response.(Alumni); ok {
					user.Password = ""
					c.JSON(200, user)
				}
			}
		}
	})

	alumni.GET("/:id", handler.GetAlumnus)
	alumni.GET("/", handler.GetAlumni)
	alumni.POST("/", handler.CreateAlumni)
	alumni.PATCH("/:id", handler.UpdateAlumni)
	alumni.DELETE("/:id", handler.DeleteAlumni)
	alumni.GET("/check", handler.CheckAlumniExistence)
	alumni.POST("/get", handler.GetAlumniByEmail)
	alumni.POST("/getByMatricNo", handler.GetAlumniByMatricNo)
	alumni.GET("/me", handler.GetAlumniByToken)
	alumni.POST("/reset_password/:id", handler.ResetAlumnusPassword)
	alumni.POST("/change_password/:id", handler.ChangeAlumnusPassword)
	alumni.POST("/reminder", handler.SendUpdateReminders)
}
