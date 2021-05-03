package data

import (
	"fmt"
	"path/filepath"
	"strings"

	gorippled "github.com/go-xrp/go-rippled"
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
	return filepath.Join("./", filename)
}

func GetMethodCategory(method string) string {
	methodToAccount := gorippled.MethodsPlusToAccount()
	method = strings.ToLower(strings.TrimSpace(method))
	if category, ok := methodToAccount[method]; ok {
		return strings.ToLower(strings.TrimSpace(category))
	}
	return ""
}

func ExampleJsonRequestFilename(method, category string) (string, error) {
	category = strings.ToLower(strings.TrimSpace(category))
	if len(category) == 0 {
		category = GetMethodCategory(method)
	}
	if len(category) == 0 {
		return "", fmt.Errorf("no category for method [%s]", method)
	}
	return fmt.Sprintf("method.%s.%s.jsonrpc.request.json", category, method), nil
}

func ExampleJsonRequest(method, category string) ([]byte, error) {
	filename, err := ExampleJsonRequestFilename(method, category)
	if err != nil {
		return []byte{}, err
	}
	return Asset(filename)
}
