package usecase

import (
	"antrein/dd-dashboard-auth/application/common/repository"
	"antrein/dd-dashboard-auth/internal/usecase/auth"
	"antrein/dd-dashboard-auth/model/config"
)

type CommonUsecase struct {
	AuthUsecase *auth.Usecase
}

func NewCommonUsecase(cfg *config.Config, repo *repository.CommonRepository) (*CommonUsecase, error) {
	authUsecase := auth.New(cfg, repo.TenantRepo)

	commonUC := CommonUsecase{
		AuthUsecase: authUsecase,
	}
	return &commonUC, nil
}
