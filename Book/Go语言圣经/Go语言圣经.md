### 第一章

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