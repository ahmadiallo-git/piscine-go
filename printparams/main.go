package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args[1:]); i++ {
		fmt.Printf(os.Args[i])
	}
}
