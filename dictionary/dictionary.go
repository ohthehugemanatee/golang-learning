package dictionary

import (
	"errors"
)

// Dictionary of terms and definitions.
type Dictionary map[string]string

// Search a given dictionary for a word.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
