package iteration

func Repeat(character string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += character
	}
	return result
}
