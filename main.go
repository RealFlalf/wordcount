package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := os.Args[1]
	fmt.Println(text)
	res := len(strings.Fields(text))
	fmt.Printf("%d", res)
}