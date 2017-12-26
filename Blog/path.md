import "path/filepath"

``` go
func Abs(path string) (string, error)
```
Abs函数返回path代表的绝对路径，如果path不是绝对路径，会加入当前工作目录以使之成为绝对路径。因为硬链接的存在，不能保证返回的绝对路径是唯一指向该地址的绝对路径。
``` go
func Base(path string) string
```
Base函数返回路径的最后一个元素。在提取元素前会去掉末尾的路径分隔符。如果路径是""，会返回"."；如果路径是只有一个斜杆构成，会返回单个路径分隔符。
``` go
func Clean(path string) string
```
Clean函数通过单纯的词法操作返回和path代表同一地址的最短路径。
格式化路径

它会不断的依次应用如下的规则，直到不能再进行任何处理：
1. 将连续的多个路径分隔符替换为单个路径分隔符
2. 剔除每一个.路径名元素（代表当前目录）
3. 剔除每一个路径内的..路径名元素（代表父目录）和它前面的非..路径名元素
4. 剔除开始一个根路径的..路径名元素，即将路径开始处的"/.."替换为"/"（假设路径分隔符是'/'）

返回的路径只有其代表一个根地址时才以路径分隔符结尾，如Unix的"/"或Windows的`C:\`。
如果处理的结果是空字符串，Clean会返回"."。参见http://plan9.bell-labs.com/sys/doc/lexnames.html
``` go
func Dir(path string) string
```
Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；如果路径由1到多个路径分隔符后跟0到多个非路径分隔符字符组成，会返回单个路径分隔符；其他任何情况下都不会返回以路径分隔符结尾的路径。

```go
func EvalSymlinks(path string) (string, error)
```
EvalSymlinks函数返回path指向的符号链接（软链接）所包含的路径。如果path和返回值都是相对路径，会相对于当前目录；除非两个路径其中一个是绝对路径。

```go
func Ext(path string) string
```
Ext函数返回path文件扩展名。返回值是路径最后一个路径元素的最后一个'.'起始的后缀（包括'.'）。如果该元素没有'.'会返回空字符串。

``` go
type ExtTest struct {
	path, ext string
}

var exttests = []ExtTest{
	{"path.go", ".go"},
	{"path.tar.gz", ".gz"},
	{"a.dir/b", ""},
	{"a.dir/b.go", ".go"},
	{"a.dir/", ""},
}

func TestExt() {
	for _, test := range exttests {
		if x := filepath.Ext(test.path); x != test.ext {
			fmt.Printf("Ext(%q) = %q, want %q\n", test.path, x, test.ext)
		}
	}
}
```
``` go
func FromSlash(path string) string
```
FromSlash函数将path中的斜杠（'/'）替换为路径分隔符并返回替换结果，多个斜杠会替换为多个路径分隔符。

``` go
const sep = filepath.Separator

var slashtests = []PathTest{
	{"", ""},
	{"/", string(sep)},
	{"/a/b", string([]byte{sep, 'a', sep, 'b'})},
	{"a//b", string([]byte{'a', sep, sep, 'b'})},
}
```

```go
func Glob(pattern string) (matches []string, err error)
```
Glob函数返回所有匹配模式匹配字符串pattern的文件或者nil（如果没有匹配的文件）。pattern的语法和Match函数相同。pattern可以描述多层的名字，如/usr/*/bin/ed（假设路径分隔符是'/'）。

```go
func IsAbs(path string) bool
```
sAbs返回路径是否是一个绝对路径。

orders[index].DataList.Datas[i].DownloadUrl = nginx + info.Path[pos+len(svc.DataDir):]