package main

import (
	"fmt"

	"github.com/jericho-yu/http-client/httpClient"
)

func printResp(hc *httpClient.HttpClient) {
	fmt.Printf("状态：%s；状态码：%d；响应内容：%s\n", hc.GetResponse().Status, hc.GetResponse().StatusCode, hc.GetResponseRawBody())
}

func main() {
	// get 请求 这里还可以使用New("").SetMethod(http.MethodGet)
	if hc1 := httpClient.NewGet("http://www.baidu.com").Send(); hc1.Err != nil {
		fmt.Println(hc1.Err.Error())
		return
	} else {
		printResp(hc1)
	}

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

	// 将请求响应保存到文件
	hc3 := httpClient.NewPost("http://www.baidu.com").
		SetHeaderAccept(httpClient.AcceptSteam).
		Send().
		SaveResponseSteamFile("./baidu.html")
	if hc3.Err != nil {
		fmt.Println(hc3.Err.Error())
		return
	}
	printResp(hc3)

	// 将请求响应体解析为json（当然baidu的解析结果会报错：invalid character '<' looking for beginning of value）
	responseJson4 := make(map[string]any)
	hc4 := httpClient.NewPost("http://www.baidu.com").SetHeaderAccept(httpClient.AcceptJson).Send().GetResponseJsonBody(&responseJson4)
	if hc4.Err != nil {
		fmt.Println(hc4.Err.Error())
		return
	}
	printResp(hc4)

	// // 还可以作为网关工具，转发本次请求到其他微服务
	// c := *gin.Context
	// requestBody5, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// // 转发给目标服务
	// hc5 := httpClient.New("http://you-target.com").
	// 	AddHeaders(map[string][]string{
	// 		"Content-Type": []string{c.Request.Header.Get("Content-Type")},
	// 		"Accept":       []string{c.Request.Header.Get("Accept")},
	// 	}).
	// 	SetMethod("这里写你接收请求的方法，以gin为例：*gin.Context.Request.Method").
	// 	SetBody(requestBody5)

	// // 原样响应给前端
	// c.String(hc5.GetResponse().StatusCode, string(hc5.GetResponseRawBody()))
}
