package entity

import "gorm.io/gorm"

// CREATE TABLE providers (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   code TEXT UNIQUE NOT NULL, -- e.g., DIGI_01, PLN_PROVIDER
//   name TEXT,
//   active BOOLEAN DEFAULT TRUE,
//   meta JSONB,
//   created_at timestamptz DEFAULT now()
// );

type Provider struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	Code     string `gorm:"column:code" json:"code,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	IsActive bool   `gorm:"column:is_active" json:"is_active,omitempty"`
	Meta     string `gorm:"column:meta" json:"meta,omitempty"`
	gorm.Model
}
