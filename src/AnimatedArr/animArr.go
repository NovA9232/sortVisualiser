package AnimatedArr

import (
  "math"
  "math/rand"
  "fmt"
  "time"
  "github.com/gen2brain/raylib-go/raylib"
)

var (
  // Base speeds (usually~ in time per comparison/change)
  QS_SLEEP = time.Millisecond  // Quick sort sleep time
  CHANGE_SLEEP = QS_SLEEP  // Time for changeDataBetween to sleep
  MS_SLEEP = time.Millisecond * 2  // Merge sort sleep time.
  BBL_SLEEP = time.Microsecond * 2  // Bubble sort sleep time
  INST_SLEEP = time.Microsecond * 2
  SHL_SLEEP = time.Millisecond * 2

  SHUFFLE_SLEEP = time.Microsecond * 500
)

type AnimArr struct {
	Data					[]float32
	sortedData		[]int
	lineNum				int
	lineWidth			int
	Active				int		// Index of current element being operated on.
	Active2				int   // Secondary active, for swapping elements.
	PivotInd			int   // For highlighting pivot when doing quickSort.
	nonLinearMult int
	ArrayAccesses int
	Comparisons		int
  W             float32
  H             float32
	maxValue			float32
	CurrentText   string
	Sorted				bool
	Sorting				bool
	Shuffling			bool
	linear				bool
	ColorOnly			bool // Do not show height if true
	Showcase			bool  // If showcase is running
}

func (a *AnimArr) Init(width, height float32, lineWidth int, linear, colorOnly bool, nonLinVarianceMult int) {  // nonLinVarianceMult is a multiplier for how variant the data is if linear is false
  a.W, a.H = width, height
	a.lineWidth = lineWidth
	a.lineNum = int(math.Floor(float64(a.W/float32(a.lineWidth))))

	a.Active		= -1
	a.Active2		= -1
	a.PivotInd	= -1
	a.Shuffling = false
	a.CurrentText = ""
	a.linear		= linear
	a.nonLinearMult = nonLinVarianceMult
	a.ColorOnly = colorOnly
	a.Sorted		= a.linear
	a.Sorting   = false

	QS_SLEEP = QS_SLEEP * time.Duration(a.lineWidth)
	CHANGE_SLEEP = QS_SLEEP
	MS_SLEEP = MS_SLEEP * time.Duration(a.lineWidth)
	BBL_SLEEP = BBL_SLEEP * time.Duration(a.lineWidth)
	INST_SLEEP = INST_SLEEP * time.Duration(a.lineWidth)
	SHL_SLEEP = SHL_SLEEP * time.Duration(a.lineWidth)

	if a.linear {
		a.Data = a.GenerateLinear(0, a.H, a.H/float32(a.lineNum))
	} else {
		a.Data = a.Generate(a.lineNum, a.lineNum*a.nonLinearMult)
	}
}

func RegularQuickSort(arr []float32) []float32 {  // Just a quick sort to sort the array for comparison (if needed)
	if len(arr) < 2 { return arr }
	var left	 []float32
	var middle []float32
	var right  []float32
	var pivot		 float32 = arr[len(arr)/2]

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			middle = append(middle, arr[i])
		}
	}

	return append(append(RegularQuickSort(left), middle...), RegularQuickSort(right)...)
}

func (a *AnimArr) getLineY(val float32) float32 {   // Lower case incase I want to have this as a package.
	return a.H-((float32(val)/float32(a.lineNum*a.nonLinearMult))*a.H)
}

func (a *AnimArr) drawLine(i int, colour rl.Color) {  // English spelling
	var x = float32((i*a.lineWidth)+(a.lineWidth/2))
	var y float32
	if a.ColorOnly {
		y = 0
	} else if a.linear {
		y = a.H-a.Data[i]
	} else {
		y = a.getLineY(a.Data[i])
	}
	rl.DrawLineEx(rl.NewVector2(x, a.H), rl.NewVector2(x, y), float32(a.lineWidth), colour)
}

func (a *AnimArr) Draw() {
	var clr rl.Color
	for i := 0; i < a.lineNum; i++ {
		if i == a.Active {
			clr = rl.Green
		} else if i == a.Active2 {
			clr = rl.Red
		} else if i == a.PivotInd {
			clr = rl.Yellow
		//} else if a.Sorted && !a.ColorOnly {   // Remove this to prevent the view going green when sorted.
		//	clr = rl.Lime
		} else {
			normal := uint8((a.Data[i]/a.maxValue)*255)  // Value normalised to 255
			//clr = rl.NewColor((normal/2)+127, (normal), (normal/3)+70, 255)  // Off yellow + coral
			//clr = rl.NewColor((normal/2)+127, (normal), (normal/3)+50, 255)  // Fire
			//clr = rl.NewColor(normal, normal, normal, 255)  // Grayscale
			//clr = rl.NewColor(normal, (normal/2)+127, normal/3, 255)  // Zesty (green --> yellow)
			clr = rl.NewColor(normal, (normal/3), (normal/2)+127, 255)  // Vapourwave/Twilight
		}
		a.drawLine(i, clr)
	}

	rl.DrawText(a.CurrentText, 10, 10, 30, rl.LightGray)

	if a.ArrayAccesses+a.Comparisons > 0 {
		rl.DrawText(fmt.Sprintf("Total length of array: %d", len(a.Data)), 10, 80, 20, rl.LightGray)
		if a.ArrayAccesses > 0 {
			rl.DrawText(fmt.Sprintf("Array accesses: %d", a.ArrayAccesses), 10, 40, 20, rl.LightGray)
		}
		if a.Comparisons > 0 {
			rl.DrawText(fmt.Sprintf("Comparisons: %d", a.Comparisons), 10, 60, 20, rl.LightGray)
		}
	}
}

func (a *AnimArr) Generate(num, max int) []float32 {	// Generates array
	a.maxValue = float32(max)
	var out []float32
	for i := 0; i < num; i++ {
		out = append(out, float32(rand.Intn(max)))
	}
	return out
}

func (a *AnimArr) GenerateLinear(start, finish, jump float32) []float32 {
	fmt.Println("Generating linear:", start, finish, jump)
	a.maxValue = finish
	var out []float32
	for i := start+jump; i <= finish; i += jump {
		out = append(out, i)
	}
	return out
}

func (a *AnimArr) Reverse() {
	for i, j := 0, len(a.Data)-1; i < j; i, j = i+1, j-1 {
		a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
	}
}

func (a *AnimArr) swapElements(i1, i2 int) {
	a.Data[i1], a.Data[i2] = a.Data[i2], a.Data[i1]
	a.ArrayAccesses += 2
}

func (a *AnimArr) Shuffle(times int, sleep bool, bogo bool) {
	a.Sorted = false
	a.Shuffling = true

	var max int = len(a.Data)
	for i := 0; i < times; i++ {
		if !bogo { a.CurrentText = fmt.Sprintf("Shuffling, round: %d", i+1) }
		for j := 0; j < len(a.Data); j++ {
			a.Active	= j
			a.Active2 = rand.Intn(max)
			a.swapElements(a.Active, a.Active2)
			if sleep { time.Sleep(SHUFFLE_SLEEP) }
		}
	}
	a.Shuffling = false
	a.Active	= -1
	a.Active2 = -1
	a.ArrayAccesses = 0
	a.Comparisons = 0
	if !bogo { a.CurrentText = "" }
}

func (a *AnimArr) changeDataBetween(start, end int, newSlice []float32, sleep bool, sleepTime *time.Duration) {
	for i := start; i < end; i++ {
		a.Active = i
		a.Data[i] = newSlice[i-start]
		a.ArrayAccesses++
		if sleep { time.Sleep(*sleepTime) }  // Sleep for half the time cause this bit should be quick
	}
}

func (a *AnimArr) cmpArrayWithData(array []float32) bool {  // Returns true if they are the same
	if len(a.Data) != len(array) {
		return false
	}

	for i := 0; i < len(a.Data); i++ {
		a.Active = i
		if a.Data[i] != array[i] {
			return false
		}
	}

	return true
}

func (a *AnimArr) BogoSort() {
	a.Sorted = false
	sorted := make([]float32, len(a.Data))
	copy(sorted, a.Data) // Copy the data into a new array
	sorted = RegularQuickSort(sorted)
	origShflSleep := SHUFFLE_SLEEP
	SHUFFLE_SLEEP = time.Millisecond * 10 // Only changed by this sort, so no point in making it an argument to a.Shuffle

	for !a.cmpArrayWithData(sorted) {
		a.Shuffle(1, true, true)
	}
	fmt.Println("Finally sorted.")
	SHUFFLE_SLEEP = origShflSleep
	a.Sorted = true
	a.Active = -1
}

func (a *AnimArr) QuickSort(start, end int) {   // Start and end of part of array to sort.
	if end-start > 1 {  // Not counting simple Comparisons like this, only between elements of a.Data
		var left []float32
		var middle []float32
		var right []float32
		a.PivotInd = int(math.Floor(float64(start+end)/2))
		var pivot float32 = a.Data[a.PivotInd]
		a.ArrayAccesses += 1

		for i := start; i < end; i++ {
			a.Active = i
			if a.Data[i] < pivot {
				a.Comparisons++
				left = append(left, a.Data[i])
				a.ArrayAccesses++
			} else if a.Data[i] > pivot {
				a.Comparisons++
				right = append(right, a.Data[i])
				a.ArrayAccesses++
			} else {
				a.Comparisons += 2 // Will have had to compare the other two, but wouldn't have added on.
				middle = append(middle, a.Data[i])
				a.ArrayAccesses++
			}
			a.ArrayAccesses++  // 1 in for loop
		}

		a.PivotInd = -1
		a.changeDataBetween(start, end, append(append(left, middle...), right...), true, &QS_SLEEP)
		a.QuickSort(start, start+len(left))
		a.QuickSort(start+len(left)+len(middle), start+len(left)+len(middle)+len(right))
	}
}

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

func (a *AnimArr) InsertionSort() {
	for i := 1; i < len(a.Data); i++ {
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


func (a *AnimArr) resetVals() {
	a.Sorting = false
	a.Active = -1
	a.Active2 = -1
	a.PivotInd = -1
	a.Sorted = true
}

func (a *AnimArr) DoSort(sort string) {
	a.Sorting = true
	a.Sorted = false
	a.ArrayAccesses = 0
	a.Comparisons = 0
	if sort == "quick" {
		a.CurrentText = "Quick Sort"
		go func() {
			a.QuickSort(0, len(a.Data))
			a.resetVals()
		}()
	} else if sort == "bogo" {
		a.CurrentText = "Bogo Sort"
		go func() {
			a.BogoSort()
			a.resetVals()
		}()
	} else if sort == "bubble" {
		a.CurrentText = "Bubble Sort"
		go func() {
			a.BubbleSort()
			a.resetVals()
		}()
	} else if sort == "insertion" {
		a.CurrentText = "Insertion Sort"
		go func() {
			a.InsertionSort()
			a.resetVals()
		}()
	} else if sort == "shell" {
		a.CurrentText = "Shell Sort"
		go func() {
			a.ShellSort()
			a.resetVals()
		}()
	} else if sort == "merge" {
		a.CurrentText = "Merge Sort"
		go func() {
			a.MergeSort(0, len(a.Data))
			a.resetVals()
		}()
	} else {
		panic("Invalid sort: "+sort)
	}
}

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
