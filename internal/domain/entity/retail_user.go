package entity

// -- retail users (customer)
// CREATE TABLE retail_users (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   email TEXT UNIQUE,
//   phone TEXT,
//   name TEXT,
//   password_hash TEXT,
//   status TEXT DEFAULT 'active',
//   created_at timestamptz DEFAULT now()
// );

type RetailUser struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
}
