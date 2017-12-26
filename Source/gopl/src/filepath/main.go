package main

import (
	"fmt"
	"path/filepath"
)

func Abs_Test(path string) {
	des, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("filepath.Abs(\"%s\") : %s \n", path, des)
}

func Base_Test(path string) {
	des := filepath.Base(path)
	fmt.Printf("filepath.Base(\"%s\") : %s \n", path, des)
}

func Clean_Test(path string) {
	des := filepath.Clean(path)
	fmt.Printf("filepath.Clean(\"%s\") : %s \n", path, des)
}

//
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

func main() {
	Abs_Test("main.go")
	Abs_Test("/root/main.go")

	Base_Test("main.go")
	Base_Test("/root/main.go")

	Clean_Test("./main.go")

	TestExt()

	ret, err := filepath.Glob("/root/workspace/Go/Go-Programming-Language/Blog/*.md")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range ret {
		fmt.Println(filepath.Base(file))
	}
}
