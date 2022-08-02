package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("look up a word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("didn't expect error but but for one")
		}
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("look up a word that does not exist", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
