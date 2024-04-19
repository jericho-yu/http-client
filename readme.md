# 一个GO语言 http客户端 

# (HTTP Client Library in Go)

这是一个用Go语言编写的HTTP客户端库。它提供了一种简单的方式来发送HTTP请求，并处理返回的响应。

This is an HTTP client library written in Go. It provides a simple way to send HTTP requests and handle the returned responses.

## 用途 (Usage)

这个库可以用于发送HTTP请求，包括GET、POST、PUT、DELETE等方法。它还允许你设置请求头和请求体，并获取响应。

This library can be used to send HTTP requests, including GET, POST, PUT, DELETE, etc. It also allows you to set request headers and request bodies, and get responses.

## 使用方法 (Example Usage)

以下是一个简单的使用示例：

Here is a simple example of usage:

```go
package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/jericho-yu/http-client/httpClient"
)

func printResp(hc *httpClient.HttpClient) {
	fmt.Printf("状态：%s；状态码：%d；响应内容：%s\n", hc.GetResponse().Status, hc.GetResponse().StatusCode, hc.GetResponseRawBody())
	fmt.Printf("Status: %s; Status Code: %d; Response Body: %s\n", hc.GetResponse().Status, hc.GetResponse().StatusCode, hc.GetResponseRawBody())
}

func main() {
	// get 请求 这里还可以使用New("").SetMethod(http.MethodGet)
	// GET Request (You can also use New("").SetMethod(http.MethodGet))
	if hc1 := httpClient.NewGet("http://www.baidu.com").Send(); hc1.Err != nil {
		fmt.Println(hc1.Err.Error())
		return
	} else {
		printResp(hc1)
	}

	// post 请求 这里还可以使用New("").SetMethod(http.MethodPost)
	// POST Request (You can also use New("").SetMethod(http.MethodPost))
	body2 := map[string]any{"name": "jericho-yu"}
	hc2 := httpClient.
		NewPost("http://www.baidu.com").
		SetJsonBody(body2).
		Send()
	if hc2.Err != nil {
		fmt.Println(hc2.Err.Error())
		return
	}
	printResp(hc2)

	// 将请求响应保存到文件
	// Save the request response to a file
	hc3 := httpClient.
		NewPost("http://www.baidu.com").
		SetHeaderAccept(httpClient.AcceptSteam).
		Send().
		SaveResponseSteamFile("./baidu.html")
	if hc3.Err != nil {
		fmt.Println(hc3.Err.Error())
		return
	}
	printResp(hc3)

	// 将请求响应体解析为json（当然baidu的解析结果会报错：invalid character '<' looking for beginning of value）
	// Parse the request response body as json (of course, Baidu's parsing result will report an error: invalid character '<' looking for beginning of value)
	responseJson4 := make(map[string]any)
	hc4 := httpClient.NewPost("http://www.baidu.com").
		SetHeaderAccept(httpClient.AcceptJson).
		Send().
		GetResponseJsonBody(&responseJson4)
	if hc4.Err != nil {
		fmt.Println(hc4.Err.Error())
		return
	}
	printResp(hc4)

	// 还可以作为网关工具，转发本次请求到其他微服务
	// It can also be used as a gateway tool to forward the current request to other microservices
	c := *gin.Context
	// 转发给目标服务
	// Forward to the target service
	hc5 := httpClient.New("http://you-target.com").
		AddHeaders(map[string][]string{
			"Content-Type": c.Request.Header["Content-Type"],
			"Accept":       c.Request.Header["Accept"],
		}).
		SetMethod(c.Request.Method).
		SetBody(func() []byte {
			b, e := io.ReadAll(c.Request.Body)
			if e != nil {
				panic(e)
			}
			return b
		}()).
		Send()
	if hc5.Err != nil {
		panic(hc5.Err)
	}

	// 原样响应给前端
	// Respond to the front-end in the same way
	c.Data(hc5.GetResponse().StatusCode, hc5.GetResponse().Header.Get("Content-Type"), hc5.GetResponseRawBody())

	// 带有SSL证书的请求
    // Request with SSL certificate
	hc6 := httpClient.NewPost("http://you-target.com").
		SetCert("/path/to/your/cert.pem").
		Send()
	if hc6.Err != nil {
		panic(hc6.Err)
	}
	printResp(hc6)

    // 批量发送请求
	// Send requests in batches
	hc7 := httpClient.HttpClientMultiple{}.
		New().
		SetClients([]*httpClient.HttpClient{
			httpClient.NewGet("http://www.baidu.com"),
			httpClient.NewGet("http://www.google.com"),
		}).
		Send()
	for _, hc := range hc7.GetClients() {
		if hc.Err != nil {
			panic(hc.Err.Error())
		}
	}
}
```
**注意：** 如果没有使用NewGet（类似：NewGet,NewPost等）方法或者没有设置SetMethod()，那么`method`默认：GET。
**Note:** If you do not use the NewGet (similar to: NewGet, NewPost, etc.) method or do not set SetMethod(), then the method defaults to GET.