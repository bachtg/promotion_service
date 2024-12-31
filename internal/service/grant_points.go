package service

import (
	"context"

	"google.golang.org/genproto/googleapis/rpc/code"

	api "promotion_service/api"
	"promotion_service/constant"
)

func (s *Service) GrantPoints(ctx context.Context, request *api.GrantPointsRequest) (*api.GrantPointsResponse, error) {
	err := s.biz.GrantPoints(ctx, request)
	if err != nil {
		return nil, err
	}

	return &api.GrantPointsResponse{
		Code:    constant.ResponseCode_SUCCESS,
		Message: code.Code_OK.String(),
	}, nil
}
