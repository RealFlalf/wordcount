package main

import (
	"fmt"
	"os"
)

func main() {
	text := os.Args[1]
	fmt.Println(len(text))
}