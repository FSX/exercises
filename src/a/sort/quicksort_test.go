package sort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	expected := []int{1, 2, 3, 5, 7, 9}
	result := []int{1, 5, 7, 2, 9, 3}
	Quicksort(result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v\n", result, expected)
	}
}

var unsortedInts = []int{
	44, 47, 21, 77, 92, 82, 8, 1, 40, 83, 72, 100, 68, 60, 33, 55, 66, 15,
	89, 86, 94, 90, 59, 61, 30, 93, 63, 20, 9, 57, 43, 7, 11, 14, 69, 41,
	91, 52, 12, 80, 17, 19, 51, 58, 56, 13, 96, 75, 70, 42, 37, 25, 54, 32,
	29, 81, 18, 65, 46, 22, 78, 39, 34, 53, 16, 85, 76, 50, 67, 35, 4, 97,
	28, 48, 2, 27, 62, 10, 24, 98, 84, 5, 6, 95, 36, 38, 87, 73, 74, 79,
	23, 31, 26, 3, 45, 64, 71, 99, 49, 88}

var sortedInts = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

func BenchmarkQuicksort(b *testing.B) {
	result := append([]int(nil), unsortedInts...)
	Quicksort(result)

	if !reflect.DeepEqual(result, sortedInts) {
		b.Errorf("\nGot:\n%v\nExpected:\n%v\n", result, sortedInts)
	}
}
