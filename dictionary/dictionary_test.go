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
	t.Run("Add a new word", func(t *testing.T) {
		testWord := "teknonymy"
		testDefinition := "The custom of naming a parent after their child"
		dictionary := Dictionary{}
		dictionary.Add(testWord, testDefinition)
		assertDefinition(dictionary, testWord, testDefinition, t)
	})
	t.Run("Add an existing word", func(t *testing.T) {
		testWord := "teknonymy"
		testDefinition := "The custom of naming a parent after their child"
		dictionary := Dictionary{testWord: testDefinition}
		err := dictionary.Add(testWord, "Some other definition")
		assertError(err, ErrWordExists, t)
		got, _ := dictionary.Search(testWord)
		assertString(got, testDefinition, t)
	})
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

func assertString(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q instead of %q", got, want)
	}
}
