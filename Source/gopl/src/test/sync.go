package main

import (
	"fmt"
	"sync"
	"time"
)

func f1() {
	fmt.Println("This is f1 function")
}

func f2() {
	fmt.Println("This is f2 function")
}

func f3() {
	fmt.Println("This is f3 function")
}

/////////////////////////////////////////////

var str string
var once sync.Once

func setup() {
	fmt.Println("setup begins.")
	str = "hello, " + time.Now().Format("2006-01-02 15:04:15")

	for i := 1; i <= 10; i++ {
		//time.Sleep(1e9)
		fmt.Print(".")
	}
	fmt.Println("\nsetup ends.")
}

func print(wg *sync.WaitGroup) {
	defer wg.Done()

	once.Do(setup)
	fmt.Println(str)
}

func TestSyncOnce() {
	// single go routine

	//	var once sync.Once
	//	once.Do(f1)
	//	once.Do(f2)
	//	once.Do(f3)

	// multi go routine

	var wg sync.WaitGroup
	wg.Add(2)
	go print(&wg)
	go print(&wg)
	wg.Wait()
}
