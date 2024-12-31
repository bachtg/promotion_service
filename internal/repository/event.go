package repository

import (
	"context"
	"time"
)

func (r *repository) CreateEvent(ctx context.Context, event *Event) (int64, error) {
	return event.Id, r.WithContext(ctx).Table(Event{}.TableName()).Create(event).Error
}

func (r *repository) GetListEvents(ctx context.Context, params *GetListEventsParams) (*GetListEventsResult, error) {
	var (
		events     = make([]*Event, 0)
		totalItems int64
		paginate   = params.Paginate
	)

	query := r.WithContext(ctx).Table(Event{}.TableName())

	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.FromDate != 0 {
		query = query.Where("from_date >= ?", time.Unix(params.FromDate, 0))
	}
	if params.ToDate != 0 {
		query = query.Where("to_date < ?", time.Unix(params.ToDate, 0))
	}

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	if err := query.Limit(paginate.Limit).Offset(paginate.Limit * (paginate.CurrentPage - 1)).Find(&events).Error; err != nil {
		return nil, err
	}

	return &GetListEventsResult{
		Items: events,
		Paginate: &Paginate{
			Limit:       paginate.Limit,
			CurrentPage: paginate.CurrentPage,
			TotalItems:  totalItems,
		},
	}, nil
}

func (r *repository) GetEventById(ctx context.Context, eventId int64) (*Event, error) {
	var (
		event Event
		err   = r.WithContext(ctx).Table(Event{}.TableName()).Where("id = ?", eventId).First(&event).Error
	)
	return &event, err
}
