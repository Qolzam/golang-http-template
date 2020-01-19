package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("hello multi platform, I am %s-%s\n", runtime.GOOS, runtime.GOARCH)
}
