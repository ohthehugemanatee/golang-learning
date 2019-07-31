package dictionary

import (
	"errors"
)

// ErrNotFound for when the word cannot be found.
var ErrNotFound = errors.New("could not find the word you were looking for")

// Dictionary of terms and definitions.
type Dictionary map[string]string

// Search a given dictionary for a word.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add a word/definition pair to the dictionary.
func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}
