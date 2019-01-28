package animatedArr

import (
  "time"
)

func (a *AnimArr) showcaseRstr() {
	a.ArrayAccesses = 0
	a.Comparisons = 0
	a.Sorting = true
	a.Sorted = false
	a.Shuffle(2, true, false)
	time.Sleep(time.Second)
}

func (a *AnimArr) RunShowcase() {
	a.Showcase = true

	a.showcaseRstr()
	a.CurrentText = "Quick Sort"
	a.QuickSort(0, len(a.Data))
	a.resetVals()

	time.Sleep(time.Second)
	a.showcaseRstr()
	a.CurrentText = "Merge Sort"
	a.MergeSort(0, len(a.Data))
	a.resetVals()

	time.Sleep(time.Second)
	a.showcaseRstr()
	a.CurrentText = "Bubble Sort"
	a.BubbleSort()
	a.resetVals()

	time.Sleep(time.Second)
	a.showcaseRstr()
	a.CurrentText = "Insertion Sort"
	a.InsertionSort()
	a.resetVals()

	time.Sleep(time.Second)
	a.showcaseRstr()
	a.CurrentText = "Shell Sort"
	a.ShellSort()
	a.resetVals()

	a.ArrayAccesses = 0
	a.Comparisons = 0
	a.Showcase = false
}
