package files

import (
	"fmt"
)

func ExampleReadLines() {
	readlines := ReadLines("testinput.txt")

	fmt.Println(readlines)

	// Output:
	// [one two three]
}
