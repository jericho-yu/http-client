package httpClient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type (
	// HttpClient http客户端
	HttpClient struct {
		Err            error
		requestUrl     string
		requestQueries map[string]string
		requestMethod  string
		requestBody    []byte
		requestHeaders map[string][]string
		request        *http.Request
		response       *http.Response
		responseBody   []byte
		isReady        bool
	}
)

// New 实例化：http客户端
func New(url string) *HttpClient {
	return &HttpClient{
		requestUrl:     url,
		requestQueries: map[string]string{},
		requestHeaders: map[string][]string{},
		responseBody:   []byte{},
	}
}

// NewGet 实例化：http客户端get请求
func NewGet(url string) *HttpClient {
	return New(url).SetMethod(http.MethodGet)
}

// NewPost 实例化：http客户端post请求
func NewPost(url string) *HttpClient {
	return New(url).SetMethod(http.MethodPost)
}

// NewPut 实例化：http客户端put请求
func NewPut(url string) *HttpClient {
	return New(url).SetMethod(http.MethodPut)
}

// NewDelete 实例化：http客户端delete请求
func NewDelete(url string) *HttpClient {
	return New(url).SetMethod(http.MethodDelete)
}

// SetUrl 设置请求地址
func (receiver *HttpClient) SetUrl(url string) *HttpClient {
	receiver.requestUrl = url
	return receiver
}

// SetMethod 设置请求方法
func (receiver *HttpClient) SetMethod(method string) *HttpClient {
	receiver.requestMethod = method
	return receiver
}

// AddHeaders 设置请求头
func (receiver *HttpClient) AddHeaders(headers map[string][]string) *HttpClient {
	receiver.requestHeaders = headers
	return receiver
}

// SetQueries 设置请求参数
func (receiver *HttpClient) SetQueries(queries map[string]string) *HttpClient {
	receiver.requestQueries = queries
	return receiver
}

// SetBody 设置请求体
func (receiver *HttpClient) SetBody(body []byte) *HttpClient {
	receiver.requestBody = body
	return receiver
}

// SetJsonBody 设置json请求体
func (receiver *HttpClient) SetJsonBody(body any) *HttpClient {
	receiver.SetHeaderContentType("json")

	receiver.requestBody, receiver.Err = json.Marshal(body)
	return receiver
}

// SetXmlBody 设置xml请求体
func (receiver *HttpClient) SetXmlBody(body any) *HttpClient {
	receiver.SetHeaderContentType("xml")

	receiver.requestBody, receiver.Err = xml.Marshal(body)
	return receiver
}

// SetFormBody 设置表单请求体
func (receiver *HttpClient) SetFormBody(body map[string]string) *HttpClient {
	receiver.SetHeaderContentType("form")

	params := url.Values{}
	for k, v := range body {
		params.Add(k, v)
	}
	receiver.requestBody = []byte(params.Encode())
	return receiver
}

// SetFormDataBody 设置表单数据请求体
func (receiver *HttpClient) SetFormDataBody(texts map[string]string, files map[string]string) *HttpClient {
	var (
		e      error
		buffer bytes.Buffer
	)

	receiver.SetHeaderContentType("form-data")

	writer := multipart.NewWriter(&buffer)

	if len(texts) > 0 {
		for k, v := range texts {
			e = writer.WriteField(k, v)
			if e != nil {
				receiver.Err = e
				return receiver
			}
		}
	}

	if len(files) > 0 {
		for k, v := range files {
			fileWriter, _ := writer.CreateFormFile("fileField", k)
			file, _ := os.Open(v)
			_, e = io.Copy(fileWriter, file)
			if e != nil {
				receiver.Err = e
				return receiver
			}
			defer func(file *os.File) {
				e = file.Close()
				if e != nil {
					panic(e)
				}
			}(file)
		}
	}

	receiver.requestBody = []byte(writer.FormDataContentType())

	return receiver
}

// SetPlainBody 设置纯文本请求体
func (receiver *HttpClient) SetPlainBody(text string) *HttpClient {
	receiver.SetHeaderContentType("plain")

	receiver.requestBody = []byte(text)

	return receiver
}

// SetHtmlBody 设置html请求体
func (receiver *HttpClient) SetHtmlBody(text string) *HttpClient {
	receiver.SetHeaderContentType("html")

	receiver.requestBody = []byte(text)

	return receiver
}

// SetCssBody 设置Css请求体
func (receiver *HttpClient) SetCssBody(text string) *HttpClient {
	receiver.SetHeaderContentType("css")

	receiver.requestBody = []byte(text)

	return receiver
}

// SetJavascriptBody 设置Javascript请求体
func (receiver *HttpClient) SetJavascriptBody(text string) *HttpClient {
	receiver.SetHeaderContentType("javascript")

	receiver.requestBody = []byte(text)

	return receiver
}

func (receiver *HttpClient) SetSteamBody(file string) *HttpClient {
	receiver.SetHeaderContentType("steam")

	fileData, e := os.ReadFile(file)
	if e != nil {
		receiver.Err = e
		return receiver
	}
	receiver.requestBody = fileData

	return receiver
}

// SetHeaderContentType 设置请求头内容类型
func (receiver *HttpClient) SetHeaderContentType(key string) *HttpClient {
	value := ContentType{}.GetValue(key)
	if value != "" {
		receiver.requestHeaders["Content-Type"] = []string{value}
	}

	return receiver
}

// SetHeaderAccept 设置请求头接受内容类型
func (receiver *HttpClient) SetHeaderAccept(key string) *HttpClient {
	value := Accept{}.GetValue(key)
	if value != "" {
		receiver.requestHeaders["Accept"] = []string{value}
	}

	return receiver
}

// GetResponse 获取响应对象
func (receiver *HttpClient) GetResponse() *http.Response {
	return receiver.response
}

// GetResponseRawBody 获取原始响应体
func (receiver *HttpClient) GetResponseRawBody() []byte {
	return receiver.responseBody
}

// GetResponseJsonBody 获取json格式响应体
func (receiver *HttpClient) GetResponseJsonBody(target any) *HttpClient {
	if e := json.Unmarshal(receiver.responseBody, &target); e != nil {
		receiver.Err = e
	}
	return receiver
}

// GetResponseXmlBody 获取xml格式响应体
func (receiver *HttpClient) GetResponseXmlBody(target any) *HttpClient {
	if e := xml.Unmarshal(receiver.responseBody, &target); e != nil {
		receiver.Err = e
	}
	return receiver
}

// SaveResponseSteamFile 保存二进制到文件
func (receiver *HttpClient) SaveResponseSteamFile(filename string) *HttpClient {
	// 创建一个新的文件
	file, err := os.Create(filename)
	if err != nil {
		receiver.Err = err
		return receiver
	}
	defer func() { file.Close() }()

	// 将二进制数据写入文件
	_, err = file.Write(receiver.responseBody)
	if err != nil {
		receiver.Err = err
		return receiver
	}

	return receiver
}

// GetRequest 获取请求
func (receiver *HttpClient) GetRequest() *http.Request {
	return receiver.request
}

// GenerateRequest 生成请求对象
func (receiver *HttpClient) GenerateRequest() *HttpClient {
	var e error

	receiver.request, e = http.NewRequest(receiver.requestMethod, receiver.requestUrl, bytes.NewReader(receiver.requestBody))
	if e != nil {
		receiver.Err = fmt.Errorf("生成请求对象失败：%s", e.Error())
		return receiver
	}

	// 设置请求头
	receiver.addHeaders()

	// 设置url参数
	receiver.setQueries()

	// 检查请求对象
	if receiver.Err = receiver.check(); receiver.Err != nil {
		return receiver
	}

	receiver.isReady = true

	return receiver
}

// Send 发送请求
func (receiver *HttpClient) Send() *HttpClient {
	if !receiver.isReady {
		receiver.GenerateRequest()
		if receiver.Err != nil {
			return receiver
		}
	}

	// 发送新的请求
	client := &http.Client{}
	receiver.response, receiver.Err = client.Do(receiver.request)
	if receiver.Err != nil {
		receiver.Err = fmt.Errorf("发送失败：%s", receiver.Err.Error())
		return receiver
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(receiver.response.Body)

	// 读取新的响应的主体
	receiver.responseBody, receiver.Err = io.ReadAll(receiver.response.Body)
	if receiver.Err != nil {
		receiver.Err = fmt.Errorf("读取响应体失败：%s", receiver.Err.Error())
		return receiver
	}

	receiver.isReady = false

	return receiver
}

// 检查条件是否满足
func (receiver *HttpClient) check() error {
	if receiver.requestUrl == "" {
		return errors.New("url不能为空")
	}
	if receiver.requestMethod == "" {
		receiver.requestMethod = http.MethodGet
	}
	return nil
}

// 设置url参数
func (receiver *HttpClient) setQueries() {
	if len(receiver.requestQueries) > 0 {
		queries := url.Values{}
		for k, v := range receiver.requestQueries {
			queries.Add(k, v)
		}
		receiver.requestUrl += "?" + queries.Encode()
	}
}

// 设置请求头
func (receiver *HttpClient) addHeaders() {
	for k, v := range receiver.requestHeaders {
		receiver.request.Header[k] = append(receiver.request.Header[k], v...)
	}
}
