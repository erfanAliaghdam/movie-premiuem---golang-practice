package entity

import "time"

type LicensePayment struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	PaymentAmount int64     `json:"payment_amount"`
	UserLicenseID int64     `json:"user_license_id"`
	MerchantCode  string    `json:"merchant_code"`
}
