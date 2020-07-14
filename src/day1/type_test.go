package day1

import (
	"fmt"
	"testing"
)

/**
类型定义和类型别名
 */

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func TestType(t *testing.T) {
	var a NewInt
	fmt.Printf("a type %T \n", a)

	var b IntAlias
	fmt.Printf("b type %T \n", b)
}