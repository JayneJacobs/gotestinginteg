package services

import (
	"fmt"
	"testing"
)

func Test_Constants(t *testing.T) {
	if privateConstant != "private" {
		t.Error("The constant is not private")
	}
}
func Test_Sort(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)

	Sort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 0 {
		t.Error("The lastElement should be 0. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 9 {
		t.Error("The first element should be 9, it is ", felement)
	}
	fmt.Println(elements)
}
