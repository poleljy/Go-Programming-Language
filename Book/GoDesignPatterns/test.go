package main

import (
	"fmt"
)

func main() {
	addN := func(m int) {
		return func(n int) {
			return m + n
		}
	}

	//addFive := addN(5)
	result := addN(6)
	fmt.Println(result)
}
