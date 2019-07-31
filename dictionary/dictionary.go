package dictionary

import (
	"errors"
)

var (
	// ErrNotFound for when the word cannot be found.
	ErrNotFound = errors.New("could not find the word you were looking for")

	// ErrWordExists when trying to add a word which already exists in the dict.
	ErrWordExists = errors.New("Attempted to add a word which already exists in the dictionary")
)

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
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
