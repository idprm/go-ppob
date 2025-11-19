package entity

import "gorm.io/gorm"

// CREATE TABLE provider_products (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   provider_id UUID REFERENCES providers(id) ON DELETE CASCADE,
//   product_code TEXT NOT NULL, -- e.g., PULSA_5K, PLN_TOKEN
//   category TEXT,
//   description TEXT,
//   attributes JSONB, -- extra fields per product
//   created_at timestamptz DEFAULT now()
// );

type ProviderProduct struct {
	ID          int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	ProviderID  int64  `gorm:"column:provider_id" json:"provider_id,omitempty"`
	ProductCode string `gorm:"column:product_code" json:"product_code,omitempty"`
	Category    string `gorm:"column:category" json:"category,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	Attributes  string `gorm:"column:attributes" json:"attributes,omitempty"`
	gorm.Model
}
