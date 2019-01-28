package animatedArr

import (
	"math"
	"time"
)

func (a *AnimArr) mergeArrays(start, mid, end int, sleepTime *time.Duration) {
	var sorted []float32
	left := make([]float32, mid-start)
	right := make([]float32, end-mid)
	copy(left, a.Data[start:mid])
	copy(right, a.Data[mid:end])
	a.ArrayAccesses += len(a.Data)
	var popped float32
	var count int

	for len(left) > 0 && len(right) > 0 {
		a.Active = start+len(left)+count
		if left[0] > right[0] {
			a.Comparisons++
			popped, right = right[0], right[1:]
			sorted = append(sorted, popped)
		} else {
			a.Comparisons++ // Add it anyway
			popped, left = left[0], left[1:]
			sorted = append(sorted, popped)
		}
		a.ArrayAccesses++
		count++
	}
	sorted = append(append(sorted, left...), right...)
	a.changeDataBetween(start, end, sorted, true, &MS_SLEEP)
}

func (a *AnimArr) MergeSort(start, end int) {  // Using a quick sort to merge lists
	if end-start > 1 { // Not counting this one in a.Comparisons
		var mid int = int(math.Floor(float64(start+end)/2))
		a.MergeSort(start, mid)  // go through left
		a.MergeSort(mid, end)

		a.mergeArrays(start, mid, end, &MS_SLEEP)
		time.Sleep(MS_SLEEP)
	}
}
