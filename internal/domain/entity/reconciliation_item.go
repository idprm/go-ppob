package entity

// CREATE TABLE reconciliation_items (
//   id BIGSERIAL PRIMARY KEY,
//   run_id UUID REFERENCES reconciliation_runs(id),
//   transaction_id UUID,
//   provider_tx_id TEXT,
//   expected_status TEXT,
//   provider_status TEXT,
//   resolution TEXT,
//   notes TEXT
// );

type ReconciliationItem struct {
	ID             int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id,omitempty"`
	RunID          int64  `gorm:"column:run_id" json:"run_id,omitempty"`
	TransactionID  int64  `gorm:"column:transaction_id" json:"transaction_id,omitempty"`
	ProviderTxID   string `gorm:"column:provider_tx_id" json:"provider_tx_id,omitempty"`
	ExpectedStatus string `gorm:"column:expected_status" json:"expected_status,omitempty"`
	ProviderStatus string `gorm:"column:provider_status" json:"provider_status,omitempty"`
	Resolution     string `gorm:"column:resolution" json:"resolution,omitempty"`
	Notes          string `gorm:"column:notes" json:"notes,omitempty"`
}
