package gorippled

import (
	"fmt"
	"strings"
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

	MethodValidationCreate     = "validation_create"
	MethodWalletPropose        = "wallet_propose"
	MethodCanDelete            = "can_delete"
	MethodCrawlShards          = "crawl_shards"
	MethodDownloadShard        = "download_shard"
	MethodLedgerCleaner        = "ledger_cleaner"
	MethodLedgerRequest        = "ledger_request"
	MethodLogLevel             = "log_level"
	MethodLogrotate            = "logrotate"
	MethodLedgerAccept         = "ledger_accept"
	MethodStop                 = "stop"
	MethodValidationSeed       = "validation_seed"
	MethodConnect              = "connect"
	MethodPeerReservationsAdd  = "peer_reservations_add"
	MethodPeerReservationsDel  = "peer_reservations_del"
	MethodPeerReservationsList = "peer_reservations_list"
	MethodPeers                = "peers"
	MethodConsensusInfo        = "consensus_info"
	MethodFeature              = "feature"
	MethodFetchInfo            = "fetch_info"
	MethodGetCounts            = "get_counts"
	MethodPrint                = "print"
	MethodValidatorInfo        = "validator_info"
	MethodValidatorListSites   = "validator_list_sites"
	MethodValidators           = "validators"
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
	methods = append(methods, StatusAndDebuggingMethods()...)
	for i, method := range methods {
		cat, err := GetCategory(method.Category.Name)
		if err != nil {
			panic(fmt.Sprintf("category not found [%s]", method.Category.Name))
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
	HasApiCli       bool
	HasApiJsonRpc   bool
	HasApiWebsocket bool
}

func (m *Method) APIReferenceURL() string {
	m.Name = strings.TrimSpace(m.Name)
	if len(m.Name) > 0 {
		return fmt.Sprintf("https: //xrpl.org/%s.html", m.Name)
	}
	return ""
}

func MethodToCategory(methodName string) (string, error) {
	methods := Methods()
	for _, tryMethod := range methods {
		if tryMethod.Name == methodName {
			if len(tryMethod.Category.Name) > 0 {
				return tryMethod.Category.Name, nil
			} else {
				return "", fmt.Errorf("method has no category name: method [%s]", tryMethod.Name)
			}
		}
	}
	return "", fmt.Errorf("method not found for methodName[%s]", methodName)
}
