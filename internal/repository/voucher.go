package repository

import "context"

func (r *repository) CreateVoucher(ctx context.Context, voucher *Voucher) (int64, error) {
	return voucher.Id, r.WithContext(ctx).Table(Voucher{}.TableName()).Create(voucher).Error
}
