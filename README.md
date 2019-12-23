# ShopeeGo [![GoDoc](https://godoc.org/github.com/teacat/shopeego?status.svg)](https://godoc.org/github.com/teacat/shopeego) [![Coverage Status](https://coveralls.io/repos/github/teacat/shopeego/badge.svg?branch=master)](https://coveralls.io/github/teacat/shopeego?branch=master) [![Build Status](https://travis-ci.org/teacat/shopeego.svg?branch=master)](https://travis-ci.org/teacat/shopeego) [![Go Report Card](https://goreportcard.com/badge/github.com/teacat/rushia)](https://goreportcard.com/report/github.com/teacat/shopeego)

由 [Golang](https://golang.org/) 所實作的 Shopee Open Platform API，包含所有 Shopee API v1 所提供的功能，這些功能都已經列舉在[本專案的 Go Doc](https://godoc.org/github.com/teacat/shopeego) 之中。使用方式詳見 [Shopee Open Platform Documentation](https://open.shopee.com/documents)。

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get gopkg.in/teacat/shopeego.v1
```

# 開始使用

```golang
client := NewClient(&ClientOptions{
	Secret: "0c2c7b3bd59c2503f49a307fcf49dc985df43a1214821d5453e9ba54ca8e2e44",
})
client.GetShopInfo(&GetShopInfoRequest{
	PartnerID: 841237,
	ShopID:    307271,
	Timestamp: int(time.Now().Unix()),
})
```

# 驗證介面

# 簽署