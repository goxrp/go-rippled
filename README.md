<img align="right" width="200" src="https://s2.coinmarketcap.com/static/img/coins/200x200/52.png" />

# Rippled OpenAPI

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This package is designed to provide OpenAPI specification defintions for the [`rippled` server](https://github.com/ripple/rippled).

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

 [build-status-svg]: https://github.com/go-xrp/rippled-openapi/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/go-xrp/rippled-openapi/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/go-xrp/rippled-openapi
 [goreport-url]: https://goreportcard.com/report/github.com/go-xrp/rippled-openapi
 [codeclimate-status-svg]: https://codeclimate.com/github/go-xrp/rippled-openapi/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/go-xrp/rippled-openapi
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/go-xrp/rippled-openapi
 [docs-godoc-url]: https://pkg.go.dev/github.com/go-xrp/rippled-openapi
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/go-xrp/rippled-openapi/blob/master/LICENSE