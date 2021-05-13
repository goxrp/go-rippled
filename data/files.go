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

func ExampleJsonRequestFilename(methodName, categoryName string) (string, error) {
	categoryName = strings.ToLower(strings.TrimSpace(categoryName))
	if len(categoryName) == 0 {
		categoryName = GetMethodCategory(methodName)
	}
	if len(categoryName) == 0 {
		return "", fmt.Errorf("no category for method [%s]", methodName)
	}
	cat, err := gorippled.GetCategory(categoryName)
	if err != nil {
		return "", fmt.Errorf("no category for categoryName [%s]", categoryName)
	}
	if cat.Type == gorippled.TypeAdmin {
		return fmt.Sprintf("method.admin.%s.%s.jsonrpc.request.json", categoryName, methodName), nil
	}
	return fmt.Sprintf("method.%s.%s.jsonrpc.request.json", categoryName, methodName), nil
}

func ExampleJsonRequest(method, category string) ([]byte, error) {
	filename, err := ExampleJsonRequestFilename(method, category)
	if err != nil {
		return []byte{}, err
	}
	return Asset(filename)
}
