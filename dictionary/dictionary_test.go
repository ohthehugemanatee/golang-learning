package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Valid entry", func(t *testing.T) {
		word := "testEntry"
		definition := "This is a test entry"
		dictionary := Dictionary{word: definition}
		got := dictionary.Search(word)
		if got != definition {
			t.Errorf("Got %q, wanted %q, given %q", got, definition, word)
		}
	})
}
