package repository

import "time"

type CreateEventParams struct {
	Name             string
	Image            string
	VouchersQuantity int32
	FromDate         time.Time
	ToDate           time.Time
	PartnerId        int32
}

type (
	GetListEventsParams struct {
		Name     string
		FromDate int64
		ToDate   int64
		Paginate *Paginate
	}
	GetListEventsResult struct {
		Items    []*Event
		Paginate *Paginate
	}
)

type GetPromotionParams struct {
	UserId  int64
	EventId int64
	GameId  int64
}
