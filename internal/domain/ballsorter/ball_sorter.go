package ballsorter

import "slices"

func CanSort(containers [][]int) bool {
	n := len(containers)
	containersCapacity := make([]int, n)
	colorsCount := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			containersCapacity[i] += containers[i][j]
			colorsCount[i] += containers[j][i]
		}
	}

	slices.Sort(containersCapacity)
	slices.Sort(colorsCount)

	for i := 0; i < n; i++ {
		if containersCapacity[i] < colorsCount[i] {
			return false
		}
	}
	return true
}
