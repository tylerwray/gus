package app

// AddBankAccount adds a bank account for a user
func (s *Service) AddBankAccount(userID, publicToken string) error {
	res, err := s.plaid.ExchangePublicToken(publicToken)

	if err != nil {
		return err
	}

	if _, err := s.db.Exec("INSERT INTO user_bank_accounts (user_id, plaid_access_token, plaid_item_id) VALUES ($1, $2, $3)", userID, res.AccessToken, res.ItemID); err != nil {
		return err
	}

	return nil
}
