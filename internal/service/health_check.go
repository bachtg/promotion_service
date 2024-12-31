package service

import (
	"context"
	"log"

	api "promotion_service/api"
)

func (s *Service) HealthCheck(ctx context.Context, in *api.Empty) (*api.Empty, error) {
	log.Print("HealthCheck OK")
	return &api.Empty{}, nil
}
