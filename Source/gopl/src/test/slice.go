package main

import (
	"fmt"
)

func changeSlice(s []int) []int {
	fmt.Printf("func slice addr: %p \n", &s)

	s[1] = 111
	return s
}

func TestSlice() {
	slice := make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		slice[i] = i
	}
	fmt.Println("slice: ", slice)

	fmt.Println("----------------------------------")

	//	fmt.Printf("slice: %v, addr: %p \n", slice, &slice)
	//	ret := changeSlice(slice)
	//	fmt.Printf("slice: %v, ret: %v, slice addr: %p, ret addr: %p \n", slice, ret, &slice, &ret)
	//	fmt.Printf("slice[0] addr:%p, ret[0] addr: %p \n", &slice[0], &ret[0])

	fmt.Println("----------------------------------")

	// arr := [5]int{0, 1, 2, 3, 4}
	// fmt.Println(arr)

	//	slice = arr[1:4]
	//	slice2 := arr[2:5]
	//
	//	fmt.Printf("arr %v, slice1 %v, slice2 %v, %p %p %p\n", arr, slice, slice2, &arr, &slice, &slice2)
	//	fmt.Printf("arr[2]%p slice[1] %p slice2[0]%p\n", &arr[2], &slice[1], &slice2[0])
	//	arr[2] = 2222
	//	fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)
	//	slice[1] = 1111
	//	fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)

	fmt.Println("----------------------------------")
	slice2 := slice[1:4]

	slice4 := make([]int, len(slice2))

	copy(slice4, slice2)

	fmt.Printf("slice %v, slice4 %v \n", slice, slice4)
	slice[1] = 1111
	fmt.Printf("slice %v, slice4 %v \n", slice, slice4)
}
