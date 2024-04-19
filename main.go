package main

import (
	"fmt"
	"net/http"

	"github.com/jericho-yu/http-client/httpClient"
)

func printResp(hc *httpClient.HttpClient) {
	fmt.Printf("状态：%s；状态码：%d；响应内容：%s\n", hc.GetResponse().Status, hc.GetResponse().StatusCode, hc.GetResponseBody())
}

func main() {
	// get 请求
	hc1 := httpClient.New("http://www.baidu.com").
		Send()
	if hc1.Err != nil {
		fmt.Println(hc1.Err.Error())
		return
	}
	printResp(hc1)

	// post 请求
	body2 := map[string]any{"name": "jericho-yu"}
	hc2 := httpClient.New("http://www.baidu.com").
		SetMethod(http.MethodPost).
		SetJsonBody(body2).
		Send()
	if hc2.Err != nil {
		fmt.Println(hc2.Err.Error())
		return
	}
	printResp(hc2)
}
