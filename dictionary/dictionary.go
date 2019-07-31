package dictionary

const (
	// ErrNotFound for when the word cannot be found.
	ErrNotFound = DictionaryError("could not find the word you were looking for")

	// ErrWordExists when trying to add a word which already exists in the dict.
	ErrWordExists = DictionaryError("Attempted to add a word which already exists in the dictionary")

	// ErrWordDoesNotExist when trying to update a word which doesn't exist in the dict.
	ErrWordDoesNotExist = DictionaryError("Attempted to update a word which does not exist in the dictionary")
)

// DictionaryError covers problems in dictionary functions
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

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

// Update the stored definition of a word.
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}
