package main

import (
	"os"
	"fmt"
)

func TestAlloc() {
	type T struct {
	n string
	i int
	f float64
	fd *os.File
	b []byte
	s bool
}
	
	t1 := new(T)
	fmt.Println("t1 new : ", t1)
	
	t2:= T{}
	fmt.Println("t2 init : ", t2)
	
	t3 := T{"hello", 1, 3.1415926, nil, make([]byte, 2), true}
	fmt.Println("t3 init : ", t3)
}