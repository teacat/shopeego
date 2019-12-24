# (WIP 尚未完成) ShopeeGo [![GoDoc](https://godoc.org/github.com/teacat/shopeego?status.svg)](https://godoc.org/github.com/teacat/shopeego) [![Coverage Status](https://coveralls.io/repos/github/teacat/shopeego/badge.svg?branch=master)](https://coveralls.io/github/teacat/shopeego?branch=master) [![Build Status](https://travis-ci.org/teacat/shopeego.svg?branch=master)](https://travis-ci.org/teacat/shopeego) [![Go Report Card](https://goreportcard.com/badge/github.com/teacat/rushia)](https://goreportcard.com/report/github.com/teacat/shopeego)

由 [Golang](https://golang.org/) 所實作的 Shopee Open Platform API，包含所有 Shopee API v1 所提供的功能，這些功能都已經列舉在[本專案的 Go Doc](https://godoc.org/github.com/teacat/shopeego) 之中。使用方式詳見 [Shopee Open Platform Documentation](https://open.shopee.com/documents)。

**注意：**Shopeego 是根據 Shopee 所提供的 API 文件與資料型態進行開發，但有些資料型態明顯是不正確的，且有些[「必填欄位」被標註為「可選」](https://github.com/minchao/shopee-php/issues/5)。基於這個狀況，當遇到 API 無法正常使用或是出現 `error_param` 時，請考慮到 Issue 中回報。

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get gopkg.in/teacat/shopeego.v1
```

# 開始使用

```golang
client := shopeego.NewClient(&ClientOptions{
	Secret: "0c2c7b3bd59c2503f49a307fcf49dc985df43a1214821d5453e9ba54ca8e2e44",
})
shop := client.GetShopInfo(&GetShopInfoRequest{
	PartnerID: 841237,
	ShopID:    307271,
	Timestamp: int(time.Now().Unix()),
})
fmt.Println(shop.ShopName) // 輸出：Yami Odymel 的早餐店！
```

# 驗證介面

欲連接的賣場必須授權給你的 [Shopee Open Platform 應用程式](https://open.shopee.com/documents?module=63&type=2&id=51) 才能夠與其對接。

# 簽署

應[蝦皮官方要求](https://open.shopee.com/documents?module=63&type=2&id=53)所有的請求都必須簽署 HMAC-SHA256，很明顯這是有點神奇的蹦蹦安全手段，但這個套件仍然有幫你作這件事。如果你不明白，這裡舉個範例。

```
POST /api/v1/orders/detail HTTP/1.1
Host: partner.shopeemobile.com
Authorization: b37c061daf2fcfa2baffe7539110938be5b7525041c147e78ad8afa78cc1a72d

{"ordersn": "160726152598865", "shopid": 61299, "partner_id": 1, "timestamp": 1470198856}
```

在請求標頭的 `Authorization` 欄位中，這個雜湊碼來自於「`請求網址|內容`」並透過你的 API 金鑰作為祕密（Secret）演算而成。