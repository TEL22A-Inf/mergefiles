package stringlists

// Join expects two lists of strings and joins them.
// Returns the resulting list.
// It always returns a new list.
func Join(list1, list2 []string) []string {
	result := []string{}

	result = append(result, list1...)
	result = append(result, list2...)

	return result
}
