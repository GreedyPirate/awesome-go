package day2

import (
	"strconv"
	"testing"
)

/**
切片就是可以动态伸缩的数组 ？
 */
func TestSlice(t *testing.T) {
	// 不写长度就是切片？
	var s1 []int
	t.Log(len(s1), cap(s1))

	s1 = append(s1, 1)
	t.Log(len(s1), cap(s1))

	s1 = append(s1, 2)
	t.Log(s1)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	t.Log(len(s2), cap(s2), s2)

	// index=3的元素设置为1，注意是大括号
	s3 := []int{3:1}
	t.Log(s3)

	// nil切片
	var null []int
	// 空切片
	var empty []int = make([]int,0)

	t.Log(null, empty)

	s4 := make([]int, 5)
	s4[0] = 10
	s4[1] = 20
	s4[2] = 30
	t.Log("slice:", s4, " get i:", s4[2])

	str := SliceInt2String(s2)
	t.Log(str)

	// 添加一个切片，需要解包
	var a []int
	a = append(a, []int{1,2,3}...)
	// 在指定位置处添加
	a = append(a[:3], 4)
	// 嵌套调用
	a = append(append(a, 5),6)
	t.Log("appent test: ", a)

	// byte切片和string的互转
	text := "hello"
	b := []byte(text)
	t.Log(b)

	con := string(b)
	t.Log(con)

	t.Log("=====================copy==================")
	copyArr(t)

	t.Log("=====================delete================")
	deleteArr(t)
}

/**
删除用截取的方式实现的，可以理解为要保留哪些
 */
func deleteArr(t *testing.T) {
	src := make([]int, 10)
	for i := 0; i < 10; i++ {
		src[i] = i
	}
	// 删除第n个元素
	i := 3
	src = append(src[0:i-1], src[i:]...)
	t.Log(src)

	src = src[1:]
	t.Log(src)

}

func copyArr(t *testing.T)  {
	src := make([]int, 10)
	dest := make([]int, 10)
	part := make([]int, 10)
	ref := src

	for i := 0; i < 10; i++ {
		src[i] = i
	}
	copy(dest, src)

	src[0] = 100 // 不影响dest, 影响ref
	t.Log(dest)
	t.Log(ref)

	// 复制部分
	copy(part, src[0:5])
	t.Log(part)
}

func SliceInt2String(s []int) string {
	if len(s) < 1 {
		return ""
	}

	b := make([]byte, 0, 256)
	// b是byte切片，append的是一个个byte，所以字符串先转为byte[]，然后...表示打散byte[], 一个个添加到切片
/*	arr := []byte(strconv.Itoa(s[0]))
	b = append(b, arr...)*/
	b = append(b, strconv.Itoa(s[0])...) // 字符不转byte[]，直接打散成byte
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}

	return string(b)
}