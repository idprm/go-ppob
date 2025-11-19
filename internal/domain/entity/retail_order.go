package entity

import "gorm.io/gorm"

// -- retail orders
// CREATE TABLE retail_orders (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   order_no TEXT UNIQUE NOT NULL,
//   retail_user_id UUID REFERENCES retail_users(id),
//   provider_product_id UUID REFERENCES provider_products(id),
//   customer_no TEXT, -- target msisdn/meter/etc.
//   amount BIGINT NOT NULL, -- price displayed to user
//   fee BIGINT DEFAULT 0,
//   total BIGINT NOT NULL, -- total charged to user = amount + fee
//   payment_method TEXT, -- e.g., VA_BNI, QRIS_SHOPEEPAY, CARD_VISA
//   payment_provider_id UUID, -- which PG/gateway used
//   payment_reference TEXT, -- PG transaction id
//   status TEXT NOT NULL DEFAULT 'CREATED', -- CREATED, PENDING_PAYMENT, PAID, FULFILLING, SUCCESS, FAILED, REFUNDED
//   idempotency_key TEXT,
//   created_at timestamptz DEFAULT now(),
//   updated_at timestamptz DEFAULT now()
// );

type RetailOrder struct {
	ID                int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	OrderNo           string `gorm:"column:order_no" json:"order_no,omitempty"`
	RetailUserID      int64  `gorm:"column:retail_user_id" json:"retail_user_id,omitempty"`
	ProviderProductID int64  `gorm:"column:provider_product_id" json:"provider_product_id,omitempty"`
	CustomerNo        string `gorm:"column:customer_no" json:"customer_no,omitempty"`
	Amount            int64  `gorm:"column:amount" json:"amount,omitempty"`
	Fee               int64  `gorm:"column:fee" json:"fee,omitempty"`
	Total             int64  `gorm:"column:total" json:"total,omitempty"`
	PaymentMethod     string `gorm:"column:payment_method" json:"payment_method,omitempty"`
	PaymentProviderID int64  `gorm:"column:payment_provider_id" json:"payment_provider_id,omitempty"`
	PaymentReference  string `gorm:"column:payment_reference" json:"payment_reference,omitempty"`
	Status            string `gorm:"column:status" json:"status,omitempty"`
	IdempotencyKey    string `gorm:"column:idempotency_key" json:"idempotency_key,omitempty"`
	gorm.Model
}
