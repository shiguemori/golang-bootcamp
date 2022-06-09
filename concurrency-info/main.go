package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("No. of CPUs:", runtime.NumCPU())

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

}
