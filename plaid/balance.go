package plaid

import (
	"bytes"
	"encoding/json"
)

// Account.Numbers will be empty
type BalanceResponse struct {
	// Normal response fields
	AccessToken      string        `json:"access_token"`
	Accounts         []Account     `json:"accounts"`
}

// Balance (POST /balance) retrieves real-time balance for a given access token.
//
// See https://plaid.com/docs/api/#balance.
func (c *Client) Balance(accessToken string) (res *BalanceResponse, err error) {
	jsonText, err := json.Marshal(balanceJson{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, err
	}
	postRes, _, err := c.postAndUnmarshal("/balance", bytes.NewReader(jsonText))
	if err != nil  {
		return nil, err
	}
	return &BalanceResponse{
		AccessToken: postRes.AccessToken,
		Accounts: postRes.Accounts,
	}, err
}

type balanceJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
