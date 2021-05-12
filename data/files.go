package data

import (
	"fmt"
	"path/filepath"
	"strings"

	gorippled "github.com/goxrp/go-rippled"
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

func GetMethodCategory(methodName string) string {
	categoryName, err := gorippled.MethodToCategory(methodName)
	if err != nil {
		return ""
	}
	return categoryName
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
