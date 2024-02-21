package enums

type PaymentStatusEnum string

const (
	COMPLETED PaymentStatusEnum = "COMPLETED"
	FAILED    PaymentStatusEnum = "FAILED"
	PENDING   PaymentStatusEnum = "PENDING"
	REFUNDED  PaymentStatusEnum = "REFUNDED"
	UNPAID    PaymentStatusEnum = "UNPAID"
)

type AccountStatusEnum string

const (
	ACTIVE      AccountStatusEnum = "ACTIVE"
	CLOSED      AccountStatusEnum = "CLOSED"
	CANCELLED   AccountStatusEnum = "CANCELLED"
	BLACKLISTED AccountStatusEnum = "BLACKLISTED"
	NONE        AccountStatusEnum = "NONE"
)