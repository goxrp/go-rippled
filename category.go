package gorippled

import (
	"fmt"
	"strings"

	"github.com/grokify/simplego/text/stringcase"
)

const (
	CategoryAccount                  = "account"
	CategoryLedger                   = "ledger"
	CategoryTransaction              = "transaction"
	CategoryPathAndOrderBook         = "path_and_order_book"
	CategoryPaymentChannel           = "payment_channel"
	CategorySubscription             = "subscription"
	CategoryServerInfo               = "server_info"
	CategoryUtility                  = "utility"
	CategoryKeyGeneration            = "key_generation"
	CategoryLoggingAndDataManagement = "logging_and_data_management"
	CategoryServerControl            = "server_control"
	CategoryPeerManagement           = "peer_management"
	CategoryStatusAndDebugging       = "status_and_debugging"

	TypePublic = "public"
	TypeAdmin  = "admin"
)

type Category struct {
	Name        string
	DisplayName string
	Type        string
	Description string
	IsReadOnly  bool
}

func (cat *Category) APIReferenceURL() string {
	cat.Name = strings.TrimSpace(cat.Name)
	if len(cat.Name) > 0 {
		return fmt.Sprintf("https: //xrpl.org/%s.html", stringcase.CaseSnakeToKebab(cat.Name))
	}
	return ""
}

func Categories() []Category {
	return []Category{
		{
			Name:        CategoryAccount,
			DisplayName: "Account",
			Type:        TypePublic,
			IsReadOnly:  true,
			Description: "An account in the XRP Ledger represents a holder of XRP and a sender of transactions. Use these methods to work with account info."},
		{
			Name:        CategoryLedger,
			DisplayName: "Ledger",
			Type:        TypePublic,
			IsReadOnly:  true,
			Description: "A ledger version contains a header, a transaction tree, and a state tree, which contain account settings, trustlines, balances, transactions, and other data. Use these methods to retrieve ledger info."},
		{
			Name:        CategoryTransaction,
			DisplayName: "Transaction",
			Type:        TypePublic,
			IsReadOnly:  false,
			Description: "Transactions are the only thing that can modify the shared state of the XRP Ledger. All business on the XRP Ledger takes the form of transactions. Use these methods to work with transactions."},
		{
			Name:        CategoryPathAndOrderBook,
			DisplayName: "Path and Order Book",
			Type:        TypePublic,
			IsReadOnly:  true,
			Description: "Paths define a way for payments to flow through intermediary steps on their way from sender to receiver. Paths enable cross-currency payments by connecting sender and receiver through order books. Use these methods to work with paths and other books."},
		{
			Name:        CategoryPaymentChannel,
			DisplayName: "Payment Channel",
			Type:        TypePublic,
			IsReadOnly:  false,
			Description: "Payment channels are a tool for facilitating repeated, unidirectional payments, or temporary credit between two parties. Use these methods to work with payment channels."},
		{
			Name:        CategorySubscription,
			DisplayName: "Subscription",
			Type:        TypePublic,
			IsReadOnly:  false,
			Description: "Use these methods to enable the server to push updates to your client when various events happen, so that you can know and react right away. WebSocket API only."},
		{
			Name:        CategoryServerInfo,
			DisplayName: "Server Info",
			Type:        TypePublic,
			IsReadOnly:  true,
			Description: "Use these methods to retrieve information about the current state of the rippled server."},
		{
			Name:        CategoryUtility,
			DisplayName: "Utility",
			Type:        TypePublic,
			IsReadOnly:  true,
			Description: "Use these methods to perform convenient tasks, such as ping and random number generation."},
		{
			Name:        CategoryKeyGeneration,
			DisplayName: "Key Generation",
			Type:        TypeAdmin,
			Description: "Use these methods to generate and manage keys."},
		{
			Name:        CategoryLoggingAndDataManagement,
			DisplayName: "Logging and Data Management",
			Type:        TypeAdmin,
			Description: "Use these methods to manage log levels and other data, such as ledgers."},
		{
			Name:        CategoryServerControl,
			DisplayName: "Server Control",
			Type:        TypeAdmin,
			Description: "Use these methods to manage the rippled server."},
		{
			Name:        CategoryPeerManagement,
			DisplayName: "Peer Management",
			Type:        TypeAdmin,
			Description: "Use these methods to manage your server's peer-to-peer connections."},
		{
			Name:        CategoryStatusAndDebugging,
			DisplayName: "Status and Debugging",
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
