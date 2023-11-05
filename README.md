# Go Rippled

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This package is designed to provide[`rippled` server](https://github.com/ripple/rippled) utilities including:

1. API Method and Category information
1. API example request and response JSON objects ([`data`](data) folder)
1. Simple JSON RPC request handling

## API Method and Category Information

Method and Category information is provided programmatically.

See:

* [`category.go`](category.go)
* [`method.go`](method.go)

## Example Requests and Responses

Request and response samples are provided in the `data` folder and available programmatically.

## Test Client

A test API client is provided that executes the requests on XRPL.org.

Display example request:

```
$ cd cmd/rippledapiexample
$ go run main.go --method account_info
```

Display example request and response:

```
$ cd cmd/rippledapiexample
$ go run main.go --method account_info --exec
```

## Models

* [AccountRoot](spec.model.accountroot.json)
* [PayChannel](spec.model.paychannel.json)

 [build-status-svg]: https://github.com/goxrp/go-rippled/workflows/test/badge.svg?branch=master
 [build-status-url]: https://github.com/goxrp/go-rippled/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/goxrp/go-rippled
 [goreport-url]: https://goreportcard.com/report/github.com/goxrp/go-rippled
 [codeclimate-status-svg]: https://codeclimate.com/github/goxrp/go-rippled/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/goxrp/go-rippled
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/goxrp/go-rippled
 [docs-godoc-url]: https://pkg.go.dev/github.com/goxrp/go-rippled
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/goxrp/go-rippled/blob/master/LICENSE