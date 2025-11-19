package entity

// -- wallet (current balance cached, for quick read)
// CREATE TABLE wallets (
//   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
//   balance BIGINT NOT NULL DEFAULT 0, -- in smallest unit (e.g., cents or rupiah)
//   currency TEXT DEFAULT 'IDR',
//   updated_at timestamptz DEFAULT now()
// );

// -- immutable ledger: every mutation to wallet creates wallet_entries row
// CREATE TABLE wallet_entries (
//   id BIGSERIAL PRIMARY KEY,
//   wallet_id UUID REFERENCES wallets(id) ON DELETE CASCADE,
//   user_id UUID, -- denormalized
//   tx_type TEXT NOT NULL, -- debit, credit, fee, adjustment
//   amount BIGINT NOT NULL, -- positive numbers only, sign determined by tx_type
//   balance_after BIGINT NOT NULL,
//   reference TEXT, -- e.g. transaction_id or topup_id
//   metadata JSONB,
//   created_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_wallet_entries_wallet ON wallet_entries(wallet_id, created_at);

type Wallet struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
}

type WalletEntry struct {
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
}
