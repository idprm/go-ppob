package entity

import (
	"time"

	"gorm.io/gorm"
)

// CREATE TABLE provider_price_list (
//   id BIGSERIAL PRIMARY KEY,
//   provider_product_id UUID REFERENCES provider_products(id),
//   price BIGINT NOT NULL,  -- price from provider (in cents)
//   effective_from timestamptz DEFAULT now(),
//   effective_to timestamptz,
//   created_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_price_product ON provider_price_list(provider_product_id, effective_from);

type ProviderPricelist struct {
	ID                int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	ProviderProductID int64     `gorm:"column:provider_product_id" json:"provider_product_id,omitempty"`
	Price             int64     `gorm:"column:price" json:"price,omitempty"`
	EffectiveFrom     time.Time `gorm:"column:effective_from" json:"effective_from,omitempty"`
	EffectiveTo       time.Time `gorm:"column:effective_to" json:"effective_to,omitempty"`
	gorm.Model
}
