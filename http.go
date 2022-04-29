package gorequest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Response 返回内容
type Response struct {
	RequestUrl            string      //【请求】链接
	RequestParams         Params      //【请求】参数
	RequestMethod         string      //【请求】方式
	RequestHeader         Headers     //【请求】头部
	ResponseHeader        http.Header //【返回】头部
	ResponseStatus        string      //【返回】状态
	ResponseStatusCode    int         //【返回】状态码
	ResponseBody          []byte      //【返回】内容
	ResponseContentLength int64       //【返回】大小
}

type App struct {
	Url             string   // 全局请求地址，没有设置url才会使用
	httpUrl         string   // 请求地址
	httpMethod      string   // 请求方法
	httpHeader      Headers  // 请求头
	httpParams      Params   // 请求参数
	responseContent Response // 返回内容
	httpContentType string   // 请求内容类型
	Error           error    // 错误
}

var (
	httpParamsModeJson = "JSON"
	httpParamsModeForm = "FORM"
)

// NewHttp 实例化
func NewHttp() *App {
	return &App{
		httpHeader: NewHeaders(),
		httpParams: NewParams(),
	}
}

// SetUrl 设置请求地址
func (app *App) SetUrl(url string) {
	app.httpUrl = url
}

// SetMethod 设置请求方式地址
func (app *App) SetMethod(method string) {
	app.httpMethod = method
}

// SetHeader 设置请求头
func (app *App) SetHeader(key, value string) {
	if key == "" {
		panic("url is empty")
	}
	app.httpHeader.Set(key, value)
}

// SetHeaders 批量设置请求头
func (app *App) SetHeaders(headers Headers) {
	for key, value := range headers {
		app.httpHeader.Set(key, value)
	}
}

// SetAuthToken 设置身份验证令牌
func (app *App) SetAuthToken(token string) {
	app.httpHeader.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

// SetUserAgent 设置用户代理，空字符串就随机设置
func (app *App) SetUserAgent(ua string) {
	if ua == "" {
		ua = GetRandomUserAgent()
	}
	app.httpHeader.Set("User-Agent", ua)
}

// SetContentTypeJson 设置JSON格式
func (app *App) SetContentTypeJson() {
	app.httpContentType = httpParamsModeJson
}

// SetContentTypeForm 设置FORM格式
func (app *App) SetContentTypeForm() {
	app.httpContentType = httpParamsModeForm
}

// SetParam 设置请求参数
func (app *App) SetParam(key string, value interface{}) {
	if key == "" {
		panic("url is empty")
	}
	app.httpParams.Set(key, value)
}

// SetParams 批量设置请求参数
func (app *App) SetParams(params Params) {
	for key, value := range params {
		app.httpParams.Set(key, value)
	}
}

// Get 发起GET请求
func (app *App) Get() (httpResponse Response, err error) {
	// 设置请求方法
	app.httpMethod = http.MethodGet
	return request(app)
}

// Post 发起POST请求
func (app *App) Post() (httpResponse Response, err error) {
	// 设置请求方法
	app.httpMethod = http.MethodPost
	return request(app)
}

// Request 发起请求
func (app *App) Request() (httpResponse Response, err error) {
	return request(app)
}

// 请求
func request(app *App) (httpResponse Response, err error) {

	// 判断网址
	if app.httpUrl == "" {
		app.httpUrl = app.Url
	}
	if app.httpUrl == "" {
		return httpResponse, errors.New("没有设置Url")
	}

	// 创建 http 客户端
	client := &http.Client{}

	// 赋值
	httpResponse.RequestUrl = app.httpUrl
	httpResponse.RequestMethod = app.httpMethod
	httpResponse.RequestParams = app.httpParams

	var reqBody io.Reader

	if app.httpMethod == http.MethodPost && app.httpContentType == httpParamsModeJson {
		app.httpHeader.Set("Content-Type", "application/json")
		jsonStr, err := json.Marshal(app.httpParams)
		if err != nil {
			return httpResponse, errors.New(fmt.Sprintf("解析出错 %s", err))
		}
		// 赋值
		reqBody = bytes.NewBuffer(jsonStr)
	}

	if app.httpMethod == http.MethodPost && app.httpContentType == httpParamsModeForm {
		// 携带 form 参数
		form := url.Values{}
		app.httpHeader.Set("Content-Type", "application/x-www-form-urlencoded")
		if len(app.httpParams) > 0 {
			for k, v := range app.httpParams {
				form.Add(k, GetParamsString(v))
			}
		}
		// 赋值
		reqBody = strings.NewReader(form.Encode())
	}

	// 创建请求
	req, err := http.NewRequest(app.httpMethod, app.httpUrl, reqBody)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("创建请求出错 %s", err))
	}

	// GET 请求携带查询参数
	if app.httpMethod == http.MethodGet {
		if len(app.httpParams) > 0 {
			q := req.URL.Query()
			for k, v := range app.httpParams {
				q.Add(k, GetParamsString(v))
			}
			req.URL.RawQuery = q.Encode()
		}
	}

	// 设置请求头
	if len(app.httpHeader) > 0 {
		for key, value := range app.httpHeader {
			req.Header.Set(key, value)
		}
	}
	// 赋值
	httpResponse.RequestHeader = app.httpHeader

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}

	// 最后关闭连接
	defer resp.Body.Close()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}

	// 赋值
	httpResponse.ResponseStatus = resp.Status
	httpResponse.ResponseStatusCode = resp.StatusCode
	httpResponse.ResponseHeader = resp.Header
	httpResponse.ResponseBody = body
	httpResponse.ResponseContentLength = resp.ContentLength

	return httpResponse, err
}
