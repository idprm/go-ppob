package entity

// CREATE TABLE idempotency_keys (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   key TEXT UNIQUE NOT NULL,
//   user_id UUID,
//   request_hash TEXT,
//   response_snapshot JSONB,
//   status TEXT DEFAULT 'used', -- used, in-progress
//   created_at timestamptz DEFAULT now(),
//   expires_at timestamptz
// );
// CREATE INDEX idx_idemp_key ON idempotency_keys(key);

type IdempotencyKey struct {
	ID  int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	Key string `gorm:"column:key" json:"key,omitempty"`
}
