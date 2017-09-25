#### Unicode
ASCII，更准确地说是美国的ASCII，使用`7bit`来表示128个字符：包含英文字母的大小写、数字、各种标点符号和设置控制符。

通用的表示一个Unicode码点的数据类型是int32，也就是Go语言中rune对应的类型；它的同义词rune符文正是这个意思。

#### UTF-8
UTF8是一个将Unicode码点编码为字节序列的变长编码。
UTF8编码使用1到4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。每个符号编码后第一个字节的高端bit位用于表示总共有多少编码个字节。如果第一个字节的高端bit为0，则表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容。如果第一个字节的高端bit是110，则说明需要2个字节；后续的每个高端bit都以10开头。更大的Unicode码点也是采用类似的策略处理。

> 一个汉字在UTF-8中占3个字节


#### 定义赋值
```
func TestStrings() {
	// 赋值
	var str string
	str = "TestString"

	ch := str[0]    // 第一个字符
	len := len(str) // 长度

	fmt.Println("ch[0]: ", ch, ", length:", len)
}
输出: ch[0]:  84 , length: 10
```

#### 中文字符
```
func TestChinese() {
	str := "测试中文"
	fmt.Println("length:", len(str))

	for i, s := range str {
		fmt.Println(i, "Unicode:", s, "string:", string(s))
	}

	// rune
	r := []rune(str)
	fmt.Println("rune:", r)

	for i := 0; i < len(r); i++ {
		fmt.Printf("r[%d]: %v, string:%s \n", i, r[i], string(r[i]))
	}
}

输出：
length: 12
0 Unicode: 27979 string: 测
3 Unicode: 35797 string: 试
6 Unicode: 20013 string: 中
9 Unicode: 25991 string: 文
rune: [27979 35797 20013 25991]
r[0]: 27979, string:测 
r[1]: 35797, string:试 
r[2]: 20013, string:中 
r[3]: 25991, string:文 
```

#### 获取总字节数len函数
```
func Len(v type) int
```
* len(string) : 返回的是字符串的字节数
* len(Array) : 数组的元素个数
* len(*Array): 数组指针中的元素个数,如果入参为nil则返回0
* len(Slice) : 数组切片中元素个数,如果入参为nil则返回0
* len(map) : 字典中元素个数,如果入参为nil则返回0
* len(Channel) : Channel buffer队列中元素个数

#### 字符串和Byte切片

标准库中有四个包对字符串处理尤为重要：`bytes`、`strings`、`strconv`和`unicode`包

```
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
```
字符串和字节slice之间可以相互转换：
```
s := "abc"
b := []byte(s)
s2 := string(b)
```

bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、
byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要处理化，因为零
值也是有效的