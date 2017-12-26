#### 1.Append类型的方法
将各种类型转换为(带引号的)字符串后追加到 dst 尾部。
``` go
func AppendBool(dst []byte, b bool) []byte
func AppendInt(dst []byte, i int64, base int) []byte
func AppendUint(dst []byte, i uint64, base int) []byte
func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte

func AppendQuote(dst []byte, s string) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte
func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```

#### 2.Format类型的方法
``` go
func FormatBool(b bool) string
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

#### 3.Parse类型的方法 
``` go
func ParseBool(str string) (value bool, err error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
func ParseFloat(s string, bitSize int) (f float64, err error)
```

#### 4.Quote类型的方法
将 s 转换为双引号字符串
``` go
func Quote(s string) string
func QuoteToASCII(s string) string
func QuoteRune(r rune) string
func QuoteRuneToASCII(r rune) string

func Unquote(s string) (t string, err error)
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
```

#### 5.其他 
``` go
func Itoa(i int) string
将整数转换为十进制字符串形式（即：FormatInt(i, 10) 的简写）
func Atoi(s string) (int, error)
将字符串转换为十进制整数，即：ParseInt(s, 10, 0) 的简写）
```