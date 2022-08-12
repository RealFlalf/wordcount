package main

import (
	"fmt"
	"os"
)

func main() {
	text := os.Args
	fmt.Println(len(text[1:]))
}