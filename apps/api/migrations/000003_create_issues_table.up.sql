CREATE TABLE IF NOT EXISTS issues(
  id UUID DEFAULT generate_ulid() PRIMARY KEY,
  title VARCHAR NOT NULL,
  detail TEXT NOT NULL,
  customer_name VARCHAR NOT NULL,
  customer_email VARCHAR NOT NULL,
  taken_by UUID,
  solved_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);