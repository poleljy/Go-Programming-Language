package main

import (
	_ "fmt"
	_ "os"

	_ "gopl/ch1"
	_ "gopl/ch10"
	"gopl/ch7"

	_ "unicode/utf8"
)

func main() {
	// ch1
	//ch1.TestDefer()

	// ch7
	ch7.TestHTTP2()

	// ch10
	/*
		if err := ch10.ToJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "jpeg: ", err)
			os.Exit(1)
		}
	*/
}
