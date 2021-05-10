package gorippled

import (
	"fmt"
	"strings"
)

const (
	AccountChannels   = "account_channels"
	AccountCurrencies = "account_currencies"
	AccountInfo       = "account_info"
	AccountLines      = "account_lines"
	AccountObjects    = "account_objects"
	AccountOffers     = "account_offers"
	AccountTx         = "account_tx"
	GatewayBalances   = "gateway_balances"
	NorippleCheck     = "noripple_check"
	Ledger            = "ledger"
	LedgerClosed      = "ledger_closed"
	LedgerCurrent     = "ledger_current"
	LedgerData        = "ledger_data"
	LedgerEntry       = "ledger_entry"
	Sign              = "sign"
	SignFor           = "sign_for"
	Submit            = "submit"
	SubmitMultisigned = "submit_multisigned"
	TransactionEntry  = "transaction_entry"
	Tx                = "tx"
	TxHistory         = "tx_history"
	BookOffers        = "book_offers"
	DepositAuthorized = "deposit_authorized"
	PathFind          = "path_find"
	RipplePathFind    = "ripple_path_find"
	ChannelAuthorize  = "channel_authorize"
	ChannelVerify     = "channel_verify"
	Subscribe         = "subscribe"
	Unsubscribe       = "unsubscribe"
	Fee               = "fee"
	Manifest          = "manifest"
	ServerInfo        = "server_info"
	ServerState       = "server_state"
	JSON              = "json"
	Ping              = "ping"
	Random            = "random"

	CategoryAccount          = "Account"
	CategoryLedger           = "Ledger"
	CategoryTransaction      = "Transaction"
	CategoryPathAndOrderBook = "Path And Order Book"
	CategoryPaymentChannel   = "PaymentChannel"
	CategorySubscription     = "Subscription"
	CategoryServerInfo       = "Server Info"
	CategoryUtility          = "Utility"
)

func Methods() []Method {
	methods := []Method{}
	methods = append(methods, accountMethods...)
	methods = append(methods, ledgerMethods...)
	return methods
}

type Method struct {
	Name        string
	Category    string
	Summary     string
	Description string
}

func (m *Method) APIReferenceURL() string {
	m.Name = strings.TrimSpace(m.Name)
	if len(m.Name) > 0 {
		return fmt.Sprintf("https: //xrpl.org/%s.html", m.Name)
	}
	return ""
}

var accountMethods = []Method{
	{
		Name:        AccountChannels,
		Category:    CategoryAccount,
		Summary:     "Get a list of payment channels where the account is the source of the channel.",
		Description: "The `account_channels` method returns information about an account's Payment Channels. This includes only channels where the specified account is the channel's source, not the destination. (A channel's \"source\" and \"owner\" are the same.) All information retrieved is relative to a particular version of the ledger.",
	},
	{
		Name:        AccountCurrencies,
		Category:    CategoryAccount,
		Summary:     "Get a list of currencies an account can send or receive.",
		Description: "The `account_currencies` command retrieves a list of currencies that an account can send or receive, based on its trust lines. (This is not a thoroughly confirmed list, but it can be used to populate user interfaces.)",
	},
	{
		Name:        AccountInfo,
		Category:    CategoryAccount,
		Summary:     "Get basic data about an account.",
		Description: "The `account_info` command retrieves information about an account, its activity, and its XRP balance. All information retrieved is relative to a particular version of the ledger.",
	},
	{
		Name:        AccountLines,
		Category:    CategoryAccount,
		Summary:     "Get info about an account's trust lines.",
		Description: "The `account_lines` method returns information about an account's trust lines, including balances in all non-XRP currencies and assets. All information retrieved is relative to a particular version of the ledger.",
	},
	{
		Name:        AccountObjects,
		Category:    CategoryAccount,
		Summary:     "Get all ledger objects owned by an account.",
		Description: "The `account_objects` command returns the raw [ledger format](https://xrpl.org/ledger-object-types.html) for all objects owned by an account. For a higher-level view of an account's trust lines and balances, see the [`account_lines` method](https://xrpl.org/account_lines.html) instead.\n\nThe types of objects that may appear in the account_objects response for an account include:\n\n* [Offer objects](https://xrpl.org/offer.html) for orders that are currently live, unfunded, or expired but not yet removed. (See [Lifecycle of an Offer](https://xrpl.org/offers.html#lifecycle-of-an-offer) for more information.)\n* [RippleState objects](https://xrpl.org/ripplestate.html) for trust lines where this account's side is not in the default state.\n* The account's [SignerList](https://xrpl.org/signerlist.html), if the account has [multi-signing](https://xrpl.org/multi-signing.html) enabled.\n* [Escrow objects](https://xrpl.org/escrow.html) for held payments that have not yet been executed or canceled.\n* [PayChannel objects](https://xrpl.org/paychannel.html) for open payment channels.\n* [Check objects](https://xrpl.org/check.html) for pending Checks.\n* [DepositPreauth objects](https://xrpl.org/depositpreauth-object.html) for deposit preauthorizations. New in: rippled 1.1.0\n* [Ticket objects](https://xrpl.org/known-amendments.html#tickets) for Tickets.",
	},
	{
		Name:        AccountOffers,
		Category:    CategoryAccount,
		Summary:     "Get info about an account's currency exchange offers.",
		Description: "The `account_offers` method retrieves a list of [offers](https://xrpl.org/offers.html) made by a given [account](https://xrpl.org/accounts.html) that are outstanding as of a particular [ledger version](https://xrpl.org/ledgers.html).",
	},
	{
		Name:        AccountTx,
		Category:    CategoryAccount,
		Summary:     "Get info about an account's transactions.",
		Description: "The `account_tx` method retrieves a list of transactions that involved the specified account.",
	},
	{
		Name:        GatewayBalances,
		Category:    CategoryAccount,
		Summary:     "Calculate total amounts issued by an account.",
		Description: "The `gateway_balances` command calculates the total balances issued by a given account, optionally excluding amounts held by [operational addresses](https://xrpl.org/issuing-and-operational-addresses.html).",
	},
	{
		Name:        NorippleCheck,
		Category:    CategoryAccount,
		Summary:     "Get recommended changes to an account's Default Ripple and No Ripple settings.",
		Description: "The `noripple_check` command provides a quick way to check the status of [the Default Ripple field for an account and the No Ripple flag of its trust lines](https://xrpl.org/rippling.html), compared with the recommended settings.",
	},
}

var ledgerMethods = []Method{
	{
		Name:        Ledger,
		Category:    CategoryLedger,
		Summary:     "Get info about a ledger version.",
		Description: "Retrieve information about the public [ledger](https://xrpl.org/ledgers.html).",
	},
	{
		Name:        LedgerClosed,
		Category:    CategoryLedger,
		Summary:     "Get the latest closed ledger version.",
		Description: "The `ledger_closed` method returns the unique identifiers of the most recently closed ledger. (This ledger is not necessarily validated and immutable yet.)",
	},
	{
		Name:        LedgerCurrent,
		Category:    CategoryLedger,
		Summary:     "Get the current working ledger version.",
		Description: "The ledger_current method returns the unique identifiers of the current in-progress [ledger](https://xrpl.org/ledgers.html). This command is mostly useful for testing, because the ledger returned is still in flux.",
	},
	{
		Name:        LedgerData,
		Category:    CategoryLedger,
		Summary:     "Get the raw contents of a ledger version.",
		Description: "The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.",
	},
	{
		Name:        LedgerEntry,
		Category:    CategoryLedger,
		Summary:     "Get one element from a ledger version.",
		Description: "The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.",
	},
}
