package day1

import (
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {

	hp := 100
	ptr := &hp

	t.Log(reflect.TypeOf(ptr))
	t.Log(reflect.TypeOf(*ptr))
	t.Log(ptr, *ptr)

	// new创建指针
	p := new(string)
	*p = "text"
	str := *p

	t.Log(str)
}
