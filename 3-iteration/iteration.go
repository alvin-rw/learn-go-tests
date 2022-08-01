package iteration

func Repeat(character string, counter int) string {
	var repeated string
	for i := 0; i < counter; i++ {
		repeated += character
	}

	return repeated
}
