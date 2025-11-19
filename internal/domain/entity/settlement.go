package entity

import (
	"time"

	"gorm.io/gorm"
)

// CREATE TABLE settlements (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   user_id UUID REFERENCES users(id),
//   period_date date,
//   amount BIGINT NOT NULL,
//   type TEXT, -- payout, settlement
//   status TEXT DEFAULT 'PENDING',
//   details JSONB,
//   created_at timestamptz DEFAULT now(),
//   processed_at timestamptz
// );

type Settlement struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	UserID      int64     `gorm:"column:user_id" json:"user_id,omitempty"`
	PeriodDate  string    `gorm:"column:period_date" json:"period_date,omitempty"`
	Amount      int64     `gorm:"column:amount" json:"amount,omitempty"`
	Type        string    `gorm:"column:type" json:"type,omitempty"`
	Status      string    `gorm:"column:status" json:"status,omitempty"`
	Details     string    `gorm:"column:details" json:"details,omitempty"`
	ProcessedAt time.Time `gorm:"column:processed_at" json:"processed_at,omitempty"`
	gorm.Model
}
