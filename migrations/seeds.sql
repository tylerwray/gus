/*
 Users 
 */
INSERT INTO users (id, username, PASSWORD)
    VALUES (1, 'tommy', '$2a$10$rbNvZXqhHZSND1jEoAAqlOEbkts2HdPFASUgs1A2cd0NZXNJUn5tu') ON CONFLICT username DO NOTHING;
INSERT INTO users (id, username, PASSWORD)
    VALUES (1, 'sally', '$2a$10$P8oijecN/DAmRxe2KOQUUurgKlbk1cyHYWfEVsVLODTD.HDiiQ.gO') ON CONFLICT username DO NOTHING;
/*
 User bank accounts
 */
INSERT INTO user_bank_accounts (user_id, plaid_access_token, plaid_item_id)
    VALUES (2, 'access-development-7b23962a-067c-497b-9d64-5814f2bae474', 'ewFlcQZzFfsYrVz63Np6cEm1wYMYgPUBDZnnd') ON CONFLICT (user_id, plaid_access_token, plaid_item_id)
    DO NOTHING;
INSERT INTO user_bank_accounts (user_id, plaid_access_token, plaid_item_id)
    VALUES (2, 'access-development-8a42966q-067c-497b-9d64-5814f2bae474', 'pdMiSTh3c0o1eSt63Np6cEm1wYMYgPUBDZnnd') ON CONFLICT (user_id, plaid_access_token, plaid_item_id)
    DO NOTHING;
