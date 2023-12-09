CREATE TABLE IF NOT EXISTS issues(
  id VARCHAR PRIMARY KEY,
  title VARCHAR NOT NULL,
  detail TEXT NOT NULL,
  customer_name VARCHAR NOT NULL,
  customer_email VARCHAR NOT NULL,
  taken_by VARCHAR,
  solved_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);