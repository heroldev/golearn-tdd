package maps

import (
	"testing"
)

// tests relating to searching the directory
func TestSearch(test *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	test.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStringsEqual(t, got, want)

	})

	test.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("word")
		want := "this word is not in the dictionary"

		assertError(t, err, ErrWordNotFound)
		assertStringsEqual(t, err.Error(), want)

	})

}

//test adding words to the dictionary
func TestAdd(test *testing.T) {

	test.Run("adding words", func(t *testing.T) {

		dictionary := Dictionary{}

		dictionary.Add("test", "this is just a test")

		assertDefinition(t, dictionary, "test", "this is just a test")

	})

	test.Run("existing words", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})

}

func TestUpdate(test *testing.T) {

	test.Run("updating a word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)

	})

	test.Run("updating existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})

}

func TestDelete(test *testing.T) {

	test.Run("deleting a word", func(t *testing.T) {

		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrWordNotFound {
			t.Errorf("expected %q to be deleted", word)
		}

	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q but wanted %q", got, definition)
	}

}

func assertStringsEqual(t testing.TB, got, want string) {

	t.Helper()

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}

}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
