### Package fmt

```
import "fmt"
```

#### Overview


#### Index

##### 返回error类型
`func Errorf(format string, a ...interface{}) error`
> 将参数列表 a 填写到格式字符串 format 的占位符中，并将填写后的结果转换为 error 类型返回

源码：
```
// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error {
 	return errors.New(Sprintf(format, a...))
} 

```
实例：
```
// Errorf
if err := fmt.Errorf("Test errorf : %s", "success"); err != nil {
	//fmt.Fprintln(os.Stdout, err)
	fmt.Println(err)
}
```
##### 写入标准输出
`func Print(a ...interface{}) (n int, err error)`
> 非字符串参数之间会添加空格，返回写入的字节数。


`func Printf(format string, a ...interface{}) (n int, err error)`
> 参数列表 a 填写到格式字符串 format 的占位符中


`func Println(a ...interface{}) (n int, err error)`
> 所有参数之间会添加空格，最后会添加一个换行符

实例：
```
// ab1 2 3cd
fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")


```


##### 写入特定输出
`func Fprint(w io.Writer, a ...interface{}) (n int, err error)`
`func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`
`func Fprintln(w io.Writer, a ...interface{}) (n int, err error)`
> 将参数列表 a [填写到格式字符串 format 的占位符中并将填写后的结果]写入 w 中，返回写入的字节数

##### 写入字符串
`func Sprint(a ...interface{}) string`
`func Sprintf(format string, a ...interface{}) string`
`func Sprintln(a ...interface{}) string`

3.
`func Fscan(r io.Reader, a ...interface{}) (n int, err error)`
`func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)`
`func Fscanln(r io.Reader, a ...interface{}) (n int, err error)`