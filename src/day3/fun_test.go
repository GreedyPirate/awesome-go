package day3

import "testing"

func TestFunction(t *testing.T) {

}

// 类型一致的形参可以这样写
func typeful(a, b int, s1, s2 string)  {
	print(a, b, s1, s2)
}

func multi() (int, int)  {
	return 1,2
}
