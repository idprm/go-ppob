package entity

import "gorm.io/gorm"

// CREATE TABLE users (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   username TEXT UNIQUE NOT NULL,
//   name TEXT,
//   email TEXT,
//   role TEXT NOT NULL, -- admin, merchant, agent, provider
//   status TEXT DEFAULT 'active',
//   created_at timestamptz DEFAULT now(),
//   updated_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_users_role ON users(role);

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;column:id"`
	Email    string `gorm:"column:email" json:"email,omitempty"`
	Role     string `gorm:"column:role" json:"role,omitempty"`
	IsActive bool   `gorm:"column:is_active" json:"is_active,omitempty"`
	gorm.DB
}
