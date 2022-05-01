package gorequest

import (
	"net/url"
)

type ResponseUrlParse struct {
	Scheme   string `json:"scheme"`    // 协议
	Host     string `json:"host"`      // 主机
	Hostname string `json:"hostname"`  // 主机名
	Port     string `json:"port"`      // 端口
	HostPath string `json:"host_path"` // 主机加路径
	Path     string `json:"path"`      // 路径
	RawQuery string `json:"raw_query"` // 参数 ?
	Fragment string `json:"fragment"`  // 片段 #
}

// UrlParse 解析URl
func UrlParse(inputUrl string) (resp ResponseUrlParse) {
	parse, err := url.Parse(inputUrl)
	if err != nil {
		return
	}
	resp.Scheme = parse.Scheme
	resp.Host = parse.Host
	resp.Hostname = parse.Hostname()
	resp.Port = parse.Port()
	resp.HostPath = parse.Host + parse.Path
	resp.Path = parse.Path
	resp.RawQuery = parse.RawQuery
	resp.Fragment = parse.Fragment
	return
}