package entity

import (
	"time"

	"gorm.io/gorm"
)

// CREATE TABLE deposits (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   user_id UUID REFERENCES users(id),
//   amount BIGINT NOT NULL,
//   method TEXT, -- bank_transfer, va, manual
//   reference TEXT,
//   status TEXT DEFAULT 'PENDING',
//   proof JSONB,
//   created_at timestamptz DEFAULT now(),
//   verified_at timestamptz
// );

type Deposit struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	UserID    int64     `gorm:"column:user_id" json:"user_id,omitempty"`
	Amount    int64     `gorm:"column:amount" json:"amount,omitempty"`
	Method    string    `gorm:"column:method" json:"method,omitempty"`
	Reference string    `gorm:"column:reference" json:"reference,omitempty"`
	Status    string    `gorm:"column:status" json:"status,omitempty"`
	Proof     string    `gorm:"column:proof" json:"proof,omitempty"`
	VerifyAt  time.Time `gorm:"column:verified_at" json:"verified_at,omitempty"`
	gorm.Model
}
