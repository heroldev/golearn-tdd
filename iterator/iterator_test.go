package iterator

import (
	"fmt"
	"testing"
)

// Tests the Repeat function in iterator.go
func TestIterator(test *testing.T) {

	assert := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	}

	test.Run("iterate a character 5 times", func(t *testing.T) {

		got := Repeat("a", 5)
		want := "aaaaa"

		assert(t, got, want)
	})

}

func BenchmarkIterator(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("y", 8))
	// Output: yyyyyyyy
}
