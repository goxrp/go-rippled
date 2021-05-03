package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/jessevdk/go-flags"

	gorippled "github.com/go-xrp/go-rippled"
	"github.com/go-xrp/go-rippled/data"
)

type Options struct {
	Method   string `short:"m" long:"method" description:"Method" required:"true"`
	Category string `short:"c" long:"category" description:"Category"`
	Exec     []bool `short:"x" long:"exec" description:"Execute API Call"`
	Pretty   []bool `short:"p" long:"pretty" description:"Pretty Print Result"`
}

const (
	RippledJsonRpcUrl = "https://s1.ripple.com:51234/"
)

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		methods := gorippled.MethodsPlus()
		fmt.Printf("No Method Provided. Valid Methods [%s]\n", strings.Join(methods, ","))
		log.Fatal(err)
	}

	bytes, err := data.ExampleJsonRequest(opts.Method, opts.Category)
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

		fmtutil.PrintJSON(req)

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
