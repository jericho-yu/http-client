# HTTP Client Library in Go

这是一个用Go语言编写的HTTP客户端库。它提供了一种简单的方式来发送HTTP请求，并处理返回的响应。

## 用途

这个库可以用于发送HTTP请求，包括GET、POST、PUT、DELETE等方法。它还允许你设置请求头和请求体，并获取响应。

## 使用方法

以下是一个简单的使用示例：

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jericho-yu/http-client/httpClient"
)

func printResp(hc *httpClient.HttpClient) {
	fmt.Printf("状态：%s；状态码：%d；响应内容：%s\n", hc.GetResponse().Status, hc.GetResponse().StatusCode, hc.GetResponseBody())
}

func main() {
    // get 请求 这里还可以使用New("").SetMethod(http.MethodGet)
	hc1 := httpClient.NewGet("http://www.baidu.com").
		Send()
	if hc1.Err != nil {
		fmt.Println(hc1.Err.Error())
		return
	}
	printResp(hc1)

	// post 请求 这里还可以使用New("").SetMethod(http.MethodPost)
	body2 := map[string]any{"name": "jericho-yu"}
	hc2 := httpClient.NewPost("http://www.baidu.com").
		SetJsonBody(body2).
		Send()
	if hc2.Err != nil {
		fmt.Println(hc2.Err.Error())
		return
	}
	printResp(hc2)
}

```
**注意：** `method`默认：GET。