package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Valid entry", func(t *testing.T) {
		word := "testEntry"
		definition := "This is a test entry"
		dictionary := Dictionary{word: definition}
		got, _ := dictionary.Search(word)
		if got != definition {
			t.Errorf("Got %q, wanted %q, given %q", got, definition, word)
		}
	})
	t.Run("Nonexistent entry", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("test")
		assertError(err, ErrNotFound, t)

	})
}

func assertError(got error, want error, t *testing.T) {
	t.Helper()
	if got.Error() == "" {
		t.Errorf("Did not get an error where one was expected")
	}
	if got != want {
		t.Errorf("Got the wrong error. Expected %q, got %q", want.Error(), got.Error())
	}
}
