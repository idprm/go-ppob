package entity

import "gorm.io/gorm"

// CREATE TABLE provider_logs (
//   id BIGSERIAL PRIMARY KEY,
//   transaction_id UUID REFERENCES transactions(id),
//   provider_code TEXT,
//   request_body JSONB,
//   response_body JSONB,
//   status_code INT,
//   created_at timestamptz DEFAULT now()
// );

type ProviderLog struct {
	ID            int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	TransactionID int64  `gorm:"column:transaction_id" json:"transaction_id,omitempty"`
	ProviderCode  string `gorm:"column:provider_code" json:"provider_code,omitempty"`
	RequestBody   string `gorm:"column:request_body" json:"request_body,omitempty"`
	ResponseBody  string `gorm:"column:response_body" json:"response_body,omitempty"`
	StatusCode    int    `gorm:"column:status_code" json:"status_code,omitempty"`
	gorm.Model
}
