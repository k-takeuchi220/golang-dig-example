package di

import (
	"sam-app/internal/controller"
	"sam-app/pkg/app/repository"
	"sam-app/pkg/app/usecase"
	"sam-app/pkg/infrastructure/database"

	"go.uber.org/dig"
)

// BuildContainer はDIコンテナを構築して返します
func BuildContainer() *dig.Container {
	container := dig.New()

	// データベース接続関数を登録
	container.Provide(database.ConnectDB)

	// Controllerの実装を登録
	container.Provide(controller.NewScoreController)

	// Repositoryの実装を登録
	container.Provide(repository.NewScoreRepository)

	// // UseCaseの実装を登録
	container.Provide(usecase.NewScoreUseCase)

	return container
}
