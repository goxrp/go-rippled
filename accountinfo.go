package gorippled

/*
{
    "method": "account_info",
    "params": [
        {
            "account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
            "strict": true,
            "ledger_index": "current",
            "queue": true
        }
    ]
}
*/

type AccountInfoResponse struct {
	Result AccountInfoResult `json:"result"`
}

type AccountInfoResult struct {
	AccountData       AccountRoot `json:"account_data"`
	LedgerCurrentData string      `json:"ledger_current_index"`
	Status            string      `json:"status"`
	Validated         bool        `json:"validated"`
}

type AccountRoot struct {
	Account           string
	AccountTxnID      string `json:"AccountTxnID,omitempty"`
	Balance           string
	Domain            string `json:"Domain,omitempty"`
	EmailHash         string `json:"EmailHash,omitempty"`
	Flags             uint32
	LedgerEntryType   string
	MessageKey        string `json:"MessageKey,omitempty"`
	OwnerCount        uint32
	PreviousTxnID     string
	PreviousTxnLgrSeq uint32
	RegularKey        string `json:"RegularKey,omitempty"`
	Sequence          uint32
	TicketCount       uint32 `json:"TicketCount,omitempty"`
	TickSize          uint8  `json:"TickSize,omitempty"`
	TransferRate      uint32 `json:"TransferRate,omitempty"`
	WalletLocator     string `json:"WalletLocator,omitempty"`
	WalletSize        uint32 `json:"WalletSize,omitempty"`
	Index             string `json:"index"`
	Urlgravatar       string `json:"urlgravatar"`
}
