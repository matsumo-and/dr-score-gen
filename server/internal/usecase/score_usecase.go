package usecase

import (
	"context"

	"dr-score-gen/internal/domain/entity"
	"dr-score-gen/internal/domain/repository"
)

// ScoreUsecase defines the business logic for score operations
type ScoreUsecase interface {
	CreateScore(ctx context.Context, title, artist string, tempo int, timeSign string) (*entity.Score, error)
	GetScore(ctx context.Context, id string) (*entity.Score, error)
	GetAllScores(ctx context.Context) ([]*entity.Score, error)
	UpdateScore(ctx context.Context, score *entity.Score) error
	DeleteScore(ctx context.Context, id string) error
}

// scoreUsecase implements ScoreUsecase interface
type scoreUsecase struct {
	scoreRepo repository.ScoreRepository
}

// NewScoreUsecase creates a new score usecase
func NewScoreUsecase(scoreRepo repository.ScoreRepository) ScoreUsecase {
	return &scoreUsecase{
		scoreRepo: scoreRepo,
	}
}

func (u *scoreUsecase) CreateScore(ctx context.Context, title, artist string, tempo int, timeSign string) (*entity.Score, error) {
	score := entity.NewScore(title, artist, tempo, timeSign)

	if err := u.scoreRepo.Create(ctx, score); err != nil {
		return nil, err
	}

	return score, nil
}

func (u *scoreUsecase) GetScore(ctx context.Context, id string) (*entity.Score, error) {
	return u.scoreRepo.FindByID(ctx, id)
}

func (u *scoreUsecase) GetAllScores(ctx context.Context) ([]*entity.Score, error) {
	return u.scoreRepo.FindAll(ctx)
}

func (u *scoreUsecase) UpdateScore(ctx context.Context, score *entity.Score) error {
	return u.scoreRepo.Update(ctx, score)
}

func (u *scoreUsecase) DeleteScore(ctx context.Context, id string) error {
	return u.scoreRepo.Delete(ctx, id)
}
