package stringlists

import "fmt"

// ExampleJoin_common demonstrates the common case of joining two non empty lists.
func ExampleJoin_common() {
	// Create example lists.
	l1 := []string{"first", "second"}
	l2 := []string{"third", "fourth"}

	// Join the lists in both orderings.
	result1 := Join(l1, l2)
	result2 := Join(l2, l1)

	// Print the results.
	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// [first second third fourth]
	// [third fourth first second]
}

// ExampleJoin_empty demonstrates the case when one of the lists is empty.
// In addition to checking the result's contents, it also shows that new lists
// are actually created in all cases.
func ExampleJoin_one_empty() {
	// Create example lists.
	// One of them ist empty.
	l1 := []string{"first", "second"}
	l2 := []string{}

	// Join the lists in both orderings.
	result1 := Join(l1, l2)
	result2 := Join(l2, l1)

	// Print the results.
	fmt.Println(result1)
	fmt.Println(result2)

	// Change an element in each result and show that
	// nothing has been changed in the original list.
	result1[0] = "I am different"
	result2[1] = "I am different"
	fmt.Println(l1)
	fmt.Println(l2)

	// Output:
	// [first second]
	// [first second]
	// [first second]
	// []
}

func ExampleJoin_both_empty() {
	l1 := []string{}
	l2 := []string{}

	result1 := Join(l1, l2)
	result2 := Join(l2, l1)

	fmt.Println(result1)
	fmt.Println()
	fmt.Println(result2)

	// Output:
	// []
	//
	// []
}
