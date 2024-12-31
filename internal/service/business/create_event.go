package business

import (
	"context"
	"strings"
	"time"

	uuid "github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/code"

	api "promotion_service/api"
	"promotion_service/constant"
	"promotion_service/internal/repository"
)

func (b *biz) CreateEvent(ctx context.Context, request *api.CreateEventRequest) (*api.CreateEventResponse, error) {
	newEventId, err := b.repo.CreateEvent(ctx, &repository.Event{
		Name:             request.Name,
		Image:            request.Image,
		VouchersQuantity: request.VouchersQuantity,
		Status:           constant.EventStatus_CREATED,
		FromDate:         time.Unix(request.FromDate, 0),
		ToDate:           time.Unix(request.ToDate, 0),
		PartnerId:        request.PartnerId,
	})
	if err != nil {
		return nil, err
	}

	for i := 0; i < int(request.VouchersQuantity); i++ {
		voucherCode := uuid.New().String()
		voucherCode = strings.ToUpper(strings.ReplaceAll(voucherCode, "-", "")[:constant.VoucherLength])
		b.repo.CreateVoucher(ctx, &repository.Voucher{
			Code:      voucherCode,
			Price:     constant.VoucherPrice,
			Currency:  constant.Currency_VND,
			PartnerId: request.PartnerId,
			EventId:   newEventId,
			Status:    constant.VoucherStatus_ACTIVE,
			ExpiredAt: time.Unix(request.ToDate, 0),
		})
	}

	return &api.CreateEventResponse{
		Code:    constant.ResponseCode_SUCCESS,
		Message: code.Code_OK.String(),
		Data: &api.Event{
			Id:               newEventId,
			Name:             request.Name,
			Image:            request.Image,
			VouchersQuantity: request.VouchersQuantity,
			FromDate:         request.FromDate,
			ToDate:           request.ToDate,
		},
	}, nil
}
