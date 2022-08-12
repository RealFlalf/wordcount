package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var text string
	// fmt.Scanf("%s", &text)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	words := strings.Fields(text)
	fmt.Println(len(words))
}