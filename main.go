package main

import (
	"fmt"

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

	// // 还可以作为网关工具，转发本次请求到其他微服务
	// // It can also be used as a gateway tool to forward the current request to other microservices
	// c := *gin.Context
	// // 转发给目标服务
	// // Forward to the target service
	// hc5 := httpClient.New("http://you-target.com").
	// 	AddHeaders(map[string][]string{
	// 		"Content-Type": c.Request.Header["Content-Type"],
	// 		"Accept":       c.Request.Header["Accept"],
	// 	}).
	// 	SetMethod(c.Request.Method).
	// 	SetBody(func() []byte {
	// 		b, e := io.ReadAll(c.Request.Body)
	// 		if e != nil {
	// 			panic(e)
	// 		}
	// 		return b
	// 	}()).
	// 	Send()
	// if hc5.Err != nil {
	// 	panic(hc5.Err)
	// }

	// // 原样响应给前端
	// // Respond to the front-end in the same way
	// c.Data(hc5.GetResponse().StatusCode, hc5.GetResponse().Header.Get("Content-Type"), hc5.GetResponseRawBody())

	// 带有SSL证书的请求
	// Request with SSL certificate
	hc6 := httpClient.NewPost("http://you-target.com").
		SetCert("/path/to/your/cert.pem").
		Send()
	if hc6.Err != nil {
		panic(hc6.Err)
	}
	printResp(hc6)
}
