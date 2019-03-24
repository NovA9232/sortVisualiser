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

	for !a.Sorted && len(left) > 0 && len(right) > 0 {
		a.Active = start+len(left)+count
		if left[0] > right[0] {
			popped, right = right[0], right[1:]
			sorted = append(sorted, popped)
		} else {
			popped, left = left[0], left[1:]
			sorted = append(sorted, popped)
		}
		a.Comparisons++  // Will have had to compare it once regardless
		a.ArrayAccesses++
		count++
	}
	sorted = append(append(sorted, left...), right...)
	a.changeDataBetween(start, end, sorted, MS_SLEEP)
}

func (a *AnimArr) MergeSort(start, end int) {  // Using a quick sort to merge lists
	if !a.Sorted && end-start > 1 { // Not counting this one in a.Comparisons
		var mid int = int(math.Floor(float64(start+end)/2))
		a.MergeSort(start, mid)  // go through left
		a.MergeSort(mid, end)

		a.mergeArrays(start, mid, end, &MS_SLEEP)
		time.Sleep(MS_SLEEP)
		a.totalSleepTime += MS_SLEEP.Seconds()
	}
}
