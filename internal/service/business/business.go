package business

import (
	"context"

	api "promotion_service/api"
	"promotion_service/internal/repository"
)

type Biz interface {
	CreateEvent(ctx context.Context, request *api.CreateEventRequest) (*api.CreateEventResponse, error)
	GetListEvents(ctx context.Context, request *api.GetListEventsRequest) (*api.GetListEventsResponse_Data, error)
	GrantPoints(ctx context.Context, request *api.GrantPointsRequest) error
}

type biz struct {
	repo repository.Repository
}

func NewBiz(
	repo repository.Repository,
) Biz {
	return &biz{
		repo: repo,
	}
}
