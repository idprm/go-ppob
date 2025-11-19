package entity

// CREATE TABLE transactions (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   reference TEXT UNIQUE NOT NULL, -- client idempotency key
//   user_id UUID REFERENCES users(id),
//   provider_id UUID REFERENCES providers(id),
//   provider_product_id UUID REFERENCES provider_products(id),
//   customer_no TEXT, -- destination phone, meter, etc (consider encryption)
//   amount BIGINT NOT NULL, -- amount charged to customer (in cents)
//   provider_price BIGINT,  -- price paid to provider
//   fee BIGINT DEFAULT 0,
//   total_debit BIGINT NOT NULL, -- amount + fee
//   status TEXT NOT NULL DEFAULT 'PENDING', -- PENDING, PROCESSING, SUCCESS, FAILED, REFUNDED
//   last_attempt_at timestamptz,
//   provider_tx_id TEXT, -- provider-side transaction id
//   provider_response JSONB,
//   created_at timestamptz DEFAULT now(),
//   updated_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_tx_user_status ON transactions(user_id, status);
// CREATE INDEX idx_tx_reference ON transactions(reference);

type Transaction struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	Reference string `gorm:"column:reference" json:"reference,omitempty"`
}
