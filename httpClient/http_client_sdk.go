package httpClient

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
)

type (
	HttpClientSdk struct {
		host    string
		url     string
		queries map[string]string
		method  string
		hc      *HttpClient
	}
)

func Get(host, url string, queries map[string]string) *HttpClientSdk {
	hcs := &HttpClientSdk{host: host, url: url, queries: queries, method: http.MethodGet}
	hcs.generateHttpClient()
	return hcs
}

func Post(host, url string, queries map[string]string) *HttpClientSdk {
	hcs := &HttpClientSdk{host: host, url: url, queries: queries, method: http.MethodPost}
	hcs.generateHttpClient()
	return hcs
}

func Put(host, url string, queries map[string]string) *HttpClientSdk {
	hcs := &HttpClientSdk{host: host, url: url, queries: queries, method: http.MethodPut}
	hcs.generateHttpClient()
	return hcs
}

func Delete(host, url string, queries map[string]string) *HttpClientSdk {
	hcs := &HttpClientSdk{host: host, url: url, queries: queries, method: http.MethodDelete}
	hcs.generateHttpClient()
	return hcs
}

// 实例化http client对象
func (receiver *HttpClientSdk) generateHttpClient() {
	receiver.hc = New().SetUrl(receiver.url).SetQueries(receiver.queries).SetMethod(receiver.method)
}

// Json 设置json格式
func (receiver *HttpClientSdk) Json(body any) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("json")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("json")
	}

	b, e := json.Marshal(body)
	if e != nil {
		receiver.hc.Err = e
		return receiver.hc
	}

	return receiver.hc.SetBody(b).GenerateRequest().Send()
}

// Xml 设置xml格式
func (receiver *HttpClientSdk) Xml(body []byte) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("xml")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("xml")
	}

	b, e := xml.Marshal(body)
	if e != nil {
		receiver.hc.Err = e
		return receiver.hc
	}

	return receiver.hc.SetBody(b).GenerateRequest().Send()
}

// Form 设置form格式
func (receiver *HttpClientSdk) Form(body map[string]string) *HttpClient {
	switch receiver.method {
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("form")
	}

	params := url.Values{}
	if len(body) > 0 {
		for k, v := range body {
			params.Add(k, v)
		}
	}

	return receiver.hc.SetBody([]byte(params.Encode())).GenerateRequest().Send()
}

// Plain 设置plain格式
func (receiver *HttpClientSdk) Plain(body string) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("plain")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("plain")
	}

	return receiver.hc.SetBody([]byte(body)).GenerateRequest().Send()
}

// Html 设置html格式
func (receiver *HttpClientSdk) Html(body string) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("html")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("html")
	}

	return receiver.hc.SetBody([]byte(body)).GenerateRequest().Send()
}

// Css 设置css格式
func (receiver *HttpClientSdk) Css(body string) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("css")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("css")
	}

	return receiver.hc.SetBody([]byte(body)).GenerateRequest().Send()
}

// Javascript 设置javascript格式
func (receiver *HttpClientSdk) Javascript(body string) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("javascript")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("javascript")
	}

	return receiver.hc.SetBody([]byte(body)).GenerateRequest().Send()
}

// Steam 设置steam格式
func (receiver *HttpClientSdk) Steam(body []byte) *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("steam")
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		receiver.hc.SetHeaderContentType("steam")
	}

	return receiver.hc.SetBody(body).GenerateRequest().Send()
}

// Any 设置任意格式
func (receiver *HttpClientSdk) Any() *HttpClient {
	switch receiver.method {
	case http.MethodGet:
		receiver.hc.SetHeaderAccept("any")
	}

	return receiver.hc.GenerateRequest().Send()
}
