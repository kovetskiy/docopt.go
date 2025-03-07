package main

import (
	"fmt"
	"github.com/kovetskiy/docopt-go"
)

func main() {
	usage := `usage: type_assert [-x] [-y]`

	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arguments)

	var x = arguments["-x"].(bool) // type assertion required
	if x == true {
		fmt.Println("x is true")
	}

	y := arguments["-y"] // no type assertion needed
	if y == true {
		fmt.Println("y is true")
	}
	y2 := arguments["-y"]
	if y2 == 10 { // this will never be true, a type assertion would have produced a build error
		fmt.Println("y is 10")
	}
}
