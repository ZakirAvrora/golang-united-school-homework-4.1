package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)

	if len(runes) < 2 {
		return input
	}
	i, j := 0, len(runes)-1

	for i !=j {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) 
}
