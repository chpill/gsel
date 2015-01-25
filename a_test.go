package main

import (
	"os"
)

func ExampleBasic() {

	os.Args = []string{"gsel", ".a.c", "test_data/simple.json"}
	main()
	// Output:
	// [true]
}
