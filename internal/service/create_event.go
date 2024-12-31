package service

import (
	"context"

	api "promotion_service/api"
)

func (s *Service) CreateEvent(ctx context.Context, request *api.CreateEventRequest) (*api.CreateEventResponse, error) {
	res, err := s.biz.CreateEvent(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
