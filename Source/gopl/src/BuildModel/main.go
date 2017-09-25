package main

import (
	"fmt"
	_ "os"
	"time"
)

func main() {
	//fmt.Fprintln(os.Stderr, "[error message]")

	time.Sleep(2 * time.Second)
	fmt.Println("[初始化]")
	fmt.Println("[0%]")

	//fmt.Fprintln(os.Stderr, "[初始化失败]")
	//os.Exit(100)

	time.Sleep(5 * time.Second)
	fmt.Println("[相机重建]")
	fmt.Println("[10%]")

	time.Sleep(6 * time.Second)
	fmt.Println("[生成密集点云]")
	fmt.Println("[15%]")

	time.Sleep(3 * time.Second)
	fmt.Println("[20%]")
	time.Sleep(3 * time.Second)
	fmt.Println("[30%]")

	time.Sleep(4 * time.Second)
	fmt.Println("[网格重建]")
	fmt.Println("[40%]")

	time.Sleep(3 * time.Second)
	fmt.Println("[45%]")
	time.Sleep(4 * time.Second)
	fmt.Println("[55%]")
	time.Sleep(2 * time.Second)
	fmt.Println("[65%]")

	time.Sleep(3 * time.Second)
	fmt.Println("[纹理优化]")
	fmt.Println("[70%]")
	time.Sleep(3 * time.Second)
	fmt.Println("[75%]")
	time.Sleep(2 * time.Second)
	fmt.Println("[80%]")

	time.Sleep(1 * time.Second)
	fmt.Println("[构建LOD]")
	fmt.Println("[85%]")
	time.Sleep(3 * time.Second)
	fmt.Println("[95%]")

	time.Sleep(2 * time.Second)
	fmt.Println("[完成创建]")
	fmt.Println("[100%]")
}
