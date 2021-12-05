package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello World", Calculate(5))
	var x int = 3
	var a [4]int
	b := []int{5, 4, 3, 2, 1}
	b = append(b, 0)
	c := make(map[string]int)
	c["ok"] = 1
	c["sure"] = 2
	delete(c, "sure")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i, v := range b {
		fmt.Println(i, sum(v, 10))
	}

	p := person{name: "erik", age: 41}
	fmt.Println(x, a, b, c, p)

	// pointers

	n := 7
	inc(&n)
	fmt.Println("pointers: ", n)
}

func inc(n *int) {
	*n++
}
