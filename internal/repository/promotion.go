package repository

import "context"

func (r *repository) CreatePromotion(ctx context.Context, promotion *Promotion) (int64, error) {
	return promotion.Id, r.WithContext(ctx).Table(Promotion{}.TableName()).Create(promotion).Error
}

func (r *repository) GetPromotion(ctx context.Context, params *GetPromotionParams) (*Promotion, error) {
	var (
		promotion Promotion
		err       = r.WithContext(ctx).Table(Promotion{}.TableName()).Where("user_id = ? and event_id = ? and game_id = ?", params.UserId, params.EventId, params.GameId).First(&promotion).Error
	)
	return &promotion, err
}

func (r *repository) UpdatePromotion(ctx context.Context, promotionId int64, points int64) error {
	return r.WithContext(ctx).
		Table(Promotion{}.TableName()).
		Where("id = ?", promotionId).Update("points", points).Error
}
