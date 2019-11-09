package services

import (
	"gotrain/GoTestingInteg/mygolang-testing/api/utils/sort"
)

const (
	privateConstant = "private"
)

// Sort takes a list of its and sorts them
func Sort(elements []int) {
	sort.BubbleSort(elements)
}
