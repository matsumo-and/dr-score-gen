package router

import (
	"dr-score-gen/internal/adapter/controller"
	"dr-score-gen/internal/infrastructure/persistence/memory"
	"dr-score-gen/internal/usecase"

	"github.com/gin-gonic/gin"
)

// NewRouter creates and configures a new router with all dependencies
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Initialize repositories
	scoreRepo := memory.NewScoreRepository()

	// Initialize use cases
	scoreUsecase := usecase.NewScoreUsecase(scoreRepo)

	// Initialize controllers
	scoreController := controller.NewScoreController(scoreUsecase)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "OK",
		})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		scores := v1.Group("/scores")
		{
			scores.POST("", scoreController.CreateScore)
			scores.GET("", scoreController.GetAllScores)
			scores.GET("/:id", scoreController.GetScore)
			scores.PUT("/:id", scoreController.UpdateScore)
			scores.DELETE("/:id", scoreController.DeleteScore)
		}
	}

	return router
}
