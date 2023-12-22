CREATE TABLE IF NOT EXISTS users(
  id UUID DEFAULT generate_ulid() PRIMARY KEY,
  name VARCHAR NOT NULL,
  email VARCHAR NOT NULL UNIQUE,
  password TEXT NOT NULL,
  is_valid BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);