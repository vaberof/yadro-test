package main

import (
	"fmt"
	"github.com/vaberof/yadro-test/internal/domain/ballsorter"
)

func main() {
	var n int
	fmt.Scan(&n)

	containers := make([][]int, n)

	for i := 0; i < n; i++ {
		containers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&containers[i][j])
		}
	}

	if ballsorter.CanSort(containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
