package entity

import "time"

// CREATE TABLE reconciliation_runs (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   run_by UUID,
//   run_at timestamptz DEFAULT now(),
//   report JSONB,
//   status TEXT DEFAULT 'COMPLETED'
// );

type ReconciliationRun struct {
	ID     int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	RunBy  int64     `gorm:"column:run_by" json:"run_by,omitempty"`
	RunAt  time.Time `gorm:"column:run_at" json:"run_at,omitempty"`
	Report string    `gorm:"column:report" json:"report,omitempty"`
	Status string    `gorm:"column:status" json:"status,omitempty"`
}
