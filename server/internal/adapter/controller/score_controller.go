package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matsumo-and/dr-score-gen/internal/usecase"
)

// ScoreController handles HTTP requests for score operations
type ScoreController struct {
	scoreUsecase usecase.ScoreUsecase
}

// NewScoreController creates a new score controller
func NewScoreController(scoreUsecase usecase.ScoreUsecase) *ScoreController {
	return &ScoreController{
		scoreUsecase: scoreUsecase,
	}
}

// CreateScore handles POST /scores
func (c *ScoreController) CreateScore(ctx *gin.Context) {
	var req struct {
		Title    string `json:"title" binding:"required"`
		Artist   string `json:"artist" binding:"required"`
		Tempo    int    `json:"tempo" binding:"required,min=1"`
		TimeSign string `json:"timeSign" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	score, err := c.scoreUsecase.CreateScore(ctx, req.Title, req.Artist, req.Tempo, req.TimeSign)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create score"})
		return
	}

	ctx.JSON(http.StatusCreated, score)
}

// GetScore handles GET /scores/:id
func (c *ScoreController) GetScore(ctx *gin.Context) {
	id := ctx.Param("id")

	score, err := c.scoreUsecase.GetScore(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	ctx.JSON(http.StatusOK, score)
}

// GetAllScores handles GET /scores
func (c *ScoreController) GetAllScores(ctx *gin.Context) {
	scores, err := c.scoreUsecase.GetAllScores(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scores"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"scores": scores})
}

// UpdateScore handles PUT /scores/:id
func (c *ScoreController) UpdateScore(ctx *gin.Context) {
	id := ctx.Param("id")

	var req struct {
		Title    string `json:"title"`
		Artist   string `json:"artist"`
		Tempo    int    `json:"tempo"`
		TimeSign string `json:"timeSign"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get existing score
	score, err := c.scoreUsecase.GetScore(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	// Update fields
	if req.Title != "" {
		score.Title = req.Title
	}
	if req.Artist != "" {
		score.Artist = req.Artist
	}
	if req.Tempo > 0 {
		score.Tempo = req.Tempo
	}
	if req.TimeSign != "" {
		score.TimeSign = req.TimeSign
	}

	if err := c.scoreUsecase.UpdateScore(ctx, score); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update score"})
		return
	}

	ctx.JSON(http.StatusOK, score)
}

// DeleteScore handles DELETE /scores/:id
func (c *ScoreController) DeleteScore(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.scoreUsecase.DeleteScore(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete score"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
