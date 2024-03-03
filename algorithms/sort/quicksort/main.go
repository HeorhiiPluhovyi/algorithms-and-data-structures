package main

import "fmt"

type Number interface {
	uint8 | uint16 | uint32 | uint | int |
		int8 | int16 | int32 | int64 | float32 | float64
}

// Function to partition the array and return the partition index
func partition[K Number](slc []K, low, high int) int {
	// Choosing the pivot
	pivot := slc[high]

	// Index of smaller element and indicates the right position of pivot found far
	i := low - 1

	for j := low; j <= high; j++ {
		// If current element is smaller than the pivot
		if slc[j] < pivot {
			// Increment index of smaller element
			i++
			slc[i], slc[j] = slc[j], slc[i]
		}
	}

	slc[i+1], slc[high] = slc[high], slc[i+1] // Swap pivot to its correct position
	return i + 1                              // Return the partition index
}

// The main function that implements QuickSort
func quickSort[K Number](slc []K, low, high int) {
	if low < high {
		// pi is the partitioning index, arr[pi] is now at the right place
		pi := partition(slc, low, high)

		// Separately sort elements before partition and after partition
		quickSort(slc, low, pi-1)
		quickSort(slc, pi+1, high)
	}
}

func main() {
	var slc = []int{10, 7, 8, 9, 1, 5, 8, 7, 67, 78, 99}
	var l = len(slc)
	fmt.Println(slc)
	quickSort(slc, 0, l-1)
	fmt.Println(slc)
}
