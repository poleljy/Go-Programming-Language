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