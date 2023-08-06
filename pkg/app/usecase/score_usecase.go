package usecase

import (
	"sam-app/pkg/app/repository"
	"sam-app/pkg/domain/model"
)

type ScoreUseCaseImpl struct {
	scoreRepository repository.ScoreRepository
}

func NewScoreUseCase(scoreRepo repository.ScoreRepository) ScoreUseCase {
	return &ScoreUseCaseImpl{
		scoreRepository: scoreRepo,
	}
}

func (su *ScoreUseCaseImpl) Create(score *model.Score) error {
	err := su.scoreRepository.Create(score)
	if err != nil {
		return err
	}
	return nil
}
