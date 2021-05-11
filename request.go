package gorippled

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	ripplenetwork "github.com/goxrp/ripple-network"
	"github.com/grokify/simplego/net/http/httpsimple"
)

type JsonRpcRequest struct {
	Method string                   `json:"method"`
	Params []map[string]interface{} `json:"params"`
}

// BuildJsonRpcRequestBody merges a `rippled` API methid with JSON RPC API request params.
func BuildJsonRpcRequestBody(rippledMethod string, paramsBodyBytes []byte) (JsonRpcRequest, error) {
	jrpcReq := JsonRpcRequest{Method: rippledMethod}
	if len(paramsBodyBytes) > 0 {
		msi := map[string]interface{}{}
		err := json.Unmarshal(paramsBodyBytes, &msi)
		if err != nil {
			return jrpcReq, err
		}
		jrpcReq.Params = []map[string]interface{}{msi}
	}
	return jrpcReq, nil
}

// DoApiJsonRpcSplit sends a JSON RPC API request to the specified `rippled` server using
// an API method and request params.
func DoApiJsonRpcSplit(jrpcURL, rippledMethod string, paramsBodyBytes []byte) (*http.Response, error) {
	jrpcReq, err := BuildJsonRpcRequestBody(rippledMethod, paramsBodyBytes)
	if err != nil {
		return nil, err
	}

	return DoApiJsonRpc(jrpcURL, jrpcReq)
}

// DoApiJsonRpc sends a JSON RPC API request to the specified `rippled` server using
// an API method and request params.
func DoApiJsonRpc(jrpcURL string, reqBody interface{}) (*http.Response, error) {
	if len(strings.TrimSpace(jrpcURL)) == 0 {
		jrpcURL = ripplenetwork.GetMainnetPublicJsonRpcUrl()
	}

	return httpsimple.Do(httpsimple.SimpleRequest{
		Method: http.MethodPost,
		URL:    jrpcURL,
		Body:   reqBody,
		IsJSON: true})
}

var methodToAccount = map[string]string{
	"account_channels":             "account",
	"account_currencies":           "account",
	"account_info":                 "account",
	"account_lines":                "account",
	"ledger":                       "ledger",
	"ledger_closed":                "ledger",
	"ledger_current":               "ledger",
	"ledger_data":                  "ledger",
	"ledger_entry-account_root":    "ledger",
	"ledger_entry-check":           "ledger",
	"ledger_entry-deposit_preauth": "ledger",
	"ledger_entry-directory":       "ledger",
	"ledger_entry-escrow":          "ledger",
	"ledger_entry-index":           "ledger",
	"ledger_entry-offer":           "ledger",
	"ledger_entry-payment_channel": "ledger",
	"ledger_entry-ripple_state":    "ledger",
	"ledger_entry-ticket":          "ledger",
}

func MethodsPlusToAccount() map[string]string {
	return methodToAccount
}

func MethodsPlus() []string {
	methods := []string{}
	for k := range methodToAccount {
		methods = append(methods, k)
	}
	sort.Strings(methods)
	return methods
}
