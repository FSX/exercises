package sorting

// Quicksort as described in:
// https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme
func Quicksort(A []int) {
	if len(A) > 0 {
		p := qsPartition(A)
		Quicksort(A[:p])
		Quicksort(A[p+1:])
	}
}

func qsPartition(A []int) int {
	pivot := A[0]
	i := 0
	j := len(A) - 1

	for {
		for A[j] > pivot {
			j--
		}
		for A[i] < pivot {
			i++
		}

		if i < j {
			A[i], A[j] = A[j], A[i]
		} else {
			return j
		}
	}
}
