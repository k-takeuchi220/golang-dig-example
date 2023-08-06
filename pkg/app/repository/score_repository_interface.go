package repository

import (
	"sam-app/pkg/domain/model"
)

type ScoreRepository interface {
	Create(score *model.Score) error
}
