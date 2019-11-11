[package services

import (
	"fmt"
	"testing"
)

func Test_Constants(t *testing.T) {
	if privateConstant != "private" {
		t.Error("The constant is not private")
	}
}
func Test_BSort(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)

	BSort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 0 {
		t.Error("The lastElement should be 9. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 9 {
		t.Error("The first element should be 9, it is ", felement)
	}
	fmt.Println(elements)
}

func Test_Sort(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)

	Sort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 9 {
		t.Error("The lastElement should be 9. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 0 {
		t.Error("The first element should be 0, it is ", felement)
	}
	fmt.Println(elements)
}

func BenchmarkBubbleSort(b *testing.B) {

	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	for i := 0; i < b.N; i++ {
		BSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = 1
		j++
	}
	return result
}
func BenchmarkBubbleSortS(b *testing.B) {
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		BSort(elements)
	}
}

func BenchmarkBSortLimit(b *testing.B) {
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSortS(b *testing.B) {
	elements := getElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
