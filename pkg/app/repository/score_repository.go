package repository

import (
	"database/sql"
	"sam-app/pkg/domain/model"
)

type ScoreRepositoryImpl struct {
	db *sql.DB
}

func NewScoreRepository(db *sql.DB) ScoreRepository {
	return &ScoreRepositoryImpl{
		db: db,
	}
}

func (sr *ScoreRepositoryImpl) Create(score *model.Score) error {
	// トランザクションを開始
	tx, err := sr.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		// エラーがある場合はトランザクションをロールバック
		if err != nil {
			tx.Rollback()
		}
	}()

	// スコアのデータを準備
	query := "INSERT INTO scores (name, score) VALUES (?, ?)"

	// スコアの登録
	_, err = tx.Exec(query, score.Name, score.Score)
	if err != nil {
		return err
	}

	// コミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
