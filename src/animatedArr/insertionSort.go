package animatedArr

import (
	"time"
)

func (a *AnimArr) InsertionSort() {
	for i := 1; !a.Sorted && i < len(a.Data); i++ {
		a.PivotInd = i
		for j := i; j > 0 && a.Data[j-1] > a.Data[j]; j-- {
			a.Comparisons++
			a.Active = j
			a.Active2 = j-1
			a.ArrayAccesses += 2 // In for loop
			a.swapElements(j, j-1)
			time.Sleep(INST_SLEEP)
		}
	}
}
