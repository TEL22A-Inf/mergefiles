package stringlists

import "fmt"

func ExampleRemoveString() {
	l1 := []string{"A", "B", "A", "C", "D", "B", "E"}

	result1 := RemoveString(l1, "A")
	result2 := RemoveString(l1, "B")
	result3 := RemoveString(l1, "C")
	result4 := RemoveString(l1, "X")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// [B C D B E]
	// [A A C D E]
	// [A B A D B E]
	// [A B A C D B E]
}
