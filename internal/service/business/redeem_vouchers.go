package business

// import (
// 	"context"
// 	"errors"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// 	"gorm.io/gorm"

// 	api "promotion_service/api"
// 	"promotion_service/internal/repository"
// )

// func (b *biz) RedeemVouchers(ctx context.Context, request *api.RedeemVouchersRequest) (*api.RedeemVouchersResponse, error) {
// 	// TODO
// 	// call user service to check user exists

// 	// check event exists
// 	_, err := b.repo.GetEventById(ctx, request.GetEventId())
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, status.Errorf(codes.NotFound, "Event %d not found", request.GetEventId())
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO
// 	// call game service to check game exists

// 	// check promotion exists
// 	promotion, err := b.repo.GetPromotion(ctx, &repository.GetPromotionParams{
// 		UserId:  request.GetUserId(),
// 		EventId: request.GetEventId(),
// 		GameId:  request.GetGameId(),
// 	})

// 	// create new promotion if not exist
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, status.Errorf(codes.NotFound, "Promotion %d not found")
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil
// }
