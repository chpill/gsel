package main

import (
	"os"
	"strings"
)

func ExampleWithFileFlag() {
	cmd := "gsel -f test_data/simple.json .a.c"
	os.Args = strings.Split(cmd, " ")
	main()
	// Output:
	// [true]
}

func ExampleWithStdio() {
	var _ error
	// open the file and put it as standard input
	os.Stdin, _ = os.Open("test_data/simple.json")

	cmd := "gsel .a.c"
	os.Args = strings.Split(cmd, " ")
	main()
	// Output:
	// [true]
}
