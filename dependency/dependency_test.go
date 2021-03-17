package dependency

import (
	"bytes"
	"testing"
)

func TestGreeting(test *testing.T) {

	test.Run("test writer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}

	})

}
