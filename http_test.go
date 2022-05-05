package gorequest

import (
	"testing"
)

func TestGet(t *testing.T) {
	appHttp := NewHttp()
	appHttp.SetDebug()
	//appHttp.SetUri("https://api.dtapp.net/ip")
	get, err := appHttp.Get("https://api.dtapp.net/ip")

	t.Logf("get：%+v\n", get)
	t.Logf("get.ResponseBody：%s\n", get.ResponseBody)
	t.Logf("err：%s\n", err)
}

func TestPost(t *testing.T) {
	appHttp := NewHttp()
	appHttp.SetUri("https://api.dtapp.net/ip")
	get, err := appHttp.Post()

	t.Logf("get：%+v\n", get)
	t.Logf("get.ResponseBody：%s\n", get.ResponseBody)
	t.Logf("err：%s\n", err)
}
