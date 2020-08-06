package day1

import (
	"fmt"
	"testing"
)

func TestGoto(t *testing.T) {
	for x := 0; x < 10; x++ {
		t.Log("x: ",x)
		for y := 0; y < 10; y++ {

			if y == 6 {
				// 跳转到标签
				goto breakHere
			}
		}
	}

	// 手动返回, 避免执行进入标签
	return

	// 标签
breakHere:
	fmt.Println("done")

OuterLoop:
	for i := 0; i < 2; i++ {

		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println("out:", i, j)
				continue OuterLoop
			}
			t.Log("all:", i,j)
		}
	}
}

/**
用goto抽取功能代码
 */
func service(price int)  {
	if(price < 0) {
		goto ReturnError
	}
	if(price > 100) {
		goto ReturnError
	}

ReturnError:
	fmt.Println("bad param")
}
