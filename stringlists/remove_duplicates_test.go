package stringlists

import "fmt"

func ExampleRemoveDuplicates() {
	l1 := []string{"A", "B", "A", "C", "D", "B", "E"}

	result1 := RemoveDuplicates(l1)

	fmt.Println(result1)

	// Output:
	// [A B C D E]
}
