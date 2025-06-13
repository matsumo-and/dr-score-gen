package repository

import (
	"context"

	"github.com/matsumo-and/dr-score-gen/internal/domain/entity"
)

// ScoreRepository defines the interface for score data operations
type ScoreRepository interface {
	// Create saves a new score
	Create(ctx context.Context, score *entity.Score) error

	// FindByID retrieves a score by its ID
	FindByID(ctx context.Context, id string) (*entity.Score, error)

	// FindAll retrieves all scores
	FindAll(ctx context.Context) ([]*entity.Score, error)

	// Update modifies an existing score
	Update(ctx context.Context, score *entity.Score) error

	// Delete removes a score by its ID
	Delete(ctx context.Context, id string) error
}
