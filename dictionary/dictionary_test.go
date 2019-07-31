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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	testWord := "teknonymy"
	testDefinition := "The custom of naming a parent after their child"
	dictionary.Add(testWord, testDefinition)
	assertDefinition(dictionary, testWord, testDefinition, t)
}

func assertDefinition(dictionary Dictionary, testWord string, testDefinition string, t *testing.T) {
	t.Helper()
	got, err := dictionary.Search(testWord)
	if err != nil {
		t.Errorf("Got error where none was expected: %q", err.Error())
	}
	if got != testDefinition {
		t.Errorf("Definition was not correct. Expected %q, got %q", testDefinition, got)
	}
}
