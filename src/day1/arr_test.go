package day1

import "testing"

/**
数组长度不可变
 */
func TestArray(t *testing.T) {
	// 初始化数组同时赋值
	var a1 = [2]int{1,2}

	// 仅初始化数组长度
	var a2 = [2]int{}
	// 按index赋值
	a2[0] = 1

	// 1,2表示索引
	a3 := [5]int{1:10, 2:20}

	t.Log(a1, a2, a3)

	// 长度可省略，go会自动计算
	var arr1 = [...]int{1,2,3,4}
	var arr2 = [...]int{1,2,3,5}

	// 数组长度相同才能比较
	t.Log("arr1 == arr2 ? ", arr1 == arr2)

	for i := 0; i < len(arr1); i++ {
		ele := arr1[i]
		t.Log(ele)
	}

	t.Log("for each")
	for _,item := range arr1{
		t.Log(item)
	}
}
