package stringlists

// Join expects two lists of strings and joins them.
// Returns the resulting list.
// It always returns a new list.
func Join(lists ...[]string) []string {
	result := []string{}

	if len(lists) == 0 {
		return result
	}

	if len(lists) >= 1 {
		result = append(result, lists[0]...)
	}

	return append(result, Join(lists[1:]...)...)
}
