# os

## os/exec
import "os/exec"

### func LookPath
搜索可执行的二进制的文件路径，返回的是执行路径和error
```go
func LookPath(file string) (string, error) 
```
示例：
```go
path, err := exec.LookPath("godoc")
if err != nil {
	log.Fatal("installing godoc in your future")
}
fmt.Printf("godoc is available at %s\n", path)
```
源码：
```go
// ErrNotFound is the error resulting if a path search failed to find an executable file.
var ErrNotFound = errors.New("executable file not found in $path")

// Error records the name of a binary that failed to be executed
// and the reason it failed.
type Error struct {
	Name string
	Err  error
}

func (e *Error) Error() string {
	return "exec: " + strconv.Quote(e.Name) + ": " + e.Err.Error()
}

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}

// LookPath searches for an executable binary named file
// in the directories named by the path environment variable.
// If file begins with "/", "#", "./", or "../", it is tried
// directly and the path is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func LookPath(file string) (string, error) {
	// skip the path lookup for these prefixes
	skip := []string{"/", "#", "./", "../"}

	for _, p := range skip {
		if strings.HasPrefix(file, p) {
			err := findExecutable(file)
			if err == nil {
				return file, nil
			}
			return "", &Error{file, err}
		}
	}

	path := os.Getenv("path")
	for _, dir := range strings.Split(path, "\000") {
		if err := findExecutable(dir + "/" + file); err == nil {
			return dir + "/" + file, nil
		}
	}
	return "", &Error{file, ErrNotFound}
}
```

### func Command
输入文件的路径，参数字符串，返回的是*Cmd的结构