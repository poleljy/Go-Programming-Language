## Map

```go
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

// 刪除
delete(m, "b")

```

```go
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	personDB := make(map[string][2]PersonInfo)
	personDB["test"] = [2]PersonInfo{{"1", "Jack", "aaa"}, {"2", "Rose", "bbb"}}
	
	v, ok := personDB["test"]
	if !ok {
		fmt.Println("Key not found.")
		return
	}
	
	for i, v := range v {
		fmt.Println(i, v, v.ID, v.Address, v.Name)
	}
}
```