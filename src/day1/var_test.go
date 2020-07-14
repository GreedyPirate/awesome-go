package day1

import (
	"fmt"
	"testing"
)

func TestBasicGrammar(t *testing.T) {
	// 变量定义
	a := 0
	fmt.Println(a)

	str := "hello"
	fmt.Println(str)
	
	ch := 'a'
	fmt.Println(ch)

	// 批量初始化1
	var (
		i int
		j string
		k bool
		l [10]int

	)
	l[0] = 1
	fmt.Println("batch:",i,j,k,l)

	// 批量初始化2
	t1, t2 := "var1", 10
	fmt.Println(t1,t2)

	// 交换2个值
	v1, v2 := 10, 20
	v1, v2 = v2, v1
	fmt.Println("swap", v1, v2)

    const TIME_OUT = 3000
}

