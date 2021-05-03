package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	gorippled "github.com/go-xrp/go-rippled"
)

// TestRequest tests API requests.
func TestRequest(t *testing.T) {
	if os.Getenv("GORIPPLED_TEST") != "LIVE" {
		return
	}
	methodsPlus := gorippled.MethodsPlus()
	n := len(methodsPlus)
	for i, methodPlus := range methodsPlus {
		if methodPlus == "ledger_entry-ticket" {
			continue
		}
		reqBodyBytes, err := ExampleJsonRequest(methodPlus, "")
		if err != nil {
			t.Errorf("[%d/%d] data.ExampleJsonRequest(\"%s\",\"\") No data err [%s]",
				i, n, methodPlus, err.Error())
		}

		resp, err := gorippled.DoApiJsonRpc("", reqBodyBytes)
		if err != nil {
			t.Errorf("[%d/%d] gorippled.DoApiJsonRpcSplit(\"\",\"%s\",...) Response err [%s]",
				i, n, methodPlus, err.Error())
		}

		respBodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("[%d/%d] ioutil.ReadAll(...) err [%s]", i, n, err.Error())
		}

		respBody := gorippled.JsonRpcResponseError{}
		err = json.Unmarshal(respBodyBytes, &respBody)
		if err != nil {
			t.Errorf("[%d/%d] json.Unmarshal(..., ...) err [%s]", i, n, err.Error())
		}
		tryStatus := respBody.Result.Status

		if tryStatus != gorippled.StatusSuccess {
			fmt.Println(string(respBodyBytes))
			t.Errorf("[%d/%d] gorippled.DoApiJsonRpcSplit(\"\",\"%s\",...) Bad status [%s]",
				i, n, methodPlus, tryStatus)
			panic("Z ")
		}
	}
}
