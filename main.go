package main

import (
	"fmt"
	"strings"

	"github.com/namlulu/learngo/something"
)

func mulitply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (lengtht int, uppercase string) {
	defer fmt.Println("I'm done")

	lengtht = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	default:
		return false
	}
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

	// 1.3 Functions part One, 1.4 Two
	fmt.Println(mulitply(2, 2))

	totalLength, upperName := lenAndUpper("namlulu")
	fmt.Println(totalLength, upperName)
	repeatMe("namlulu", "namlulu2", "namlulu3", "namlulu4")

	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)

	// 1.6 if with a Twist
	fmt.Println(canIDrink(15))

	// 1.7 Switch
	fmt.Println(canIDrink2(15))

	// 1.8 Pointers
	a := 2
	b := &a
	fmt.Println(*b)

	*b = 2024
	fmt.Println(a)

	// 1.9 Arrays and Slices
	names := []string{"namlulu", "namlulu2", "namlulu3", "namlulu4", "namlulu5"}
	names = append(names, "namlulu6")
	fmt.Println(names)

	// 1.10 Maps
	namlulu := map[string]string{"name": "namlulu", "age": "12"}
	for _, value := range namlulu {
		fmt.Println(value)
	}

	// 1.11 Structs
	type person struct {
		name    string
		age     int
		favFood []string
	}

	favFood := []string{"kimchi", "ramen"}
	sean := person{name: "namlulu", age: 18, favFood: favFood}
	fmt.Println(sean)
}
