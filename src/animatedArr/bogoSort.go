package animatedArr

import (
	"time"
)

func (a *AnimArr) BogoSort() {
	a.Sorted = false
	sorted := make([]float32, len(a.Data))
	copy(sorted, a.Data) // Copy the data into a new array
	sorted = RegularQuickSort(sorted)
	origShflSleep := SHUFFLE_SLEEP
	SHUFFLE_SLEEP = time.Millisecond * 10 // Only changed by this sort, so no point in making it an argument to a.Shuffle

	for !a.Sorted && !a.cmpArrayWithData(sorted) {
		a.Shuffle(1, true, true)
	}
	SHUFFLE_SLEEP = origShflSleep
	a.Sorted = true
	a.Active = -1
}
