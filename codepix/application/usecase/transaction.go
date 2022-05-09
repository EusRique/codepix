package usecase

import "github.com/EusRique/codepix/domain/model"

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyRepositoryInterface
}
