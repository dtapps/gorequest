package gorequest

import (
	"log"
	"testing"
)

func TestParams(t *testing.T) {
	params1 := NewParams()
	params2 := NewParams()
	params1.Set("a", "1")
	params2.Set("b", "2")
	params3 := NewParamsWith(params1, params2)
	log.Println(params1.DeepCopy())
	log.Println(params2.DeepCopy())
	log.Println(params3.DeepCopy())
	log.Println(params1.DeepCopy())
	log.Println(params2.DeepCopy())
}
