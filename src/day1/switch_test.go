package day1

import "testing"

func TestSwitch(t *testing.T) {

	i := 1

	switch i {
	case 1,2:
		t.Log("1 or 2")
	case 3:
		t.Log("error")
	default:
		t.Log("0")
	}

	switch  {
	case i < 2:
		t.Log(" < 2")
		fallthrough // 可以往下执行走case
	case i < 3:
		t.Log("  < 3")

	}
}
