package gorippled

import (
	"encoding/json"
	"net/http"
	"strings"

	ripplenetwork "github.com/go-xrp/ripple-network"
	"github.com/grokify/simplego/net/http/httpsimple"
)

type JsonRpcRequest struct {
	Method string                   `json:"method"`
	Params []map[string]interface{} `json:"params"`
}

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

func DoApiJsonRpcSplit(jrpcURL, rippledMethod string, paramsBodyBytes []byte) (*http.Response, error) {
	jrpcReq, err := BuildJsonRpcRequestBody(rippledMethod, paramsBodyBytes)
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(jrpcURL)) == 0 {
		jrpcURL = ripplenetwork.GetMainnetPublicJsonRpcUrl()
	}

	return httpsimple.Do(httpsimple.SimpleRequest{
		Method: http.MethodPost,
		URL:    jrpcURL,
		Body:   jrpcReq,
		IsJSON: true})
}
