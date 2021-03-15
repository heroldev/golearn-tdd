package integer

import (
	"fmt"
	"testing"
)

//Test the adder file.
func TestInteger(test *testing.T) {

	assert := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	}

	test.Run("testing the addition of numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assert(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println((sum))
	//Output: 6
}
