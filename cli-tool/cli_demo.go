package cli_tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Name string = "Girish"
var privateName string = "private string"

func CliDemo_Test() {
	fmt.Println("what do you want ?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}

func privatefunc() {
	fmt.Println(privateName)
}
