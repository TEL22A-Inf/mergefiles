package stringlists

// RemoveString takes a list of strings and a single string s.
// Removes all occurrences of s from the list and returns the result.
func RemoveString(list []string, s string) []string {
	result := []string{}
	if len(list) == 0 {
		return result
	}

	head, tail := list[0], list[1:]
	if head != s {
		result = append(result, head)
	}

	return append(result, RemoveString(tail, s)...)
}
