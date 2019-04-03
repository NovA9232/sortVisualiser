package animatedArr

import (
	"math"
	"time"
)

func (a *AnimArr) CombSort() {
	gap := len(a.Data)
	shrink := 1.3
	a.Sorted = false

	for !a.Sorted {
		gap = int(math.Floor(float64(gap)/shrink))
		if gap <= 1 {
			gap = 1
			a.Sorted = true
		}

		for i := 0; i + gap < len(a.Data); i++ {
			a.Active = i
			a.Active2 = i + gap
			if a.Data[a.Active] > a.Data[a.Active2] {
				a.swapElements(a.Active, a.Active2)
				a.Sorted = false
				time.Sleep(COMB_SLEEP)
				a.totalSleepTime += COMB_SLEEP.Seconds()
			}
			a.ArrayAccesses += 2
			a.Comparisons++
		}
	}
}
