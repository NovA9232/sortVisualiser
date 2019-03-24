package animatedArr

import (
	"math"
)

func RegularQuickSort(arr []float32) []float32 {  // Just a quick sort to sort the array for comparison (if needed)
	if len(arr) < 2 { return arr }
	var (
		left	 []float32
		middle []float32
		right  []float32
		pivot		 float32 = arr[len(arr)/2]
	)

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			middle = append(middle, arr[i])
		}
	}

	return append(append(RegularQuickSort(left), middle...), RegularQuickSort(right)...)
}

func (a *AnimArr) QuickSort(start, end int) {   // Start and end of part of array to sort.
	if !a.Sorted && end-start > 1 {  // Not counting simple Comparisons like this, only between elements of a.Data
		var left []float32
		var middle []float32
		var right []float32
		a.PivotInd = int(math.Floor(float64(start+end)/2))
		var pivot float32 = a.Data[a.PivotInd]
		a.ArrayAccesses += 1

		for i := start; !a.Sorted && i < end; i++ {
			a.Active = i
			if a.Data[i] < pivot {
				a.Comparisons++
				left = append(left, a.Data[i])
			} else if a.Data[i] > pivot {
				a.Comparisons++
				right = append(right, a.Data[i])
				a.ArrayAccesses++
			} else {
				a.Comparisons += 2 // Will have had to compare the other two, but wouldn't have added on.
				middle = append(middle, a.Data[i])
			}
			a.ArrayAccesses++
		}

		a.PivotInd = -1
		a.changeDataBetween(start, end, append(append(left, middle...), right...), QS_SLEEP)
		a.QuickSort(start, start+len(left))
		a.QuickSort(start+len(left)+len(middle), start+len(left)+len(middle)+len(right))
	}
}
