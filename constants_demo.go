package main

import (
	"fmt"
)

// func constants_test() {
// 	const a = 42
// 	const b string = "hello constant"
// 	const concatString = "hello, " + " World"
// 	const c = b
// 	const (
// 		d = true
// 		e = 3.14
// 	)

// 	// if no value provided copy value form above
// 	const (
// 		f = "foo"
// 		g
// 	)

// 	// assing with 0
// 	const special = iota // 0
// }

func constants_demo() {
	const c = iota
	fmt.Println(c)

	// iota works based on index
	// below is const group
	const (
		d = 2 * 5
		e
		f = iota
		g = 10 * iota
	)
	fmt.Println(d, e, f, g)
}
