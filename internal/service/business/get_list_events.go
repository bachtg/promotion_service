package business

import (
	"context"

	api "promotion_service/api"
	"promotion_service/constant"
	"promotion_service/internal/repository"
	"promotion_service/utils"
)

func (b *biz) GetListEvents(ctx context.Context, request *api.GetListEventsRequest) (*api.GetListEventsResponse_Data, error) {
	var (
		currentPage = utils.CoalesceInt32(request.CurrentPage, constant.DEFAULT_CURRENT_PAGE)
		pageSize    = utils.CoalesceInt32(request.PageSize, constant.DEFAULT_PAGE_SIZE)
		events      = make([]*api.Event, 0)
	)

	result, err := b.repo.GetListEvents(ctx, &repository.GetListEventsParams{
		Name:     request.Name,
		FromDate: request.FromDate,
		ToDate:   request.ToDate,
		Paginate: &repository.Paginate{
			Limit:       int(pageSize),
			CurrentPage: int(currentPage),
		},
	})
	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		events = append(events, &api.Event{
			Id:               item.Id,
			Name:             item.Name,
			Image:            item.Image,
			VouchersQuantity: item.VouchersQuantity,
			FromDate:         item.FromDate.Unix(),
			ToDate:           item.ToDate.Unix(),
		})
	}

	return &api.GetListEventsResponse_Data{
		TotalRecords: int32(result.Paginate.TotalItems),
		PageSize:     int32(result.Paginate.Limit),
		CurrentPage:  int32(result.Paginate.CurrentPage),
		Events:       events,
	}, nil

}
