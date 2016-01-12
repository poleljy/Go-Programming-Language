# 2.5 面向对象
带有接受者的函数：`method`

用Rob Pike的话来说就是：

>"A method is a function with an implicit first argument, called a receiver."

method的语法如下：

	func (r ReceiverType) funcName(parameters) (results)
	
示例：

	package main
	import (
		"fmt"
		"math"
	)

	type Rectangle struct {
		width, height float64
	}

	type Circle struct {
		radius float64
	}

	func (r Rectangle) area() float64 {
		return r.width*r.height
	}

	func (c Circle) area() float64 {
		return c.radius * c.radius * math.Pi
	}


	func main() {
		r1 := Rectangle{12, 2}
		r2 := Rectangle{9, 4}
		c1 := Circle{10}
		c2 := Circle{25}

		fmt.Println("Area of r1 is: ", r1.area())
		fmt.Println("Area of r2 is: ", r2.area())
		fmt.Println("Area of c1 is: ", c1.area())
		fmt.Println("Area of c2 is: ", c2.area())
	}

在使用method的时候重要注意几点

- 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
- method里面可以访问接收者的字段
- 调用method通过`.`访问，就像struct里面访问字段一样
- 指针作为Receiver会对实例对象的内容发生操作,而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。
