package animatedArr

import (
	"math"
	"time"
)

func (a *AnimArr) generateShellSortGaps() []int {
	var out = []int{0}
	for i := int(math.Floor(float64(len(a.Data))/2)); i > 0; i = int(math.Floor(float64(i)/2)) {
		out = append(out, i)
	}
	a.ArrayAccesses++ // In for loop
	return out
}

func (a *AnimArr) ShellSort() {
	gapSequence := a.generateShellSortGaps()
	for _, gap := range gapSequence {
		for i := gap; i < len(a.Data); i++ {
			temp := a.Data[i]
			for a.Active = i; a.Active >= gap && a.Data[a.Active-gap] > temp; a.Active -= gap {
				a.Comparisons++  // Remember, not counting Comparisons unless they compare an element of a.Data
				a.Active2 = a.Active-gap
				a.Data[a.Active] = a.Data[a.Active2]
				time.Sleep(SHL_SLEEP)
				a.ArrayAccesses += 3 // 2+ One in for loop
			}
			a.Data[a.Active] = temp
			a.ArrayAccesses += 2
		}
	}
}
