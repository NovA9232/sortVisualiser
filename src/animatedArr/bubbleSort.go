package animatedArr

import (
	"time"
)

func (a *AnimArr) BubbleSort() {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(a.Data)-1; i++ {
			a.Active = i
			a.Active2 = i+1
			if a.Data[i] > a.Data[i+1] {
				a.swapElements(i, i+1)
				time.Sleep(BBL_SLEEP)
				sorted = false
			}
			a.Comparisons++ // Add in case
			a.ArrayAccesses += 2
		}
	}
	a.Sorted = true
	a.Active = -1
	a.Active2 = -1
}
