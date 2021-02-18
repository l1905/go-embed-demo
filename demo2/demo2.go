package main

import (
	_ "embed"
	"fmt"
)

// 主要目的： 打印文件自身

//go:embed demo2.go
var src string

func main() {
	fmt.Print(src)
}