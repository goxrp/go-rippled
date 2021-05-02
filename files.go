package rippledopenapispec

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/grokify/rippled-openapi-spec/docs/examples"
)

const (
	ApiTypeJsonRpc   = "jsonrpc"
	ApiTypeWebsocket = "websocket"

	ObjectTypeRequest  = "request"
	ObjectTypeResponse = "response"
	ObjectTypeModel    = "model"
)

const (
	MethodAccountInfoJsonrpcRequest = "method.account.account_info.jsonrpc.request.json"
)

func examplespath(filename string) string {
	return filepath.Join("docs/examples", filename)
}

var methodToAccount = map[string]string{
	"account_channels":   "account",
	"account_currencies": "account",
	"account_info":       "account",
	"account_lines":      "account",
	"ledger":             "ledger",
	"ledger_closed":      "ledger",
	"ledger_current":     "ledger",
	"ledger_data":        "ledger",
}

func GetMethodCategory(method string) string {
	method = strings.ToLower(strings.TrimSpace(method))
	if category, ok := methodToAccount[method]; ok {
		return strings.ToLower(strings.TrimSpace(category))
	}
	return ""
}

func ExampleJsonRequestFilename(method string) (string, error) {
	category := GetMethodCategory(method)
	if len(category) == 0 {
		return "", fmt.Errorf("no category for method [%s]", method)
	}
	return fmt.Sprintf("method.%s.%s.jsonrpc.request.json", category, method), nil
}

func ExampleJsonRequest(method string) ([]byte, error) {
	filename, err := ExampleJsonRequestFilename(method)
	if err != nil {
		return []byte{}, err
	}
	return examples.Asset(filename)
}
