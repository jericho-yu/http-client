package httpClient

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	host         string
	url          string
	method       string
	body         []byte
	headers      map[string][]string
	request      *http.Request
	response     *http.Response
	Err          error
	responseBody []byte
}

// NewHttpClient 实例化：http客户端
func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

// SetHost 设置请求根地址
func (receiver *HttpClient) SetHost(host string) *HttpClient {
	receiver.host = host
	return receiver
}

// SetUrl 设置请求地址
func (receiver *HttpClient) SetUrl(url string) *HttpClient {
	receiver.url = url
	return receiver
}

// SetMethod 设置请求方法
func (receiver *HttpClient) SetMethod(method string) *HttpClient {
	receiver.method = method
	return receiver
}

// SetHeaders 设置请求头
func (receiver *HttpClient) SetHeaders(headers map[string][]string) *HttpClient {
	receiver.headers = headers
	return receiver
}

// SetBody 设置请求体
func (receiver *HttpClient) SetBody(body []byte) *HttpClient {
	receiver.body = body
	return receiver
}

// GetResponse 获取响应
func (receiver *HttpClient) GetResponse() *http.Response {
	return receiver.response
}

// 获取
func (receiver *HttpClient) GetResponseBody() []byte {
	return receiver.responseBody
}

// Send 发送请求
func (receiver *HttpClient) Send() *HttpClient {
	// 创建请求对象
	if receiver.Err = receiver.generateRequest(); receiver.Err != nil {
		return receiver
	}

	// 填充请求头
	receiver.setHeaders()

	// 检查请求对象
	if receiver.Err = receiver.check(); receiver.Err != nil {
		return receiver
	}

	// 发送新的请求
	client := &http.Client{}
	receiver.response, receiver.Err = client.Do(receiver.request)
	if receiver.Err != nil {
		receiver.Err = fmt.Errorf("发送失败：%s", receiver.Err.Error())
		return receiver
	}
	defer receiver.response.Body.Close()

	// 读取新的响应的主体
	receiver.responseBody, receiver.Err = ioutil.ReadAll(receiver.response.Body)
	if receiver.Err != nil {
		receiver.Err = fmt.Errorf("读取响应体失败：%s", receiver.Err.Error())
		return receiver
	}

	return receiver
}

// 检查条件是否满足
func (receiver *HttpClient) check() error {
	if receiver.host == "" {
		return errors.New("host不能为空")
	}
	if receiver.url == "" {
		return errors.New("url不能为空")
	}
	if receiver.method == "" {
		return errors.New("method不能为空")
	}
	return nil
}

// 生成请求对象
func (receiver *HttpClient) generateRequest() (e error) {
	receiver.request, e = http.NewRequest(receiver.method, receiver.host+receiver.url, bytes.NewReader(receiver.body))
	return
}

// 设置请求头
func (receiver *HttpClient) setHeaders() {
	for k, v := range receiver.headers {
		receiver.request.Header[k] = v
	}
}
