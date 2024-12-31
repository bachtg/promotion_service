package repository

import (
	"time"

	"promotion_service/constant"
)

type Paginate struct {
	Limit       int
	CurrentPage int
	TotalItems  int64
}

type Event struct {
	Id               int64                `gorm:"id"`
	Name             string               `gorm:"column:name"`
	Image            string               `gorm:"column:image"`
	VouchersQuantity int32                `gorm:"column:vouchers_quantity"`
	PartnerId        int64                `gorm:"column:partner_id"`
	Status           constant.EventStatus `gorm:"column:status"`
	FromDate         time.Time            `gorm:"column:from_date"`
	ToDate           time.Time            `gorm:"column:to_date"`
	CreatedAt        time.Time            `gorm:"column:created_at"`
	UpdatedAt        time.Time            `gorm:"column:updated_at"`
}

type Voucher struct {
	Id        int64                  `gorm:"id"`
	Code      string                 `gorm:"column:code"`
	QrCode    string                 `gorm:"column:qr_code"`
	Price     int64                  `gorm:"column:price"`
	Currency  string                 `gorm:"column:curency"`
	PartnerId int64                  `gorm:"column:partner_id"`
	EventId   int64                  `gorm:"column:event_id"`
	Status    constant.VoucherStatus `gorm:"column:status"`
	ExpiredAt time.Time              `gorm:"column:expired_at"`
	CreatedAt time.Time              `gorm:"column:created_at"`
	UpdatedAt time.Time              `gorm:"column:updated_at"`
}

type Promotion struct {
	Id        int64     `gorm:"id"`
	UserId    int64     `gorm:"user_id"`
	EventId   int64     `gorm:"event_id"`
	GameId    int64     `gorm:"game_id"`
	Points    int64     `gorm:"points"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Event) TableName() string     { return "events" }
func (Voucher) TableName() string   { return "vouchers" }
func (Promotion) TableName() string { return "promotions" }
