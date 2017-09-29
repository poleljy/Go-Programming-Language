package main

import (
	"fmt"
)

// 局部变量是可以返回的，且返回后该空间不会释放
type T struct {
    i, j int
}

func a(i, j int) T {
    t := T { i, j}
    return t
}
func TestFunction(){
    t:= a(1, 2)
    fmt.Println(t)
}

