package services

import (
	"gotrain/GoTestingInteg/mygolang-testing/api/utils/sort"
)

const (
	privateConstant = "private"
)

// BSort takes a list of its and sorts them
func BSort(elements []int) {
	sort.BubbleSort(elements)
}

// Sort takes a list of its and sorts them
func Sort(elements []int) {
	sort.Sort(elements)
}

// BSortLimit Limits by using native sort for elements > 10000
func BSortLimit(elements []int) {
	if len(elements) < 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
