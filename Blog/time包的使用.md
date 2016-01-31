
## 定时器timer
```go
type Timer struct {
    C <-chan Time
    // contains filtered or unexported fields
}
```
The Timer type represents a single event. 
When the Timer expires, the current time will be sent on C, unless the Timer was created by AfterFunc. 
A Timer must be created with NewTimer or AfterFunc.