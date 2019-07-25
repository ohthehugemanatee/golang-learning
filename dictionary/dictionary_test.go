package dictionary

import "testing"

func TestSearch(t *testing.T) {
	given := "testEntry"
	want := "This is a test entry"
	dictionary := Dictionary{given: want}
	got := dictionary.Search(given)
	if got != want {
		t.Errorf("Got %q, wanted %q, given %q", got, want, given)
	}
}
