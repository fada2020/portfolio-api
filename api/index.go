package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"portfolio-api/database"
	"portfolio-api/handlers"
)

var router *gin.Engine

func init() {
	// Initialize Supabase database
	if err := database.InitSupabase(); err != nil {
		log.Printf("Failed to initialize Supabase: %v", err)
	}

	// Insert sample data on startup
	if err := database.InsertSampleData(); err != nil {
		log.Printf("Warning: Failed to insert sample data: %v", err)
	}

	// Set Gin mode for production
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS configuration for portfolio frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"https://hyoukjoolee.github.io",     // GitHub Pages
		"https://portfolio-api.vercel.app",  // Vercel deployment
		"http://localhost:3000",             // Local development
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Health check endpoint
	router.GET("/health", healthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// User management
		users := v1.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)
			users.GET("/:id", handlers.GetUserByID)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}

		// Projects showcase
		projects := v1.Group("/projects")
		{
			projects.GET("", handlers.GetProjects)
			projects.GET("/:id", handlers.GetProject)
			projects.POST("", handlers.CreateProject)
			projects.PUT("/:id", handlers.UpdateProject)
			projects.DELETE("/:id", handlers.DeleteProject)
		}

		// Skills and technologies
		skills := v1.Group("/skills")
		{
			skills.GET("", handlers.GetSkills)
			skills.POST("", handlers.AddSkill)
			skills.DELETE("/:id", handlers.RemoveSkill)
		}

		// Contact form
		contact := v1.Group("/contact")
		{
			contact.POST("", handlers.SubmitContactForm)
		}

		// Portfolio statistics
		stats := v1.Group("/stats")
		{
			stats.GET("/views", handlers.GetViewStats)
			stats.GET("/projects", handlers.GetProjectStats)
			stats.POST("/visit", handlers.RecordVisit)
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Handler is the main entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"version":   "1.0.0",
		"platform":  "vercel",
	})
}