package animatedArr

import (
	"math"
	"time"
)

func (a *AnimArr) generateShellSortGaps() []int {   // Generate A083318 gaps  O(n^(3/2))
	var out = []int{1}  // Init with 1
	for i, k := 0, 1; k < len(a.Data); i, k = i + 1, int(math.Ceil(math.Pow(2, float64(i)) + 1)) {
		println(k, "k")
		out = append(out, k)
	}
	return out
}

func (a *AnimArr) ShellSort() {
	gapSequence := a.generateShellSortGaps()
	for i, gap := len(gapSequence)-1, gapSequence[len(gapSequence)-1]; i >= 0; i, gap = i - 1, gapSequence[i] {  // Go through array backwards
		println("Gap:", gap)
		for i := gap; !a.Sorted && i < len(a.Data); i++ {
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
