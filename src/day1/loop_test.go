package day1

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	for i :=0; i<10; i++ {
		fmt.Print(i)
	}

	a := 0
	for a < 5 {
		fmt.Println(a)
		a++;
	}

	if v := 1 < 2; v {
		fmt.Sprint(v)
	}
}
