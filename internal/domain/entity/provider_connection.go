package entity

import "gorm.io/gorm"

// CREATE TABLE provider_connections (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   provider_id UUID REFERENCES providers(id) ON DELETE CASCADE,
//   conn_type TEXT NOT NULL, -- api, jabber, soap, iso8583
//   endpoint TEXT,
//   credentials JSONB,
//   priority INT DEFAULT 100,
//   active BOOLEAN DEFAULT TRUE,
//   created_at timestamptz DEFAULT now()
// );

type ProviderConnection struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	ConnType string `gorm:"column:conn_type" json:"conn_type,omitempty"`
	Endpoint string `gorm:"column:endpoint" json:"endpoint,omitempty"`
	Priority int    `gorm:"column:priority;default:100" json:"priority,omitempty"`
	IsActive bool   `gorm:"column:is_active;default:true" json:"is_active,omitempty"`
	gorm.Model
}
