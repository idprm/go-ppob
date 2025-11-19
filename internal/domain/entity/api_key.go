package entity

import "gorm.io/gorm"

// CREATE TABLE api_keys (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   user_id UUID REFERENCES users(id) ON DELETE CASCADE,
//   key_id TEXT UNIQUE NOT NULL, -- public id
//   key_hash TEXT NOT NULL,      -- store hashed
//   name TEXT,
//   revoked BOOLEAN DEFAULT FALSE,
//   created_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_api_keys_user ON api_keys(user_id);

type APIKey struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id"`
	gorm.Model
}
