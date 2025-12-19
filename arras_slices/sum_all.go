package main

func SumAll(arrays ...[]int) []int {
	var sums []int
	for _, array := range arrays {
		sums = append(sums, Sum(array))
	}
	return sums
}
func SumAllTails(arrays ...[]int) []int {
	var sum []int
	for _, array := range arrays {
		if len(array) == 0 {
			sum = append(sum, 0)
			continue
		}
		sum = append(sum, Sum(array[1:]))
	}
	return sum
}
