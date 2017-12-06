package main

import (
	"fmt"
	_ "os"

	_ "gopl/ch1"
	_ "gopl/ch10"
	"gopl/ch7"

	_ "unicode/utf8"
)

func print(name []byte) {
	fmt.Println(string(name))
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func main() {
	//	name := []byte("ok")
	//	defer print(name)
	//
	//	name = []byte("no")
	//	print(name)

	// ch1
	//ch1.TestDefer()

	// ch7
	//ch7.TestHTTP2()

	//ch7.TestSleep()

	ch7.TestInterface()

	// ch10
	/*
		if err := ch10.ToJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "jpeg: ", err)
			os.Exit(1)
		}
	*/

	//fmt.Println(f())
}
