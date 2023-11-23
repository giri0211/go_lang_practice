package loops

import (
	"fmt"
	"strings"
)

func Collection_loop_test() {
	fmt.Println("Collection_loop_test")
	fmt.Println(strings.Repeat("_", 10))
	arr := [3]int{100, 101, 102}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Collection_loop_Only_index_test() {
	fmt.Println("Collection_loop_Only_index_test")
	fmt.Println(strings.Repeat("_", 10))
	arr := [3]int{100, 101, 102}
	for i, _ := range arr {
		fmt.Println(i)
	}
}

func Collection_loop_Only_value_test() {
	fmt.Println("Collection_loop_Only_value_test")
	fmt.Println(strings.Repeat("_", 10))
	arr := [3]int{100, 101, 102}
	for _, v := range arr {
		fmt.Println(v)
	}
}
