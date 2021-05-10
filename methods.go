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

	CategoryAccount                  = "Account"
	CategoryLedger                   = "Ledger"
	CategoryTransaction              = "Transaction"
	CategoryPathAndOrderBook         = "Path And Order Book"
	CategoryPaymentChannel           = "PaymentChannel"
	CategorySubscription             = "Subscription"
	CategoryServerInfo               = "Server Info"
	CategoryUtility                  = "Utility"
	CategoryKeyGeneration            = "Key Generation"
	CategoryLoggingAndDataManagement = "Logging and Data Management"
	CategoryServerControl            = "Server Control"
	CategoryPeerManagement           = "Peer Management"
	CategoryStatusAndDebugging       = "Status and Debugging"

	TypePublic = "public"
	TypeAdmin  = "admin"
)

func Methods() []Method {
	methods := []Method{}
	methods = append(methods, accountMethods...)
	methods = append(methods, ledgerMethods...)
	methods = append(methods, paymentChannelMethods...)
	methods = append(methods, serverInfoMethods...)
	for i, method := range methods {
		cat, err := GetCategory(method.Category.Name)
		if err != nil {
			panic(fmt.Sprintf("category not found [%s]", method.Category))
		}
		method.Category = cat
		methods[i] = method
	}
	return methods
}

type Method struct {
	Name        string
	Category    Category
	Type        string
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

type Category struct {
	Name        string
	Type        string
	Description string
}

func Categories() []Category {
	return []Category{
		{
			Name:        CategoryAccount,
			Type:        TypePublic,
			Description: "An account in the XRP Ledger represents a holder of XRP and a sender of transactions. Use these methods to work with account info."},
		{
			Name:        CategoryLedger,
			Type:        TypePublic,
			Description: "A ledger version contains a header, a transaction tree, and a state tree, which contain account settings, trustlines, balances, transactions, and other data. Use these methods to retrieve ledger info."},
		{
			Name:        CategoryTransaction,
			Type:        TypePublic,
			Description: "Transactions are the only thing that can modify the shared state of the XRP Ledger. All business on the XRP Ledger takes the form of transactions. Use these methods to work with transactions."},
		{
			Name:        CategoryPathAndOrderBook,
			Type:        TypePublic,
			Description: "Paths define a way for payments to flow through intermediary steps on their way from sender to receiver. Paths enable cross-currency payments by connecting sender and receiver through order books. Use these methods to work with paths and other books."},
		{
			Name:        CategoryPaymentChannel,
			Type:        TypePublic,
			Description: "Payment channels are a tool for facilitating repeated, unidirectional payments, or temporary credit between two parties. Use these methods to work with payment channels."},
		{
			Name:        CategorySubscription,
			Type:        TypePublic,
			Description: "Use these methods to enable the server to push updates to your client when various events happen, so that you can know and react right away. WebSocket API only."},
		{
			Name:        CategoryServerInfo,
			Type:        TypePublic,
			Description: "Use these methods to retrieve information about the current state of the rippled server."},
		{
			Name:        CategoryUtility,
			Type:        TypePublic,
			Description: "Use these methods to perform convenient tasks, such as ping and random number generation."},
		{
			Name:        CategoryKeyGeneration,
			Type:        TypeAdmin,
			Description: "Use these methods to generate and manage keys."},
		{
			Name:        CategoryLoggingAndDataManagement,
			Type:        TypeAdmin,
			Description: "Use these methods to manage log levels and other data, such as ledgers."},
		{
			Name:        CategoryServerControl,
			Type:        TypeAdmin,
			Description: "Use these methods to manage the rippled server."},
		{
			Name:        CategoryPeerManagement,
			Type:        TypeAdmin,
			Description: "Use these methods to manage your server's peer-to-peer connections."},
		{
			Name:        CategoryStatusAndDebugging,
			Type:        TypeAdmin,
			Description: "Use these methods to check the status of the network and server."},
	}
}

func GetCategory(categoryName string) (Category, error) {
	cats := Categories()
	for _, cat := range cats {
		if categoryName == cat.Name {
			return cat, nil
		}
	}
	return Category{}, fmt.Errorf("category not found [%s]", categoryName)
}

var accountMethods = []Method{
	{
		Name:        AccountChannels,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get a list of payment channels where the account is the source of the channel.",
		Description: "The `account_channels` method returns information about an account's Payment Channels. This includes only channels where the specified account is the channel's source, not the destination. (A channel's \"source\" and \"owner\" are the same.) All information retrieved is relative to a particular version of the ledger.",
	},
	{
		Name:        AccountCurrencies,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get a list of currencies an account can send or receive.",
		Description: "The `account_currencies` command retrieves a list of currencies that an account can send or receive, based on its trust lines. (This is not a thoroughly confirmed list, but it can be used to populate user interfaces.)",
	},
	{
		Name:        AccountInfo,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get basic data about an account.",
		Description: "The `account_info` command retrieves information about an account, its activity, and its XRP balance. All information retrieved is relative to a particular version of the ledger.",
	},
	{
		Name:        AccountLines,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get info about an account's trust lines.",
		Description: "The `account_lines` method returns information about an account's trust lines, including balances in all non-XRP currencies and assets. All information retrieved is relative to a particular version of the ledger.",
	},
	{
		Name:        AccountObjects,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get all ledger objects owned by an account.",
		Description: "The `account_objects` command returns the raw [ledger format](https://xrpl.org/ledger-object-types.html) for all objects owned by an account. For a higher-level view of an account's trust lines and balances, see the [`account_lines` method](https://xrpl.org/account_lines.html) instead.\n\nThe types of objects that may appear in the account_objects response for an account include:\n\n* [Offer objects](https://xrpl.org/offer.html) for orders that are currently live, unfunded, or expired but not yet removed. (See [Lifecycle of an Offer](https://xrpl.org/offers.html#lifecycle-of-an-offer) for more information.)\n* [RippleState objects](https://xrpl.org/ripplestate.html) for trust lines where this account's side is not in the default state.\n* The account's [SignerList](https://xrpl.org/signerlist.html), if the account has [multi-signing](https://xrpl.org/multi-signing.html) enabled.\n* [Escrow objects](https://xrpl.org/escrow.html) for held payments that have not yet been executed or canceled.\n* [PayChannel objects](https://xrpl.org/paychannel.html) for open payment channels.\n* [Check objects](https://xrpl.org/check.html) for pending Checks.\n* [DepositPreauth objects](https://xrpl.org/depositpreauth-object.html) for deposit preauthorizations. New in: rippled 1.1.0\n* [Ticket objects](https://xrpl.org/known-amendments.html#tickets) for Tickets.",
	},
	{
		Name:        AccountOffers,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get info about an account's currency exchange offers.",
		Description: "The `account_offers` method retrieves a list of [offers](https://xrpl.org/offers.html) made by a given [account](https://xrpl.org/accounts.html) that are outstanding as of a particular [ledger version](https://xrpl.org/ledgers.html).",
	},
	{
		Name:        AccountTx,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get info about an account's transactions.",
		Description: "The `account_tx` method retrieves a list of transactions that involved the specified account.",
	},
	{
		Name:        GatewayBalances,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Calculate total amounts issued by an account.",
		Description: "The `gateway_balances` command calculates the total balances issued by a given account, optionally excluding amounts held by [operational addresses](https://xrpl.org/issuing-and-operational-addresses.html).",
	},
	{
		Name:        NorippleCheck,
		Category:    Category{Name: CategoryAccount},
		Summary:     "Get recommended changes to an account's Default Ripple and No Ripple settings.",
		Description: "The `noripple_check` command provides a quick way to check the status of [the Default Ripple field for an account and the No Ripple flag of its trust lines](https://xrpl.org/rippling.html), compared with the recommended settings.",
	},
}

var ledgerMethods = []Method{
	{
		Name:        Ledger,
		Category:    Category{Name: CategoryLedger},
		Summary:     "Get info about a ledger version.",
		Description: "Retrieve information about the public [ledger](https://xrpl.org/ledgers.html).",
	},
	{
		Name:        LedgerClosed,
		Category:    Category{Name: CategoryLedger},
		Summary:     "Get the latest closed ledger version.",
		Description: "The `ledger_closed` method returns the unique identifiers of the most recently closed ledger. (This ledger is not necessarily validated and immutable yet.)",
	},
	{
		Name:        LedgerCurrent,
		Category:    Category{Name: CategoryLedger},
		Summary:     "Get the current working ledger version.",
		Description: "The ledger_current method returns the unique identifiers of the current in-progress [ledger](https://xrpl.org/ledgers.html). This command is mostly useful for testing, because the ledger returned is still in flux.",
	},
	{
		Name:        LedgerData,
		Category:    Category{Name: CategoryLedger},
		Summary:     "Get the raw contents of a ledger version.",
		Description: "The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.",
	},
	{
		Name:        LedgerEntry,
		Category:    Category{Name: CategoryLedger},
		Summary:     "Get one element from a ledger version.",
		Description: "The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.",
	},
}

var paymentChannelMethods = []Method{
	{
		Name:        ChannelAuthorize,
		Category:    Category{Name: CategoryPaymentChannel},
		Summary:     "Sign a claim for money from a payment channel.",
		Description: "The `channel_authorize` method creates a signature that can be used to redeem a specific amount of XRP from a payment channel.",
	},
	{
		Name:        ChannelVerify,
		Category:    Category{Name: CategoryPaymentChannel},
		Summary:     "Check a payment channel claim's signature.",
		Description: "The `channel_verify` method checks the validity of a signature that can be used to redeem a specific amount of XRP from a payment channel.",
	},
}

var subscriptionMethods = []Method{
	{
		Name:        Subscribe,
		Category:    Category{Name: CategorySubscription},
		Summary:     "Listen for updates about a particular subject.",
		Description: "The `subscribe` method requests periodic notifications from the server when certain events happen.",
	},
	{
		Name:        Unsubscribe,
		Category:    Category{Name: CategorySubscription},
		Summary:     "Stop listening for updates about a particular subject.",
		Description: "The `unsubscribe` command tells the server to stop sending messages for a particular subscription or set of subscriptions.",
	},
}

var serverInfoMethods = []Method{
	{
		Name:        Fee,
		Category:    Category{Name: CategoryServerInfo},
		Summary:     "Get information about transaction cost.",
		Description: "The `fee` command reports the current state of the open-ledger requirements for the [transaction cost](https://xrpl.org/transaction-cost.html). This requires the FeeEscalation amendment to be enabled.",
	},
	{
		Name:        Manifest,
		Category:    Category{Name: CategoryServerInfo},
		Summary:     "Look up the public information about a known validator.",
		Description: "The `manifest` method reports the current \"manifest\" information for a given validator public key. The \"manifest\" is the public portion of that validator's configured token.",
	},
	{
		Name:        ServerInfo,
		Category:    Category{Name: CategoryServerInfo},
		Summary:     "Retrieve status of the server in human-readable format.",
		Description: "The `server_info` command asks the server for a human-readable version of various information about the rippled server being queried.",
	},
	{
		Name:        ServerState,
		Category:    Category{Name: CategoryServerInfo},
		Summary:     "Retrieve status of the server in machine-readable format.",
		Description: "The `server_state` command asks the server for various machine-readable information about the rippled server's current state. The response is almost the same as the [`server_info` method](https://xrpl.org/server_info.html), but uses units that are easier to process instead of easier to read. (For example, XRP values are given in integer drops instead of scientific notation or decimal values, and time is given in milliseconds instead of seconds.)",
	},
}
