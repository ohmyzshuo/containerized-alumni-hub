package main

import (
	"alumni_hub/internal/alumni"
	"alumni_hub/internal/alumni/honor"
	"alumni_hub/internal/alumni/publication"
	"alumni_hub/internal/alumni/study"
	"alumni_hub/internal/alumni/work_experience"
	"alumni_hub/internal/auth"
	"alumni_hub/internal/content"
	"alumni_hub/internal/db"
	"alumni_hub/internal/email"
	"alumni_hub/internal/staff"
	"alumni_hub/internal/utils/importer"
	"alumni_hub/internal/utils/infoExporter"
	"alumni_hub/internal/utils/pubExporter"
	"alumni_hub/internal/utils/pubImporter"
	"alumni_hub/internal/utils/workExporter"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := gin.New()
	r.Use(Cors()) // Apply CORS middleware
	r.RedirectTrailingSlash = false

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Init()

	database := db.DB

	// 在初始化数据库后
	if database != nil {
		log.Println("Database connection successful")
		// 可以尝试执行一个简单的查询
		result := database.Raw("SELECT 1")
		if result.Error != nil {
			log.Printf("Database test query failed: %v", result.Error)
		} else {
			log.Println("Database test query successful")
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	log.Printf("Connecting to database with DSN: %s", dsn)
	// init service and handlers, and register routes
	// auth routes
	authService := auth.NewService(database)
	authHandler := auth.NewHandler(authService)
	auth.RegisterRoutes(r, authHandler)

	// staff routes
	staffService := staff.NewService(database)
	staffHandler := staff.NewHandler(staffService)
	staff.RegisterRoutes(r, staffHandler)

	//// alumni routes
	//alumniService := alumni.NewService(database)
	//alumniHandler := alumni.NewHandler(alumniService)
	//alumni.RegisterRoutes(r, alumniHandler)

	// honor routes
	honorSvc := honor.NewService(database)
	honorHdl := honor.NewHandler(honorSvc)
	honor.RegisterRoutes(r, honorHdl)

	// work experience routes
	workSvc := work_experience.NewService(database)
	workHdl := work_experience.NewHandler(workSvc)
	work_experience.RegisterRoutes(r, workHdl)

	// study experience routes
	studySvc := study.NewService(database)
	studyHdl := study.NewHandler(studySvc)
	study.RegisterRoutes(r, studyHdl)

	// publication routes
	publicationService := publication.NewService(database)
	publicationHandler := publication.NewHandler(publicationService)
	publication.RegisterRoutes(r, publicationHandler)

	// content routes
	contentService := content.NewService(database)
	contentHandler := content.NewHandler(contentService)
	content.RegisterRoutes(r, contentHandler)

	// importer routes
	importerSvc := importer.NewService(database)
	importerHdl := importer.NewHandler(importerSvc)
	importer.RegisterRoutes(r, importerHdl)

	// pub importer routes
	pubImporterSvc := pubImporter.NewService(database)
	pubImporterHdl := pubImporter.NewHandler(pubImporterSvc)
	pubImporter.RegisterRoutes(r, pubImporterHdl)

	// pub exporter routes
	pubExporterSvc := pubExporter.NewService(database)
	pubExporterHdl := pubExporter.NewHandler(pubExporterSvc)
	pubExporter.RegisterRoutes(r, pubExporterHdl)

	// work exporter routes
	workExporterSvc := workExporter.NewService(database)
	workExporterHdl := workExporter.NewHandler(workExporterSvc)
	workExporter.RegisterRoutes(r, workExporterHdl)

	//info exporter routes
	infoExporterSvc := infoExporter.NewService(database)
	infoExporterHdl := infoExporter.NewHandler(infoExporterSvc)
	infoExporter.RegisterRoutes(r, infoExporterHdl)

	// email routes
	emailSvc := email.NewService(database)
	emailHdl := email.NewHandler(emailSvc)
	email.RegisterRoutes(r, emailHdl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Protected routes
	authGroup := r.Group("/auth")
	authGroup.Use(auth.JWT(authService))
	{
	}
	alumniService := alumni.NewService(database)
	alumniHandler := alumni.NewHandler(alumniService)
	alumni.RegisterRoutes(r, alumniHandler)

	//// Routes for alumni
	//r.GET("/alumni", alumniHandler.GetAlumni)
	//r.GET("/alumni/:id", alumniHandler.GetAlumni)
	//r.POST("/alumni", alumniHandler.CreateAlumni)
	//r.PUT("/alumni/:id", alumniHandler.UpdateAlumni)
	//r.DELETE("/alumni/:id", alumniHandler.DeleteAlumni)
	//r.POST("/alumni/import", alumniHandler.ImportAlumni)

	r.Static("/uploads", "./uploads")

	err = r.Run(":9105")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
