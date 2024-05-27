package repository

import (
	"antrein/dd-dashboard-auth/application/common/resource"
	"antrein/dd-dashboard-auth/internal/repository/tenant"
	"antrein/dd-dashboard-auth/model/config"
)

type CommonRepository struct {
	TenantRepo *tenant.Repository
}

func NewCommonRepository(cfg *config.Config, rsc *resource.CommonResource) (*CommonRepository, error) {
	tenantRepo := tenant.New(cfg, rsc.Db)

	commonRepo := CommonRepository{
		TenantRepo: tenantRepo,
	}
	return &commonRepo, nil
}
