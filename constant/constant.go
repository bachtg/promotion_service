package constant

const (
	DEFAULT_CURRENT_PAGE = 1
	DEFAULT_PAGE_SIZE    = 10
)

type EventStatus string

const (
	EventStatus_CREATED EventStatus = "created"
	EventStatus_ONGOING EventStatus = "ongoing"
	EventStatus_ENDED   EventStatus = "ended"
)

type VoucherStatus string

const (
	VoucherStatus_ACTIVE  VoucherStatus = "active"
	VoucherStatus_EXPIRED VoucherStatus = "expired"
	VoucherStatus_USED    VoucherStatus = "used"
)

const (
	Currency_VND = "VND"
)

const (
	VoucherLength = 12
	VoucherPrice  = 100000
)

const (
	ResponseCode_SUCCESS int32 = 0
)
