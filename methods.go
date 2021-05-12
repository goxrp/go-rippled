package gorippled

import (
	"fmt"
	"strings"

	"github.com/grokify/simplego/database"
)

const (
	MethodAccountChannels   = "account_channels"
	MethodAccountCurrencies = "account_currencies"
	MethodAccountInfo       = "account_info"
	MethodAccountLines      = "account_lines"
	MethodAccountObjects    = "account_objects"
	MethodAccountOffers     = "account_offers"
	MethodAccountTx         = "account_tx"
	MethodGatewayBalances   = "gateway_balances"
	MethodNorippleCheck     = "noripple_check"
	MethodLedger            = "ledger"
	MethodLedgerClosed      = "ledger_closed"
	MethodLedgerCurrent     = "ledger_current"
	MethodLedgerData        = "ledger_data"
	MethodLedgerEntry       = "ledger_entry"
	MethodSign              = "sign"
	MethodSignFor           = "sign_for"
	MethodSubmit            = "submit"
	MethodSubmitMultisigned = "submit_multisigned"
	MethodTransactionEntry  = "transaction_entry"
	MethodTx                = "tx"
	MethodTxHistory         = "tx_history"
	MethodBookOffers        = "book_offers"
	MethodDepositAuthorized = "deposit_authorized"
	MethodPathFind          = "path_find"
	MethodRipplePathFind    = "ripple_path_find"
	MethodChannelAuthorize  = "channel_authorize"
	MethodChannelVerify     = "channel_verify"
	MethodSubscribe         = "subscribe"
	MethodUnsubscribe       = "unsubscribe"
	MethodFee               = "fee"
	MethodManifest          = "manifest"
	MethodServerInfo        = "server_info"
	MethodServerState       = "server_state"
	MethodJSON              = "json"
	MethodPing              = "ping"
	MethodRandom            = "random"

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

type MethodsSet struct {
	mapMethods      map[string]Method
	mapCategories   map[string]Category
	categoriesOrder []string
}

func NewMethodsSet() MethodsSet {
	set := MethodsSet{
		mapCategories:   map[string]Category{},
		mapMethods:      map[string]Method{},
		categoriesOrder: []string{}}
	methods := Methods()
	for _, method := range methods {
		method.Name = strings.TrimSpace(method.Name)
		if len(method.Name) == 0 {
			panic("empty method name")
		}
		set.mapMethods[method.Name] = method
	}
	categories := Categories()
	for _, category := range categories {
		category.Name = strings.TrimSpace(category.Name)
		if len(category.Name) == 0 {
			panic("empty category name")
		}
		set.categoriesOrder = append(set.categoriesOrder, category.Name)
		set.mapCategories[category.Name] = category
	}
	return set
}

func (set *MethodsSet) GetMethod(methodName string) (Method, error) {
	if m, ok := set.mapMethods[strings.TrimSpace(methodName)]; ok {
		return m, nil
	}
	return Method{}, fmt.Errorf("no method for [%s]", methodName)
}

func (set *MethodsSet) GetCategory(categoryName string) (Method, error) {
	if m, ok := set.mapMethods[strings.TrimSpace(categoryName)]; ok {
		return m, nil
	}
	return Method{}, fmt.Errorf("no category for [%s]", categoryName)
}

func Methods() []Method {
	methods := []Method{}
	methods = append(methods, AccountMethods()...)
	methods = append(methods, LedgerMethods()...)
	methods = append(methods, TransactionMethods()...)
	methods = append(methods, PathAndOrderBookMethods()...)
	methods = append(methods, PaymentChannelMethods()...)
	methods = append(methods, SubscriptionMethods()...)
	methods = append(methods, ServerInfoMethods()...)
	methods = append(methods, UtilityMethods()...)
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
	Name            string
	Category        Category
	Type            string
	Summary         string
	Description     string
	FunctionType    string
	IsOneOf         bool
	IsDeprecated    bool
	IsCLIOnly       bool
	IsWebsocketOnly bool
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

func GetCategoryMethods(categoryName string) ([]Method, error) {
	switch strings.ToLower(strings.TrimSpace(categoryName)) {
	case strings.ToLower(CategoryAccount):
		return AccountMethods(), nil
	case strings.ToLower(CategoryLedger):
		return LedgerMethods(), nil
	case strings.ToLower(CategoryTransaction):
		return TransactionMethods(), nil
	case strings.ToLower(CategoryPathAndOrderBook):
		return PathAndOrderBookMethods(), nil
	case strings.ToLower(CategoryPaymentChannel):
		return PaymentChannelMethods(), nil
	case strings.ToLower(CategorySubscription):
		return SubscriptionMethods(), nil
	case strings.ToLower(CategoryServerInfo):
		return ServerInfoMethods(), nil
	case strings.ToLower(CategoryUtility):
		return UtilityMethods(), nil
	default:
		return AccountMethods(), fmt.Errorf("category not found [%s]", categoryName)
	}
}

func AccountMethods() []Method {
	return []Method{
		{
			Name:         MethodAccountChannels,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get a list of payment channels where the account is the source of the channel.",
			Description:  "The `account_channels` method returns information about an account's Payment Channels. This includes only channels where the specified account is the channel's source, not the destination. (A channel's \"source\" and \"owner\" are the same.) All information retrieved is relative to a particular version of the ledger.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodAccountCurrencies,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get a list of currencies an account can send or receive.",
			Description:  "The `account_currencies` command retrieves a list of currencies that an account can send or receive, based on its trust lines. (This is not a thoroughly confirmed list, but it can be used to populate user interfaces.)",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodAccountInfo,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get basic data about an account.",
			Description:  "The `account_info` command retrieves information about an account, its activity, and its XRP balance. All information retrieved is relative to a particular version of the ledger.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodAccountLines,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get info about an account's trust lines.",
			Description:  "The `account_lines` method returns information about an account's trust lines, including balances in all non-XRP currencies and assets. All information retrieved is relative to a particular version of the ledger.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodAccountObjects,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get all ledger objects owned by an account.",
			Description:  "The `account_objects` command returns the raw [ledger format](https://xrpl.org/ledger-object-types.html) for all objects owned by an account. For a higher-level view of an account's trust lines and balances, see the [`account_lines` method](https://xrpl.org/account_lines.html) instead.\n\nThe types of objects that may appear in the account_objects response for an account include:\n\n* [Offer objects](https://xrpl.org/offer.html) for orders that are currently live, unfunded, or expired but not yet removed. (See [Lifecycle of an Offer](https://xrpl.org/offers.html#lifecycle-of-an-offer) for more information.)\n* [RippleState objects](https://xrpl.org/ripplestate.html) for trust lines where this account's side is not in the default state.\n* The account's [SignerList](https://xrpl.org/signerlist.html), if the account has [multi-signing](https://xrpl.org/multi-signing.html) enabled.\n* [Escrow objects](https://xrpl.org/escrow.html) for held payments that have not yet been executed or canceled.\n* [PayChannel objects](https://xrpl.org/paychannel.html) for open payment channels.\n* [Check objects](https://xrpl.org/check.html) for pending Checks.\n* [DepositPreauth objects](https://xrpl.org/depositpreauth-object.html) for deposit preauthorizations. New in: rippled 1.1.0\n* [Ticket objects](https://xrpl.org/known-amendments.html#tickets) for Tickets.",
			IsOneOf:      true,
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodAccountOffers,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get info about an account's currency exchange offers.",
			Description:  "The `account_offers` method retrieves a list of [offers](https://xrpl.org/offers.html) made by a given [account](https://xrpl.org/accounts.html) that are outstanding as of a particular [ledger version](https://xrpl.org/ledgers.html).",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodAccountTx,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get info about an account's transactions.",
			Description:  "The `account_tx` method retrieves a list of transactions that involved the specified account.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodGatewayBalances,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Calculate total amounts issued by an account.",
			Description:  "The `gateway_balances` command calculates the total balances issued by a given account, optionally excluding amounts held by [operational addresses](https://xrpl.org/issuing-and-operational-addresses.html).",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodNorippleCheck,
			Category:     Category{Name: CategoryAccount},
			Summary:      "Get recommended changes to an account's Default Ripple and No Ripple settings.",
			Description:  "The `noripple_check` command provides a quick way to check the status of [the Default Ripple field for an account and the No Ripple flag of its trust lines](https://xrpl.org/rippling.html), compared with the recommended settings.",
			FunctionType: database.FunctionRead,
		},
	}
}

func LedgerMethods() []Method {
	return []Method{
		{
			Name:         MethodLedger,
			Category:     Category{Name: CategoryLedger},
			Summary:      "Get info about a ledger version.",
			Description:  "Retrieve information about the public [ledger](https://xrpl.org/ledgers.html).",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodLedgerClosed,
			Category:     Category{Name: CategoryLedger},
			Summary:      "Get the latest closed ledger version.",
			Description:  "The `ledger_closed` method returns the unique identifiers of the most recently closed ledger. (This ledger is not necessarily validated and immutable yet.)",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodLedgerCurrent,
			Category:     Category{Name: CategoryLedger},
			Summary:      "Get the current working ledger version.",
			Description:  "The ledger_current method returns the unique identifiers of the current in-progress [ledger](https://xrpl.org/ledgers.html). This command is mostly useful for testing, because the ledger returned is still in flux.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodLedgerData,
			Category:     Category{Name: CategoryLedger},
			Summary:      "Get the raw contents of a ledger version.",
			Description:  "The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodLedgerEntry,
			Category:     Category{Name: CategoryLedger},
			Summary:      "Get one element from a ledger version.",
			Description:  "The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.",
			IsOneOf:      true,
			FunctionType: database.FunctionRead,
		},
	}
}

func TransactionMethods() []Method {
	return []Method{
		{
			Name:         MethodSign,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Cryptographically sign a transaction.",
			Description:  "The `sign` method takes a [transaction in JSON format](https://xrpl.org/transaction-formats.html) and a [seed value](https://xrpl.org/cryptographic-keys.html), and returns a signed binary representation of the transaction. To contribute one signature to a [multi-signed transaction](https://xrpl.org/multi-signing.html), use the [`sign_for` method](https://xrpl.org/sign_for.html) instead.\n\nBy default, this method is [admin-only](https://xrpl.org/admin-rippled-methods.html). It can be used as a public method if the server has [enabled public signing](https://xrpl.org/enable-public-signing.html).",
			FunctionType: database.FunctionAdd,
		},
		{
			Name:         MethodSignFor,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Contribute to a multi-signature.",
			Description:  "The `sign_for` command provides one signature for a [multi-signed transaction](https://xrpl.org/multi-signing.html).\n\nBy default, this method is [admin-only](https://xrpl.org/admin-rippled-methods.html). It can be used as a public method if the server has [enabled public signing](https://xrpl.org/enable-public-signing.html).\n\nThis command requires the [MultiSign amendment](https://xrpl.org/known-amendments.html#multisign) to be enabled.",
			FunctionType: database.FunctionAdd,
		},
		{
			Name:         MethodSubmit,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Send a transaction to the network.",
			Description:  "The `submit` method applies a [transaction](https://xrpl.org/transaction-formats.html) and sends it to the network to be confirmed and included in future ledgers.\n\nThis command has two modes:\n\n* Submit-only mode takes a signed, serialized transaction as a binary blob, and submits it to the network as-is. Since signed transaction objects are immutable, no part of the transaction can be modified or automatically filled in after submission.\n* Sign-and-submit mode takes a JSON-formatted Transaction object, completes and signs the transaction in the same manner as the sign method, and then submits the signed transaction. We recommend only using this mode for testing and development.\n\nTo send a transaction as robustly as possible, you should construct and sign it in advance, persist it somewhere that you can access even after a power outage, then `submit` it as a `tx_blob`. After submission, monitor the network with the [`tx` method](https://xrpl.org/tx.html) command to see if the transaction was successfully applied; if a restart or other problem occurs, you can safely re-submit the `tx_blob` transaction: it won't be applied twice since it has the same sequence number as the old transaction.",
			FunctionType: database.FunctionAdd,
		},
		{
			Name:         MethodSubmitMultisigned,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Send a multi-signed transaction to the network.",
			Description:  "The `submit_multisigned` command applies a [multi-signed](https://xrpl.org/multi-signing.html) transaction and sends it to the network to be included in future ledgers. (You can also submit multi-signed transactions in binary form using the [submit command in submit-only mode](https://xrpl.org/submit.html#submit-only-mode).)\n\nThis command requires the [MultiSign amendment](https://xrpl.org/known-amendments.html#multisign) to be enabled.",
			FunctionType: database.FunctionAdd,
		},
		{
			Name:         MethodTransactionEntry,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Retrieve info about a transaction from a particular ledger version.",
			Description:  "The `transaction_entry` method retrieves information on a single transaction from a specific ledger version. (The [`tx` method](https://xrpl.org/tx.html), by contrast, searches all ledgers for the specified transaction. We recommend using that method instead.)",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodTx,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Retrieve info about a transaction from all the ledgers on hand.",
			Description:  "The `tx` method retrieves information on a single [transaction](https://xrpl.org/transaction-formats.html), by its [identifying hash](https://xrpl.org/transaction-basics.html#identifying-transactions).",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodTxHistory,
			Category:     Category{Name: CategoryTransaction},
			Summary:      "Retrieve info about all recent transactions.",
			Description:  "The `tx_history` method retrieves some of the most recent transactions made.",
			FunctionType: database.FunctionRead,
			IsDeprecated: true,
		},
	}
}

func PathAndOrderBookMethods() []Method {
	return []Method{
		{
			Name:         MethodBookOffers,
			Category:     Category{Name: CategoryPathAndOrderBook},
			Summary:      "Get info about offers to exchange two currencies.",
			Description:  "The `book_offers` method retrieves a list of offers, also known as the [order book](https://www.investopedia.com/terms/o/order-book.asp), between two currencies.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodDepositAuthorized,
			Category:     Category{Name: CategoryPathAndOrderBook},
			Summary:      "Check whether an account is authorized to send money directly to another.",
			Description:  "The `deposit_authorized` command indicates whether one account is authorized to send payments directly to another. See [Deposit Authorization](https://xrpl.org/depositauth.html) for information on how to require authorization to deliver money to your account.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:            MethodPathFind,
			Category:        Category{Name: CategoryPathAndOrderBook},
			Summary:         "Find a path for a payment between two accounts and receive updates.",
			Description:     "The `path_find` method searches for a [path](https://xrpl.org/paths.html) along which a transaction can possibly be made, and periodically sends updates when the path changes over time. For a simpler version that is supported by JSON-RPC, see the [`ripple_path_find` method](https://xrpl.org/ripple_path_find.html). For payments occurring strictly in XRP, it is not necessary to find a path, because XRP can be sent directly to any account.",
			FunctionType:    database.FunctionRead,
			IsWebsocketOnly: true,
		},
		{
			Name:         MethodRipplePathFind,
			Category:     Category{Name: CategoryPathAndOrderBook},
			Summary:      "Find a path for payment between two accounts, once.",
			Description:  "The `ripple_path_find` method is a simplified version of the [`path_find` method](https://xrpl.org/path_find.html) that provides a single response with a [payment path](https://xrpl.org/paths.html) you can use right away. It is available in both the WebSocket and JSON-RPC APIs. However, the results tend to become outdated as time passes. Instead of making multiple calls to stay updated, you should instead use the [`path_find` method](https://xrpl.org/path_find.html) to subscribe to continued updates where possible.\n\nAlthough the `rippled` server tries to find the cheapest path or combination of paths for making a payment, it is not guaranteed that the paths returned by this method are, in fact, the best paths.",
			FunctionType: database.FunctionRead,
		},
	}
}

func PaymentChannelMethods() []Method {
	return []Method{
		{
			Name:         MethodChannelAuthorize,
			Category:     Category{Name: CategoryPaymentChannel},
			Summary:      "Sign a claim for money from a payment channel.",
			Description:  "The `channel_authorize` method creates a signature that can be used to redeem a specific amount of XRP from a payment channel.",
			FunctionType: database.FunctionAdd,
		},
		{
			Name:         MethodChannelVerify,
			Category:     Category{Name: CategoryPaymentChannel},
			Summary:      "Check a payment channel claim's signature.",
			Description:  "The `channel_verify` method checks the validity of a signature that can be used to redeem a specific amount of XRP from a payment channel.",
			FunctionType: database.FunctionRead,
		},
	}
}

func SubscriptionMethods() []Method {
	return []Method{
		{
			Name:         MethodSubscribe,
			Category:     Category{Name: CategorySubscription},
			Summary:      "Listen for updates about a particular subject.",
			Description:  "The `subscribe` method requests periodic notifications from the server when certain events happen.",
			FunctionType: database.FunctionAdd,
		},
		{
			Name:         MethodUnsubscribe,
			Category:     Category{Name: CategorySubscription},
			Summary:      "Stop listening for updates about a particular subject.",
			Description:  "The `unsubscribe` command tells the server to stop sending messages for a particular subscription or set of subscriptions.",
			FunctionType: database.FunctionDelete,
		},
	}
}

func ServerInfoMethods() []Method {
	return []Method{
		{
			Name:         MethodFee,
			Category:     Category{Name: CategoryServerInfo},
			Summary:      "Get information about transaction cost.",
			Description:  "The `fee` command reports the current state of the open-ledger requirements for the [transaction cost](https://xrpl.org/transaction-cost.html). This requires the FeeEscalation amendment to be enabled.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodManifest,
			Category:     Category{Name: CategoryServerInfo},
			Summary:      "Look up the public information about a known validator.",
			Description:  "The `manifest` method reports the current \"manifest\" information for a given validator public key. The \"manifest\" is the public portion of that validator's configured token.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodServerInfo,
			Category:     Category{Name: CategoryServerInfo},
			Summary:      "Retrieve status of the server in human-readable format.",
			Description:  "The `server_info` command asks the server for a human-readable version of various information about the rippled server being queried.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodServerState,
			Category:     Category{Name: CategoryServerInfo},
			Summary:      "Retrieve status of the server in machine-readable format.",
			Description:  "The `server_state` command asks the server for various machine-readable information about the rippled server's current state. The response is almost the same as the [`server_info` method](https://xrpl.org/server_info.html), but uses units that are easier to process instead of easier to read. (For example, XRP values are given in integer drops instead of scientific notation or decimal values, and time is given in milliseconds instead of seconds.)",
			FunctionType: database.FunctionRead,
		},
	}
}

func UtilityMethods() []Method {
	return []Method{
		{
			Name:         MethodJSON,
			Category:     Category{Name: CategoryUtility},
			Summary:      "Pass JSON through the commandline.",
			Description:  "The `json` method is a proxy to running other commands, and accepts the parameters for the command as a JSON value. It is *exclusive to the Commandline client*, and intended for cases where the commandline syntax for specifying parameters is inadequate or undesirable.",
			IsCLIOnly:    true,
			FunctionType: database.FunctionOther,
		},
		{
			Name:         MethodPing,
			Category:     Category{Name: CategoryUtility},
			Summary:      "Confirm connectivity with the server.",
			Description:  "The `ping` command returns an acknowledgement, so that clients can test the connection status and latency.",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodRandom,
			Category:     Category{Name: CategoryUtility},
			Summary:      "Generate a random number.",
			Description:  "The `random` command provides a random number to be used as a source of entropy for random number generation by clients.",
			FunctionType: database.FunctionRead,
		},
	}
}
