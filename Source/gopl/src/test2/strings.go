package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func TestStrings() {
	// 赋值
	var str string
	str = "TestString"

	ch := str[0]    // 第一个字符
	len := len(str) // 长度

	fmt.Println("ch[0]: ", ch, ", length:", len)
}

func TestChinese() {
	str := "测试中文"
	fmt.Println("length:", len(str))

	for i, s := range str {
		fmt.Println(i, "Unicode:", s, "string:", string(s))
	}

	// rune
	r := []rune(str)
	fmt.Println("rune:", r)

	for i := 0; i < len(r); i++ {
		fmt.Printf("r[%d]: %v, string:%s \n", i, r[i], string(r[i]))
	}
}

func TestUTF8() {
	s := "Hello,世界"
	fmt.Println(len(s))                    // "12"
	fmt.Println(utf8.RuneCountInString(s)) // "8"

	/*
		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("[%d] = %c\n", i, r)
			i += size
		}
	*/
	for i, r := range "Hello,世界" {
		fmt.Printf("[%d] = %c\n", i, r)
	}

	n := 0
	for range s { // for _, _ = range s
		n++
	}

	fmt.Printf("Count: %d\n", n)
}

// base name
func TestBaseName(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func TestBaseName2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// path/filepath

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
