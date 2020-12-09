package main

import (
	"algorithm-coding/concurrency"
	"bufio"
	"fmt"
	"os"
)

// HelloWorld.go entrance
func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	switch input {
	case "mu\n":
		concurrency.MuHelloWorld()
	case "ch\n":
		concurrency.ChHelloWorld()
	case "wg\n":
		concurrency.WgHelloWorld()
	default:
		fmt.Printf("Hello world in default")
	}
}
