package main

import (
	"fmt"
	_ "os"
)

func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值  %d%s 超过100", n, "%")
		}
		fmt.Printf("数值 :%d%s\n", n, "%")
	}
	return nil
}

func TestFmt() {
	// Errorf
	if err := fmt.Errorf("Test errorf : %s", "success"); err != nil {
		//fmt.Fprintln(os.Stdout, err)
		fmt.Println(err)
	}

	// Print, Println, Printf
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")         // ab1 2 3cd
	fmt.Println("a", "b", 1, 2, 3, "c", "d")             // a b 1 2 3 c d
	fmt.Printf("Test fmt Printf: %s, %d\n", "string", 2) // Test fmt Printf: string, 2

	if err := percent(30, 70, 90, 160); err != nil {
		fmt.Println(err)
	}
}
