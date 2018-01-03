### go 切片
----------
#### 引言
`map` 在golang中也是一个比较常见且重要的数据类型，可以从下面二个方面掌握 `map` 这一基础知识：
1. `map` 基本操作及注意事项；
2. go 1.9中 `goroutine` 安全的 `sync.map` 使用；  
----------
#### 基础操作及注意事项
##### 基本概念
`map` 的定义和特性在 [effective_go](https://golang.org/doc/effective_go.html#maps) 中有段话比较精准：
> Maps are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value) The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined on them. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

`Map`是go语言中高效的 `key/value` 存储数据类型,复杂度为O(1)。`key`支持任意可比较的数据类型，例如整形，浮点型，复数，字符串，指针，接口类型，结构体和数组。切片不能直接作为 `key` 值，因为切片不能直接进行比较查找。和切片类似，`map` 内部保留了指向底层数据结构的引用，如果 `map` 作为函数参数传递，函数内的修改也会直接影响调用者的数据。

##### 常见操作
a)、创建及初始化
``` go
var people map[string]int
//people["Pole"] = 27 	// panic: assignment to entry in nil map
fmt.Printf("map addr:%p,len:%d,nil:%v,content:%v \n", people, len(people), people == nil, people)

people = make(map[string]int)
fmt.Printf("map addr:%p,len:%d,nil:%v,content:%v \n", people, len(people), people == nil, people)

people["Pole"] = 27
fmt.Printf("map addr:%p,len:%d,nil:%v,content:%v \n", people, len(people), people == nil, people)

people = map[string]int {
	"Jack": 25,
	"Rose": 22,
}
fmt.Printf("map addr:%p,len:%d,nil:%v,content:%v \n", people, len(people), people == nil, people)

people["Rose"] = 23
fmt.Printf("map addr:%p,len:%d,nil:%v,content:%v \n", people, len(people), people == nil, people)

fmt.Printf("Pole is %d years old \n",people["Pole"])
people["Pole"]++
fmt.Printf("map addr:%p,len:%d,nil:%v,content:%v \n", people, len(people), people == nil, people)
```
>map addr:0x0,len:0,nil:true,content:map[] 
map addr:0xc042074060,len:0,nil:false,content:map[] 
map addr:0xc042074060,len:1,nil:false,content:map[Pole:27] 
map addr:0xc042074090,len:2,nil:false,content:map[Jack:25 Rose:22] 
map addr:0xc042074090,len:2,nil:false,content:map[Jack:25 Rose:23] 
Pole is 0 years old 
map addr:0xc042074090,len:3,nil:false,content:map[Jack:25 Rose:23 Pole:1]

没有什么难点， 只需要注意下面3点：
1. 声明的空(或`nil`)  `map` 不能直接添加数据；
2. 相同key值会赋值会被覆盖，不存在key时做运算会直接添加数据；
3. 对于不存在的key值获取value的时候会返回value对应类型的默认值，如int为0，string为空字符串

b)、获取指定key对应的value
``` go
key := "Rose"
if value, ok := people[key]; ok {
	fmt.Printf("%s is %d years old.\n", key, value)
} else {
	fmt.Printf("No people named %s\n", key)
}

key = "Poleljy"
if value := people[key]; value != 0 {
	fmt.Printf("%s is %d years old.\n", key, value)
} else {
	fmt.Printf("No people named %s\n", key)
}

// 遍历map
for key, value := range people {
	fmt.Printf("%s is %d years old.\n", key, value)
}
```
> Rose is 23 years old.  
> No people named Poleljy.
> Rose is 23 years old.
> Pole is 1 years old.
> Jack is 25 years old.

需要注意遍历 map 返回结果是无序的。

c)、移除数据 `delete(map, key)`

----------
#### `sync.map` 的使用

go 1.6之后，并发地读写map会报错,内建的 `map` 不是线程(goroutine)安全的
``` go
people := map[string]int {
	"Jack": 25,
	"Rose": 22,
}

var wg sync.WaitGroup
go func() {
	wg.Add(1)
	defer wg.Done()
	for {
		people["Jack"] = 27
	}
}()

go func() {
	wg.Add(1)
	defer wg.Done()
	for {
		fmt.Println(people["Jack"])
	}
}()
wg.Wait()
```
> fatal error: concurrent map read and map write

go 1.9之前的解决方案就是加 `sync.RWMutex`
``` go
people := struct {
	sync.RWMutex
	m map[string]int
} {
	m: make(map[string]int),
}

var wg sync.WaitGroup
go func() {
	wg.Add(1)
	defer wg.Done()
	for {
		people.Lock()
		people.m["Jack"] = 27
		people.Unlock()
	}
}()

go func() {
	wg.Add(1)
	defer wg.Done()
	for {
		people.RLock()
		fmt.Println(people.m["Jack"])
		people.RUnlock()
	}
}()
//time.Sleep(2*time.Second)
wg.Wait()
```
##### `sync.map`的基本使用方法
主要函数签名有：
``` go
// 移除一个键值对
func (m *Map) Delete(key interface{})
// 返回对应value,存在返回value，true;不存在返回nil，false
func (m *Map) Load(key interface{}) (value interface{}, ok bool)
// 存在对应key则直接返回value,否则存储键值对并返回value。load的返回true,store返回false
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
// 遍历所有键值对
func (m *Map) Range(f func(key, value interface{}) bool)
// 存储一个键值对
func (m *Map) Store(key, value interface{})
```    
示例：
``` go
package main

import(
	"fmt"
	"sync"
)

type userInfo struct {
	name string
	age int
}

func main() {
	print := func(key, value interface{}) bool {
		fmt.Printf("%s is %v years old.\n", key, value)
		return true
	}

	var m sync.Map

	key := "Jack"
	if val, ok := m.Load(key);ok {
		fmt.Printf("%s is %d years old.\n", key, val)
	} else {
		fmt.Printf("No people named %s.\n", key)
	}

	if val, ok := m.LoadOrStore(key, 25);ok {
		fmt.Printf("%s is %d years old.\n", key, val)
	} else {
		fmt.Printf("Insert people named %s, %d years old.\n", key, val)
	}
	m.Store("Rose", "twenty")	// 可以存储不同类型的value，orz...
	m.Delete("Rose")			// 移除key/value
	m.Range(print)				// 遍历

	user := userInfo{
		name: "Pole",
		age:27,
	}
	m.Store("three", user)
	m.Range(func(key, value interface{}) bool {
		if user,ok := value.(userInfo); ok {
			fmt.Printf("Amazeing, %s is %v years old.\n", user.name, user.age)
		} else {
			fmt.Printf("%s is %v years old.\n", key, value)
		}
		return true
	})
}
```
> No people named Jack.
> Insert people named Jack, 25 years old.
> Jack is 25 years old.
> Jack is 25 years old.
> Amazeing, Pole is 27 years old.

接口通俗易懂，没有难点。sync.map能存放任意类型的value，interface{}接口类型，这点感觉相当灵活。

#### 参考资料
1. [effective_go](https://golang.org/doc/effective_go.html#maps)
2. [go doc](https://golang.org/pkg/sync/#Map)
3. Go in Action(4.3 Map internals and fundamentals)
4. The Go Programming Language [4.3. Map]

