package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/jessevdk/go-flags"

	"github.com/go-xrp/go-rippled/data"
)

type Options struct {
	Method string `short:"m" long:"method" description:"Method" required:"true"`
	Exec   []bool `short:"x" long:"exec" description:"Execute API Call"`
	Pretty []bool `short:"p" long:"pretty" description:"Pretty Print Result"`
}

const (
	RippledJsonRpcUrl = "https://s1.ripple.com:51234/"
)

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := data.ExampleJsonRequest(opts.Method)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))

	if len(opts.Exec) > 0 {
		req := httpsimple.SimpleRequest{
			Method: http.MethodPost,
			URL:    RippledJsonRpcUrl,
			Body:   bytes,
			IsJSON: true}

		resp, err := httpsimple.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonutil.PrettyPrint(bytes, "", "  ")))
	}
	fmt.Println("DONE")
}
