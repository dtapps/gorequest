package gorequest

import "testing"

func TestGetRandomUserAgentSystem(t *testing.T) {
	t.Log(GetRandomUserAgentSystem())
}

func TestGetRandomUserAgent(t *testing.T) {
	t.Log(GetRandomUserAgent())
}

func TestDtaUa(t *testing.T) {
	t.Log(DtaUa())
}
