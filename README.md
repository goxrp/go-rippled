# Go Rippled

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This package is designed to provide[`rippled` server](https://github.com/ripple/rippled) utilities including:

* OpenAPI specifcation: partial OAS3 specification files are provided
* Simple JSON RPC request handling

## Models

* [AccountRoot](spec.model.accountroot.json)
* [PayChannel](spec.model.paychannel.json)

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

 [build-status-svg]: https://github.com/go-xrp/go-rippled/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/go-xrp/go-rippled/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/go-xrp/go-rippled
 [goreport-url]: https://goreportcard.com/report/github.com/go-xrp/go-rippled
 [codeclimate-status-svg]: https://codeclimate.com/github/go-xrp/go-rippled/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/go-xrp/go-rippled
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/go-xrp/go-rippled
 [docs-godoc-url]: https://pkg.go.dev/github.com/go-xrp/go-rippled
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/go-xrp/go-rippled/blob/master/LICENSE