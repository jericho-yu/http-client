# HTTP Client Library in Go

这是一个用Go语言编写的HTTP客户端库。它提供了一种简单的方式来发送HTTP请求，并处理返回的响应。

## 用途

这个库可以用于发送HTTP请求，包括GET、POST、PUT、DELETE等方法。它还允许你设置请求头和请求体，并获取响应。

## 使用方法

以下是一个简单的使用示例：

```go
package main

import (
	"fmt"
	"httpClient"
)

func main() {
	hc := httpClient.
		NewHttpClient().
		SetHost("http://").
		SetMethod(http.MethodGet).
		SetUrl("baidu.com").
		Send()

	if hc.Err != nil {
		println(hc.Err.Error())
	}
	println(string(hc.GetResponseBody()))
	println("status code:", hc.GetResponse().StatusCode)
	println(fmt.Sprintf("header %v", hc.GetResponse().Header))
}
```
**注意：** `host`、`method`、`url`是必须设置的，否则会返回错误。