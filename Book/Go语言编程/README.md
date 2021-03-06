## 第4章 并发编程

### 4.1 并发基础
并发则意味着程序在运行时有多个执行上下文，对应着多个调用栈。我们知道每一个进程在运行时，都有自己的调用栈和堆，有一个完整的上下文，而操作系统在调度进程的时候，会保存被调度进程的上下文环境，等该进程获得时间片后，再恢复该进程的上下文到系统中。

线程之间通信只能采用共享内存的方式。为了保证共享内存的有效性，我们采取了很多措施，比如加锁等，来避免死锁或资源竞争。实践证明，我们很难面面俱到，往往会在工程中遇到各种奇怪的故障和问题。计算机科学家们在近40年的研究中又产生了一种新的系统模型，称为“消息传递系统”。

### 4.2 协程

### 4.3 goroutine
goroutine是Go语言中的轻量级线程实现，由Go运行时（runtime）管理。
在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine中并发执行。当被调用的函数返回时，这个goroutine也自动结束了。需要注意的是，如果这个函数有返回值，那么这个返回值会被丢弃。

### 4.4 并发通信
在工程上，有两种最常见的并发通信模型：共享数据和消息。

共享数据是指多个并发单元分别保持对同一个数据的引用，实现对该数据的共享。被共享的数据可能有多种形式，比如内存数据块、磁盘文件、网络数据等。在实际工程应用中最常见的无疑是内存了，也就是常说的共享内存。

Go语言提供的是另一种通信模型，即以消息机制而非共享内存作为通信方式。
消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。这有点类似于进程的概念，每个进程不会被其他进程打扰，它只做好自己的工作就可以了。不同进程间靠消息来通信，它们不会共享内存。Go语言提供的消息通信机制被称为channel。

> 不要通过共享内存来通信，而应该通过通信来共享内存。

### 4.5 channel
我们可以使用channel在两个或多个goroutine之间传递消息。channel是进程内的通信方式，因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用Socket或者HTTP等通信协议。

channel是类型相关的。也就是说，一个channel只能传递一种类型的值，这个类型需要在声明channel时指定。
```go
package main

import (
	"fmt"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	
	for _, ch := range(chs) {
		<-ch
	}
}
```
#### 4.5.1 基本语法
声明：

	var channelName chan ElementType

示例：
	
	var ch chan int
	var m map[string] chan bool

定义：

	ch := make(chan int)
写入：
	
	ch <- value

读取：
	
	value := <-ch
> 如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。

#### 4.5.2 select
```go
select {
case <-chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程
}
```
> select不像switch，后面并不带判断条件，而是直接去查看case语句;
> 每个case语句里必须是一个面向channel的操作;

#### 4.5.3 缓冲机制



### 4.8 同步
#### 4.8.2 全局唯一性操作
对于从全局的角度只需要运行一次的代码，比如全局初始化操作， Go语言提供了一个Once类型来保证全局的唯一性操作。 once的Do()方法可以保证在全局范围内只调用指定的函数一次，而且所有其他goroutine在调用到此语句时，将会先被阻塞，直至全局唯一的once.Do()调用结束后才继续。