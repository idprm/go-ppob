package entity

// CREATE TABLE outbox (
//   id BIGSERIAL PRIMARY KEY,
//   aggregate_type TEXT,
//   aggregate_id UUID,
//   event_type TEXT,
//   payload JSONB,
//   published BOOLEAN DEFAULT FALSE,
//   created_at timestamptz DEFAULT now()
// );
// CREATE INDEX idx_outbox_published ON outbox(published, created_at);

type Outbox struct {
}
