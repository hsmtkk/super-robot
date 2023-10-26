package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("begin")
	fmt.Println("output environment variables")
	for _, kv := range os.Environ() {
		fmt.Println(kv)
	}
	fmt.Println("end")
}
