CREATE TABLE IF NOT EXISTS crud.users (
  id serial PRIMARY KEY,
  name text NOT NULL,
  email text,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);
