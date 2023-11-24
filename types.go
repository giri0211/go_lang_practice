package main

import "fmt"

func testTypes() {
	var a string

	a = "foo"
	fmt.Println(a)

	// implicit
	var c = true
	fmt.Println(c)

	d := 3.14
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)
}
