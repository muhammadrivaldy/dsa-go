package main

func getMax(arr []int) int {
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}
