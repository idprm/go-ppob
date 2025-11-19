package entity

import "gorm.io/gorm"

// CREATE TABLE transaction_attempts (
//   id BIGSERIAL PRIMARY KEY,
//   transaction_id UUID REFERENCES transactions(id) ON DELETE CASCADE,
//   attempt_no INT NOT NULL,
//   provider_connection_id UUID REFERENCES provider_connections(id),
//   request JSONB,
//   response JSONB,
//   response_status TEXT,
//   latency_ms INT,
//   created_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_attempts_tx ON transaction_attempts(transaction_id);

type TransactionAttempt struct {
	ID                   int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	AttempNo             int    `gorm:"column:attempt_no" json:"attempt_no,omitempty"`
	ProviderConnectionID int64  `gorm:"column:provider_connection_id" json:"provider_connection_id,omitempty"`
	Request              string `gorm:"column:request" json:"request,omitempty"`
	Response             string `gorm:"column:response" json:"response,omitempty"`
	ResponseStatus       string `gorm:"column:response_status" json:"response_status,omitempty"`
	LatencyMs            int    `gorm:"column:latency_ms" json:"latency_ms,omitempty"`
	gorm.Model
}
