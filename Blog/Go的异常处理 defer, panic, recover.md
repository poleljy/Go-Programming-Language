### 异常处理
在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为0了）。Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
```go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	
	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}
```
### defer
defer的思想类似于C++中的析构函数，不过Go语言中“析构”的不是对象，而是函数，defer就是用来添加函数结束时执行的语句。注意这里强调的是添加，而不是指定，因为不同于C++中的析构函数是静态的，Go中的defer是动态的。
```go
func f() (result int) {
 defer func() {
    result++  
 }() 
 return 0
}
```
> defer可以多次，这样形成一个defer栈，后defer的语句在函数返回时将先被调用

### panic
panic 是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。
函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。

### recover
recover之后，逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回。

### 检查函数是否发生了panic
```go
func ThrowPanic(f func()) (b bool) {
	defer func() {
		if e := recover(); e != nil {
			b = true
		}
	}()
	f()
	return
}
```