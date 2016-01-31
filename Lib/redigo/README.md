# redigo

## Installation

	go get github.com/garyburd/redigo/redis

## 建立连接
	conn , err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)

参数意义： 网络类型“tcp”、地址和端口、连接超时、读超时和写超时时间

示例
```go
conn, err := redis.DialTimeout("tcp", "192.168.1.139:6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to connect to database :", err)
		return
	}
	defer conn.Close()
	
	size, err := conn.Do("DBSIZE")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("size is %d.\n", size)
```

## 关闭连接
通过调用 `conn.Close()` 关闭连接。

## 基本命令执行

	Do(commandName string, args ...interface{}) (reply interface{}, err error)

在redis的协议中，都是按照字符流的，那么Do函数是如何进行序列化的呢？下面是其转换规则:

	Go Type                 Conversion
	
	[]byte                  Sent as is
	string                  Sent as is
	int, int64              strconv.FormatInt(v)
	float64                 strconv.FormatFloat(v, 'g', -1, 64)
	bool                    true -> "1", false -> "0"
	nil                     ""
	all other types         fmt.Print(v)

byte数组和字符串不变，整形和浮点数转换成对应的字符串，bool用1或者0表示，nil为空字符串

执行后得到的结果返回值的类型：

	Redis type              Go type
	
	error                   redis.Error
	integer                 int64
	simple string           string
	bulk string             []byte or nil if value not present.
	array                   []interface{} or nil if value not present.

从redis传回来得普通对象（整形、字符串、浮点数）, redis提供了类型转换函数供转换：

	func Bool(reply interface{}, err error) (bool, error)
	func Bytes(reply interface{}, err error) ([]byte, error)
	func Float64(reply interface{}, err error) (float64, error)
	func Int(reply interface{}, err error) (int, error)
	func Int64(reply interface{}, err error) (int64, error)
	func String(reply interface{}, err error) (string, error)
	func Strings(reply interface{}, err error) ([]string, error)
	func Uint64(reply interface{}, err error) (uint64, error)

示例：

	_, err = conn.Do("SET", "user:user0", 123)
	_, err = conn.Do("SET", "user:user1", 456)
	_, err = conn.Do("APPEND", "user:user0", 87)
	
	user0, err := redis.Int(conn.Do("GET", "user:user0"))
	user1, err := redis.Int(conn.Do("GET", "user:user1"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("user0 is", user0, "user1 is", user1)

## Pipelining

	Send(commandName string, args ...interface{}) error
	Flush() error
	Receive() (reply interface{}, err error)

Send writes the command to the connection's output buffer. 
Flush flushes the connection's output buffer to the server. 
Receive reads a single reply from the server.

	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	c.Receive() // reply from SET
	v, err = c.Receive() // reply from GET

## Publish and Subscribe

```go
func subscribe() {
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer c.Close()

    psc := redis.PubSubConn{c}
    psc.Subscribe("redChatRoom")
    for {
        switch v := psc.Receive().(type) {
        case redis.Message:
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
        case redis.Subscription:
            fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
        case error:
            fmt.Println(v)
            return
        }
    }
}
```

```go
go subscribe()
go subscribe()
go subscribe()
go subscribe()
go subscribe()
 
c, err := redis.Dial("tcp", "127.0.0.1:6379")
if err != nil {
     fmt.Println(err)
     return
 }
 defer c.Close()
 
for {
    var s string
    fmt.Scanln(&s)
     _, err := c.Do("PUBLISH", "redChatRoom", s)
    if err != nil {
        fmt.Println("pub err: ", err)
        return
     }
}
```


