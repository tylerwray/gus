CREATE TABLE user_bank_accounts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users (id),
  plaid_access_token VARCHAR(50) NOT NULL,
  plaid_item_id VARCHAR(60) NOT NULL,
  inserted_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

