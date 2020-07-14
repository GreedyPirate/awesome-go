// package的文档说明
package day1

import (
	"reflect"
	"testing"
	"unicode"
)

func TestChar(t *testing.T) {
	zh := '中'
	t.Log(reflect.TypeOf(zh))

	en := 'a'
	t.Log(reflect.TypeOf(en))
	t.Log(unicode.IsLetter(en))

	num := en
	t.Log(num)

	word := "中国china"
	// 直接用len返回的是字节数，一个中文占3个字节
	t.Log("word byte length：", len(word))
	// 要想获取字符串的真正长度，需要强转为[]rune
	t.Log("word actual lenth: ", len([]rune(word)))

	// 类型转换
	f := 3.14
	i := int(f)
	t.Log(i)
}