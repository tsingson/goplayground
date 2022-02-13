package main

import (
	"runtime"
)

func main() {
	// ballast := make([]byte, 10*1024*1024*1024) // 10G
	ballast := make([]byte, 10<<30)
	// do something

	runtime.KeepAlive(ballast)
}
