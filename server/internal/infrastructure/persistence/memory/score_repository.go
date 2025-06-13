package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/matsumo-and/dr-score-gen/internal/domain/entity"
	"github.com/matsumo-and/dr-score-gen/internal/domain/repository"
)

// scoreRepository is an in-memory implementation of ScoreRepository
type scoreRepository struct {
	mu     sync.RWMutex
	scores map[string]*entity.Score
	nextID int
}

// NewScoreRepository creates a new in-memory score repository
func NewScoreRepository() repository.ScoreRepository {
	return &scoreRepository{
		scores: make(map[string]*entity.Score),
		nextID: 1,
	}
}

func (r *scoreRepository) Create(ctx context.Context, score *entity.Score) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Generate ID
	score.ID = fmt.Sprintf("%d", r.nextID)
	r.nextID++

	r.scores[score.ID] = score
	return nil
}

func (r *scoreRepository) FindByID(ctx context.Context, id string) (*entity.Score, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	score, exists := r.scores[id]
	if !exists {
		return nil, errors.New("score not found")
	}

	return score, nil
}

func (r *scoreRepository) FindAll(ctx context.Context) ([]*entity.Score, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	scores := make([]*entity.Score, 0, len(r.scores))
	for _, score := range r.scores {
		scores = append(scores, score)
	}

	return scores, nil
}

func (r *scoreRepository) Update(ctx context.Context, score *entity.Score) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.scores[score.ID]; !exists {
		return errors.New("score not found")
	}

	r.scores[score.ID] = score
	return nil
}

func (r *scoreRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.scores[id]; !exists {
		return errors.New("score not found")
	}

	delete(r.scores, id)
	return nil
}
