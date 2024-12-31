package service

import (
	"context"

	"google.golang.org/genproto/googleapis/rpc/code"

	api "promotion_service/api"
	"promotion_service/constant"
)

func (s *Service) GetListEvents(ctx context.Context, request *api.GetListEventsRequest) (*api.GetListEventsResponse, error) {
	data, err := s.biz.GetListEvents(ctx, request)
	if err != nil {
		return nil, err
	}

	return &api.GetListEventsResponse{
		Code:    constant.ResponseCode_SUCCESS,
		Message: code.Code_OK.String(),
		Data:    data,
	}, nil
}
