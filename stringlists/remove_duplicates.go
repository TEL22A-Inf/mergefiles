package stringlists

// RemoveDuplicates takes a list of strings.
// For each string in the list, it removes all duplicates of that string,
// that follow later in the list.
// Returns the result
func RemoveDuplicates(list []string) []string {
	result := []string{}

	if len(list) >= 1 {
		head, tail := list[0], list[1:]
		result = append(result, head)
		result = append(result, RemoveDuplicates(RemoveString(tail, head))...)
	}

	return result
}
