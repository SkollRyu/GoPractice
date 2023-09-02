package sum

func Sum(arr []int) int {
	var result int
	// for i := 0; i < len(arr); i++ {
	// 	result += arr[i]
	// }

	for _, num := range arr {
		result += num
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, arr := range numbersToSum {
		sums = append(sums, Sum(arr))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, arr := range numbersToSum {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(arr[1:]))
		}
	}
	return sums
}
