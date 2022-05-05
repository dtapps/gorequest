package gorequest

import "testing"

func TestUrlParse(t *testing.T) {
	t.Logf("%+v", UriParse("https://api.dtapp.net/ip?ip=127.0.0.1#history"))
	t.Logf("%+v", UriParse("https://test:abcd123@golangbyexample.com:8000/tutorials/intro?type=advance&compact=false#history"))
	t.Logf("%+v", UriFilterExcludeQueryString("/"))
	t.Logf("%+v", UriFilterExcludeQueryString("/favicon.ico"))
	t.Logf("%+v", UriFilterExcludeQueryString("/ip"))
	t.Logf("%+v", UriFilterExcludeQueryString("/ip?ip=127.0.0.1"))
}
