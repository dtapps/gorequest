package gorequest

import (
	"net/url"
	"strings"
)

// ResponseUrlParse 返回参数
type ResponseUrlParse struct {
	Uri      string `json:"uri"`       // URI
	Urn      string `json:"urn"`       // URN
	Url      string `json:"url"`       // URL
	Scheme   string `json:"scheme"`    // 协议
	Host     string `json:"host"`      // 主机
	Hostname string `json:"hostname"`  // 主机名
	Port     string `json:"port"`      // 端口
	Path     string `json:"path"`      // 路径
	RawQuery string `json:"raw_query"` // 参数 ?
	Fragment string `json:"fragment"`  // 片段 #
}

// UriParse 解析URl
func UriParse(input string) (resp ResponseUrlParse) {
	parse, err := url.Parse(input)
	if err != nil {
		return
	}
	resp.Uri = input
	resp.Urn = parse.Host + parse.Path
	resp.Url = parse.Scheme + "://" + parse.Host + parse.Path
	resp.Scheme = parse.Scheme
	resp.Host = parse.Host
	resp.Hostname = parse.Hostname()
	resp.Port = parse.Port()
	resp.Path = parse.Path
	resp.RawQuery = parse.RawQuery
	resp.Fragment = parse.Fragment
	return
}

// UriFilterExcludeQueryString 过滤掉url中的参数
func UriFilterExcludeQueryString(uri string) string {
	URL, _ := url.Parse(uri)
	clearUri := strings.ReplaceAll(uri, URL.RawQuery, "")
	clearUri = strings.TrimRight(clearUri, "?")
	return strings.TrimRight(clearUri, "/")
}
