## Make和New

1. make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
2. new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：new返回指针。
3. make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型（引用），而不是*T。
4. 本质来讲，导致这三个内建类型有所不同的原因是：引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。make返回初始化后的（非零）值。
5. 故make 是内建类型初始化的方法，例如：`s :=make([]int,len,cap)`  //这个切片在元素在超过10个时，底层将会发生至少一次的内存移动动作

### make

```go
// slice
v := make([]int, 10)		// len: 10 cap: 10 value: [0 0 0 0 0 0 0 0 0 0]
fmt.Println("len:", len(v), "cap:", cap(v), "value:", v)
	
v2 := make([]int, 5, 10)	// len: 5 cap: 10 value: [0 0 0 0 0]	
v2 = v2[:cap(v2)]		   // len: 10 cap: 10 value: [0 0 0 0 0 0 0 0 0 0]	
v2 = v2[1:]				 // len: 9 cap: 9 value: [0 0 0 0 0 0 0 0 0]
c := []int{1,2,3,4,5}	   // len: 5 cap: 5 value: [1 2 3 4 5]

// map
m := make(map[string]string)
m["a"] = "aa"
m["b"] = "bb"
m["a"] = "cc"

// 查找
if v, ok := m["a"]; ok {
	fmt.Println(v)
} else {
	fmt.Println("Key not found")
}

// 遍历
for k, v := range m {
	fmt.Println("Key:", k, "Value:", v)
}

// channel
// 声明
	var channelName chan ElementType
// 示例
	var ch chan int
	var m map[string] chan bool
// 定义
	// 无缓冲channel
	ch := make(chan type)
	// 缓冲channel (channel可以存储多少元素)
	ch := make(chan type, value)
```

### new
```go
type Test struct {
	fd int
	name string
	nepipe int
}

func NewTest1(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	f := Test{fd, name, 2}
	return &f
}

func NewTest3(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	return &Test{fd, name, 3}
}

func NewTest4(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	return &Test{name:name,fd:fd}
}

func NewTest2(fd int, name string) *Test {
	if fd < 0 {
		return nil
	}
	
	f := new(Test)
	f.fd = fd;
	f.name = name
	f.nepipe = 0
	return f
}

func main() {
	newfile := NewTest1(1,"Test")
	fmt.Printf("Test:fd>%d | name>%s | nepipe>%d \n", newfile.fd, newfile.name, newfile.nepipe)
}

```
