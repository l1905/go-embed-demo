package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed version.txt
	version string

	Version string = strings.TrimSpace(version)

)

func main() {
	fmt.Printf("Version %q\n", Version)
}