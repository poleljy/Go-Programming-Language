#### 第一章

1. map是用make函数创建的数据结构的一个引用。
当一个map被作为参数传递给一个函数时，函数接收到的是一份引用的拷贝，虽然本身并不是一个东西，但因为他们指向的是同一块数据对象。
（译注：类似于C++里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）

```go
// 从文件读取
func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) > 0 {
		for _, file := range files {
			fd, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Open file %s error: %s\n", file, err)
				continue
			}
			CountLine(fd, counts)
		}
	} else {
		CountLine(os.Stdin, counts)
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("string:%s,count:%d\n", line, count)
		}
	}
}

func CountLine(file *os.File, counts map[string]int) {
	line := bufio.NewScanner(file)
	for {
		ret := line.Scan()
		if ret != true {
			return
		}
		counts[line.Text()]++
	}
}
```

#### 第四章 复合数据类型

数组、slice、map和结构体

### 数组
因为数组的长度是固定的，因此在Go语言中很少直接使用数组。和数组对应的类型是`Slice`（切片），它是可以增长和收缩动态序列，`Slice`功能也更灵活
```
type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func TestSlice() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
}
```

### Slice
一个slice由三个部分构成：指针、长度和容量
要注意的是`slice`的第一个元素并不一定就是数组的第一个元素
```
months := [...]string{1: "January", /* ... */, 12: "December"}
数组的第一个元素从索引0开始，但是月份一般是从1开始的，因此我们声明数组时直接跳过第0个元素，第0个元素会被自动初始
化为空字符串。
```
如果切片操作超出cap(s)的上限将导致一个panic异常，但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大

### 4.3. Map
基础操作
``` go
ages := make(map[string]int) // mapping from strings to ints

ages := map[string]int{
	"Jack":25,
	"Rose":22,
}

ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"

delete(ages, "alice") // remove element ages["alice"]

for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}

if age, ok := ages["bob"]; !ok { /* ... */ }
```

排序
``` go
import (
	"sort"
	"fmt"
)

func main() {
	ages := make(map[string]int)

	//var names []string
	names := make([]string, 0, len(ages)) //给slice分配一个合适的大小将会更有效

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for name, age := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
```

处理key值是slice
``` go
var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
```


#### 第七章 接口
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口  
接口指定的规则非常简单：表达一个类型属于某个接口只要这个类型实现这个接口

> *os.File类型实现了io.Reader，Writer，Closer，和ReadWriter接口;   
> *bytes.Buffer实现了Reader，Writer，和ReadWriter这些接口

一个类型持有一个方法:  
* Go中的函数接收者，可以为值类型，也可以是引用类型；  
* 对于每一个命名过的具体类型T，它一些方法的接收者是类型T本身然而另一些则是一个T的指针

在T类型的参数上调用一个T的方法是合法的，只要这个参数是一个变量；编译器隐式的获取了它的地址。但这仅仅是一个语法糖：T类型的值不拥有所有*T指针的方法，那这样它就可能只实现更少的接口。

#### 第十一章 测试
``` go
import (
	"testing"
)
```

在`*_test.go`文件中，有三种类型的函数：`测试函数`、`基准测试函数`、`示例函数`

`测试函数`是以`Test`为函数名前缀的函数，用于测试程序的一些逻辑行为是否正确；go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。

``` go
import "testing"

func TestXXX(t testing.T) { }
```

`基准测试函数`是以`Benchmark`为函数名前缀的函数，它们用于衡量一些函数的性能；go test命令会多次运行基准函数以计算一个平均的执行时间。
``` go
import "testing"

func BenchmarkXXX(b *testing.B) { }
```

`示例函数`是以`Example`为函数名前缀的函数，提供一个由编译器保证正确性的示例文档,示例函数没有函数参数和返回值。