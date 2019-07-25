package dictionary

// Search a given dictionary for a word.
func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}
