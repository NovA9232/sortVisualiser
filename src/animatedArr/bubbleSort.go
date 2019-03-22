package animatedArr

import (
	"time"
)

func (a *AnimArr) mainBubble(i int, swapped *bool, sleepTime *time.Duration) {  // Used by cocktail shake as well, so keep it separate.
  a.Active = i
  a.Active2 = i+1
  if a.Data[i] > a.Data[i+1] {
    a.swapElements(i, i+1)
    *swapped = true
    time.Sleep(*sleepTime)
  }
  a.Comparisons++ // Add in case
  a.ArrayAccesses += 2
}

func (a *AnimArr) BubbleSort() {
	swapped := true
	for swapped && !a.Sorted {
		swapped = false
		for i := 0; i < len(a.Data)-1; i++ {
			a.mainBubble(i, &swapped, &BBL_SLEEP)
		}
	}
	a.Sorted = true
	a.Active = -1
	a.Active2 = -1
}
