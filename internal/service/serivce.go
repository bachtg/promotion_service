package service

import (
	api "promotion_service/api"
	"promotion_service/config"
	"promotion_service/internal/repository"
	"promotion_service/internal/service/business"
)

type Service struct {
	cfg  *config.Config
	repo repository.Repository
	biz  business.Biz

	api.UnimplementedLoyaltyServiceServer
}

func NewService(
	cfg *config.Config,
	repo repository.Repository,
) *Service {

	biz := business.NewBiz(repo)

	return &Service{
		cfg:  cfg,
		repo: repo,
		biz:  biz,
	}
}
