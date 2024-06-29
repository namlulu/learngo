package main

import (
	"fmt"
	"strings"

	"github.com/namlulu/learngo/something"
)

func mulitply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)

}

func main() {
	// 1.1 Packages and Imports
	something.SayHello()

	// 1.2 Variables and Constants
	const name string = "namlulu"
	// name = "namlulu2"

	var name2 string = "namlulu"
	name2 = "namlulu2"

	name3 := "namlulu3"

	fmt.Println(name, name2, name3)

	// 1.3 Functions part One
	fmt.Println(mulitply(2, 2))
	totalLength, upperName := lenAndUpper("namlulu")
	fmt.Println(totalLength, upperName)

}
