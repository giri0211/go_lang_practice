package main

import (
	"demo/branching"
	"demo/functions"
	"demo/loops"
	"fmt"
)

func main() {
	a := "foo"
	fmt.Println(a)
	// execute the tests from main package.
	// loops_tests()
	// branching_tests()

}

func functions_tests() {
	functions.Test()
}
func branching_tests() {
	// branching.If_statement_test()
	// branching.Switch_test()
	// branching.Deffered_function_test()
	branching.Divide_test()
}
func loops_tests() {

	loops.Infinite_loop_test()
	loops.Counter_based_loop_test()
	loops.Collection_loop_test()
	loops.Collection_loop_Only_index_test()
	loops.Collection_loop_Only_value_test()
}

func webserver_tests() {

	fmt.Println("webserver started")
	// holding the control here for the we server to keep listening
	// web_service.Webserver_start()
}
func arithematic_tests() {

	// arthimatic_test()
	// constants_demo()
	// pointers()
	// cli_tool.CliDemo()
	// cli_demo.CliDemo_Test()

	// cli_tool.CliDemo_Test()
	// webserver_tests()
}
