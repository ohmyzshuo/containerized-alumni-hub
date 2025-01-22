package faculty

import (
	"net/http"

	"alumni_hub/internal/db"
	"github.com/gin-gonic/gin"
)

func GetFaculties(c *gin.Context) {
	var faculties []Faculty
	db.DB.Find(&faculties)
	c.JSON(http.StatusOK, faculties)
}

func CreateFaculty(c *gin.Context) {
	var faculty Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&faculty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, faculty)
}

func RegisterRoutes(r *gin.Engine) {
	faculties := r.Group("/faculties")
	{
		faculties.GET("/", GetFaculties)
		faculties.POST("/", CreateFaculty)
	}
}
