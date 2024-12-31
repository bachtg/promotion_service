package business

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	api "promotion_service/api"
	"promotion_service/internal/repository"
)

func (b *biz) GrantPoints(ctx context.Context, request *api.GrantPointsRequest) error {
	// TODO
	// call user service to check user exists

	// check event exists
	_, err := b.repo.GetEventById(ctx, request.GetEventId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return status.Errorf(codes.NotFound, "Event %d not found", request.GetEventId())
	}
	if err != nil {
		return err
	}

	// TODO
	// call game service to check game exists

	// upsert points
	promotion, err := b.repo.GetPromotion(ctx, &repository.GetPromotionParams{
		UserId:  request.GetUserId(),
		EventId: request.GetEventId(),
		GameId:  request.GetGameId(),
	})

	// create new promotion if not exist
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = b.repo.CreatePromotion(ctx, &repository.Promotion{
			UserId:  request.GetUserId(),
			EventId: request.GetEventId(),
			GameId:  request.GetGameId(),
			Points:  request.GetPoints(),
		})
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	// update promotion if exist
	err = b.repo.UpdatePromotion(ctx, promotion.Id, promotion.Points+request.GetPoints())
	if err != nil {
		return err
	}

	return nil
}
