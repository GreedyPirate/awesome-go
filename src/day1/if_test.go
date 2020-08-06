package day1

import "testing"

func TestIf(t *testing.T) {
	if 2 > 1{
		t.Log("right")
	} else{

	}

	// while(true)形式
	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	t.Log(sum)

	// while形式
	var i int
	for i <= 10 {
		i++
	}
}
