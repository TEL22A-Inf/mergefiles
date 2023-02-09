package stringlists

import "fmt"

// ExampleJoin_common_two demonstrates the common case of joining two non empty lists.
func ExampleJoin_common_two() {
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

// ExampleJoin_common_three demonstrates the common case of joining three non empty lists.
func ExampleJoin_common_three() {
	// Create example lists.
	l1 := []string{"first", "second"}
	l2 := []string{"third", "fourth"}
	l3 := []string{"fifth", "sixth"}

	// Join the lists in all possible orderings.
	result1 := Join(l1, l2, l3)
	result2 := Join(l1, l3, l2)
	result3 := Join(l2, l1, l3)
	result4 := Join(l2, l3, l1)
	result5 := Join(l3, l1, l2)
	result6 := Join(l3, l2, l1)

	// Print the results.
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// [first second third fourth fifth sixth]
	// [first second fifth sixth third fourth]
	// [third fourth first second fifth sixth]
	// [third fourth fifth sixth first second]
	// [fifth sixth first second third fourth]
	// [fifth sixth third fourth first second]
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
	fmt.Println(result2)

	// Output:
	// []
	// []
}

func ExampleJoin_no_lists() {
	result1 := Join()

	fmt.Println(result1)

	// Output:
	// []
}

func ExampleJoin_one_list() {
	l1 := []string{"first", "second"}
	result1 := Join(l1)

	fmt.Println(result1)

	// Change an element in the result and show that
	// nothing has been changed in the original list.
	result1[0] = "I am different"
	fmt.Println(l1)

	// Output:
	// [first second]
	// [first second]
}
