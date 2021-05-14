package gorippled

func AccountCurrenciesRequestSimple(account string) JsonRpcRequest {
	return JsonRpcRequest{
		Method: MethodAccountCurrencies,
		Params: []map[string]interface{}{{
			"account":       account,
			"account_index": 0,
			"ledger_index":  "validated",
			"strict":        true,
		}},
	}
}

type AccountCurrenciesResponse struct {
	Result AccountCurrenciesResult `json:"result"`
}

type AccountCurrenciesResult struct {
	LedgerHash        string   `json:"ledger_hash"`
	LedgerIndex       int64    `json:"ledger_index"`
	ReceiveCurrencies []string `json:"receive_currencies"`
	SendCurrencies    []string `json:"send_currencies"`
	Status            string   `json:"status"`
	Validated         bool     `json:"validated"`
}
