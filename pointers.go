package main

import "fmt"

func pointers() {
	s := "hello world"
	p := &s
	var x *string = &s
	fmt.Println(p, x)
	fmt.Println(*p)
	*p = "updated Hellow"

	fmt.Println(s)

	p = new(string)
	fmt.Println(p, *p)

}
