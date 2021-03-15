package iterator

// Repeat takes in a character and repeats it a given number of times
func Repeat(character string, repeatCount int) string {

	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated

}

func main() {
	Repeat("b", 12)
}
