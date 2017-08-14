## 字符切片
```
import "bytes"
```
http://www.it165.net/pro/html/201407/16999.html
### 基本处理函数
* func Compare(a, b []byte) int
```go
// Compare
a := []byte("Golang Learning")
b := []byte("Golang Learning")
result := bytes.Compare(a, b)
switch {
case result > 0:
	fmt.Println("a greater b")
case result >= 0:
	fmt.Println("a greater or equal b")
case result < 0:
	fmt.Println("a less b")
case result <= 0:
	fmt.Println("a less or equal b")
}
```

* func Contains(b, subslice []byte) bool  
 检查字节切片b是否包含子切片subslice（区分大小写）
```go
package main

import (
    "bytes"
    "fmt"
)

func main() {
    s := []byte("Golang")
    subslice1 := []byte("go")
    subslice2 := []byte("Go")
    fmt.Println(bytes.Contains(s, subslice1))
    fmt.Println(bytes.Contains(s, subslice2))
}
```

### Buffer
默认情况下Buffer对象没有定义初始值，Buffer使用结构体自带的一个[64]byte数组作为存储空间。当超出限制时，另创建一个两倍的存储空间，并复制未读取的数据。当Buffer里的数据被完全读取后，会将写入位置重置到底层数据的开始处。因此只要读写操作平衡，就无须担心内存会持续增长。

