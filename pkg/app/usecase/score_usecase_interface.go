package usecase

import (
	"sam-app/pkg/domain/model"
)

type ScoreUseCase interface {
	Create(score *model.Score) error
}
