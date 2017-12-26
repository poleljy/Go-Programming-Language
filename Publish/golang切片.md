### go 切片
----------
#### 引言
`切片(slice)`在`go`语言中的重要性不言而喻，不过在实际使用过程中还是可能会出现一些意想不到的“坑”，这些“坑”恰恰值得思考：  

1. 切片的基本操作：创建及初始化，增删改查等  
2. 切片作为函数参数传递：值传递、引用传递？  

----------
#### 切片的基本使用
1.切片基本概念  
`切片` 从命名上可以理解为从数组上切下来的一小片数据。底层还是数组。数组的特性：相同数据类型的集合，连续存放，随机读取，大小固定。引入切片的重要原因是为了消除数组大小固定的限制，使用者不需要手动指定切片大小，自动扩容。

![slice](https://i.imgur.com/WSr0pI5.png)

一个切片对象由三部分组成：指向底层数组的指针，当前切片可以访问的数据个数，当前切片底层数组存放数据的容量。
> 当length超过capacity后，会自动重新分配一个更大容量的数组，切片的指针会指向新数组，capacity也会自动增加。

2.基本操作

##### 创建及初始化

a)、`make` 函数创建  
调用 `make` 时，内部会分配一个数组，然后返回数组对应的切片。
`len()` 和 `cap()` 分别获取切片的大小和容量。 
``` go 
slice := make([]int, 3, 5) // length = 3, capcity = 5
fmt.Printf("slice addr : %p, len : %d, cap: %d, slice: %v \n", slice, len(slice), cap(slice), slice)
```
> slice addr : 0xc04203e3a0, len : 3, cap: 5, slice: [0 0 0] 

``` go
slice := make([]int, 5)     // length = 5, capcity = 5
fmt.Printf("slice addr : %p, len : %d, cap: %d, slice: %v \n", slice, len(slice), cap(slice), slice)
```
> slice addr : 0xc04203e3a0, len : 5, cap: 5, slice: [0 0 0 0 0] 

b)、带初始值
``` go
slice := []string{"Jack", "Rose", "Leonardo DiCaprio", "Kate Winslet"}
slice := []int{10, 20, 30}
slice := []string{4: "Rose"}
```
> slice addr : 0xc04203e3a0, len : 4, cap: 4, slice: [Jack Rose Leonardo Kate] 
> slice addr : 0xc04203e3a0, len : 3, cap: 3, slice: [10 20 30] 
> slice addr : 0xc04203e3a0, len : 5, cap: 5, slice: [    Rose] 

c)、创建空切片
有的时候函数返回值会返回一个空切片（pointer = nil, length = capacity = 0）
``` go
var slice []int     // 比较常见写法
slice := make([]int, 0)
slice := []int{}
```
> slice addr : 0xc04203e3a0, len : 0, cap: 0, slice: [] 

``` go
var slice []int
//slice[0] = 1       // index out of range,length = capacity = 0
slice = append(slice, 1)       // 空切片可以直接使用append追加
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
```

注：创建数组和切片的区别在于是否指定了长度
array := [3]int{10, 20, 30}
slice := []int{10, 20, 30}

##### 常见操作

a)、创建新切片
``` go
slice := []int{10, 20, 30, 40, 50}
slice[1] = 15   // 修改第二个数值为15

newSlice := slice[1:3]  // 切出一个新的切片对象，共用底层数组
```
> slice addr : 0xc04203e3a0, len : 5, cap: 5, slice: [10 15 30 40 50] 
> slice addr : 0xc04203e3c0, len : 2, cap: 4, slice: [15 30] 

![slice](https://i.imgur.com/gkTXqnY.png)

b)、修改切片值
``` go
slice := []int{10, 20, 30, 40, 50}
	
newSlice := slice[1:3]
newSlice[1] = 35  // 底层共用数组，任一切片的修改都会改变数组实际存储值，影响其他切片
fmt.Printf("slice addr : %p, len : %d, cap: %d, slice: %v \n", slice, len(slice), cap(slice), slice)
fmt.Printf("slice addr : %p, len : %d, cap: %d, slice: %v \n", newSlice, len(newSlice), cap(newSlice), newSlice)
```
> slice addr : 0xc04203e3a0, len : 5, cap: 5, slice: [10 20 35 40 50] 
> slice addr : 0xc04203e3c0, len : 2, cap: 4, slice: [20 35] 

c)、切片追加

当追加新的数值超过capacity的时候，会重新分配一个新的数组，切片指向的底层数组地址发生改变，容量扩充了。capacity的增长比例在数据小于1000的时候每扩容一次，capicity扩充2倍，超过1000之后为1.25。

除了追加数值，切片还可以直接追加切片 ： `slice := append(slice, s...)`

``` go
slice := []int{10, 20, 30, 40, 50}

newSlice := slice[1:3]
newSlice = append(newSlice, 60)

newSlice2 := slice[1:3]
newSlice2 = append(newSlice2, 60, 70, 80, 90)

fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", newSlice, len(newSlice), cap(newSlice), newSlice)
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", newSlice2, len(newSlice2), cap(newSlice2), newSlice2)
```
> slice addr : 0xc04203e3a0,len : 5,cap: 5,slice: [10 20 30 60 50] 
> slice addr : 0xc04203e3c0,len : 3,cap: 4,slice: [20 30 60]
> slice addr : 0xc04203e3e0,len : 6,cap: 8,slice: [20 30 60 70 80 90] 

``` go
var slice []int
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
for i:=0; i<10; i++{
	slice = append(slice, i)
	fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
}
```
> slice addr : 0x0,len : 0,cap: 0,slice: [] 
> slice addr : 0xc042044088,len : 1,cap: 1,slice: [0] 
> slice addr : 0xc0420440e0,len : 2,cap: 2,slice: [0 1] 
> slice addr : 0xc0420420c0,len : 3,cap: 4,slice: [0 1 2] 
> slice addr : 0xc0420420c0,len : 4,cap: 4,slice: [0 1 2 3] 
> slice addr : 0xc04206a100,len : 5,cap: 8,slice: [0 1 2 3 4] 
> slice addr : 0xc04206a100,len : 6,cap: 8,slice: [0 1 2 3 4 5] 
> slice addr : 0xc04206a100,len : 7,cap: 8,slice: [0 1 2 3 4 5 6] 
> slice addr : 0xc04206a100,len : 8,cap: 8,slice: [0 1 2 3 4 5 6 7] 
> slice addr : 0xc04207a080,len : 9,cap: 16,slice: [0 1 2 3 4 5 6 7 8] 
> slice addr : 0xc04207a080,len : 10,cap: 16,slice: [0 1 2 3 4 5 6 7 8 9] 

每次扩容会重新分配一个数组同时复制之前的数据，会带来额外的开销，所以如果在事先知道容量的情况下，最好在创建切片时候指定容量，避免内存拷贝
``` go
slice := make([]int, 0, 10)  // length = 0, capacity = 10
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
for i:=0; i<10; i++{
	slice = append(slice, i)
	fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
}
```
> slice addr : 0xc042076000,len : 0,cap: 10,slice: [] 
> slice addr : 0xc042076000,len : 1,cap: 10,slice: [0] 
> slice addr : 0xc042076000,len : 2,cap: 10,slice: [0 1] 
> slice addr : 0xc042076000,len : 3,cap: 10,slice: [0 1 2] 
> slice addr : 0xc042076000,len : 4,cap: 10,slice: [0 1 2 3] 
> slice addr : 0xc042076000,len : 5,cap: 10,slice: [0 1 2 3 4] 
> slice addr : 0xc042076000,len : 6,cap: 10,slice: [0 1 2 3 4 5] 
> slice addr : 0xc042076000,len : 7,cap: 10,slice: [0 1 2 3 4 5 6] 
> slice addr : 0xc042076000,len : 8,cap: 10,slice: [0 1 2 3 4 5 6 7] 
> slice addr : 0xc042076000,len : 9,cap: 10,slice: [0 1 2 3 4 5 6 7 8] 
> slice addr : 0xc042076000,len : 10,cap: 10,slice: [0 1 2 3 4 5 6 7 8 9] 

d)、移除指定下标元素
``` go
slice := []int{10, 20, 30, 40, 50}

// 删除切片元素remove element at index
index:=3;
slice=append(slice[:index],slice[index+1:]...)
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
```
> slice addr : 0xc04203e3a0,len : 4,cap: 5,slice: [10 20 30 50]

e)、指定位置插入数据
需要创建一个临时切片（底层新建了一个新的数组）保存后面数据
``` go
slice := []int{10, 20, 30, 40, 50}

// 方法1
index:=3;
tail := append([]int{}, slice[index:]...)
slice = append(slice[0:index], 35)
slice = append(slice, tail...)
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)

// 方法2
head := slice[:index:index] // 第三个index参数
head = append(head, 35)
slice = append(head, slice[index:]...)
    
```
> slice addr : 0xc042076000,len : 6,cap: 10,slice: [10 20 30 35 40 50] 
> slice addr : 0xc042060060,len : 6,cap: 6,slice: [10 20 30 35 40 50] 

f)、切片拷贝
`copy`函数可以用来拷贝切片,函数签名: `func copy(dst, src []Type) int`

``` go
slice := []int{10, 20, 30, 40, 50}
copySlice := make([]int, 4, 10)     // 拷贝时目标切片length需要大于等于原切片，否则会丢失数据或报错
if num := copy(copySlice, slice); num > 0 {
	fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", copySlice, len(copySlice), cap(copySlice), copySlice)
} else {
	fmt.Println("failed to copy slice")
}
fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
```
> slice addr : 0xc042076000,len : 4,cap: 10,slice: [10 20 30 40] 
> slice addr : 0xc042060030,len : 5,cap: 5,slice: [10 20 30 40 50] 

g)、遍历
比较简单，不做描述，需要注意的就是遍历的时候是值传递，会做拷贝，在某些需要修改值得情况下可以考虑使用index来动态修改

##### 进阶用法
a)、第三个index参数
`slice := source[i:j:k]`
第一、二个参数是新切片的起始下标，左开右闭
第三个参数表示新切片capacity的下标位置，即限制新切片的capacity
length = j - i 
capacity = k - i 

使用第三个参数的好处就是能够限制新切片容量大小，在某些场景下能自动的隔离开新切片和旧切片，使新切片的修改不会改变旧切片的数值
``` go
head := slice[:index:index] // 限制新切片容量等于当前长度
head = append(head, 35)     // 追加数据会导致新切片重新分配底层数据，和之前的切片数组分离
```

b)、切片作为函数参数传递
``` go
func main {
	slice := []int{10, 20, 30, 40, 50}
	slice = foo(slice)
	fmt.Printf("slice addr : %p,len : %d,cap: %d,slice: %v \n", slice, len(slice), cap(slice), slice)
}

func foo(s []int) []int {
	s[0] = 0
	return s
}
```
> slice addr : 0xc042060030,len : 5,cap: 5,slice: [0 20 30 40 50] 

如果弄懂上面的知识，值传递和引用传递就不是问题了，毫无疑问切片做参数是值传递。每次都会做一次拷贝，指向的还是同一个底层数组。
另外一个比较容易混淆的问题也可以很好的解释，切片作为参数的时候不需要做引用传值，在64位系统中一个切片只需要24字节（8+8+8），谁用谁知道。

#### 参考资料
1. Go in Action [ 4 Arrays, slices, and maps ]
2. The Go Programming Language [4.2. Slice]

