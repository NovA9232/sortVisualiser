package animatedArr

import (
  "fmt"
  "math/rand"
  "time"
)

func (a *AnimArr) Shuffle(times int, sleep bool, bogo bool) {
	a.Sorted = false
	a.Shuffling = true

	var max int = len(a.Data)
	for i := 0; !a.Sorted && i < times; i++ {
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

func (a *AnimArr) swapElements(i1, i2 int) {
	a.Data[i1], a.Data[i2] = a.Data[i2], a.Data[i1]
	a.ArrayAccesses += 2
}

func (a *AnimArr) changeDataBetween(start, end int, newSlice []float32, sleep bool, sleepTime *time.Duration) {
	for i := start; i < end; i++ {
		a.Active = i
		a.Data[i] = newSlice[i-start]
		a.ArrayAccesses++
		if sleep {
			time.Sleep(*sleepTime)
			a.totalSleepTime += sleepTime.Seconds()
		}
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

func (a *AnimArr) Reverse(array []float32) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
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
	a.maxValue = finish
	var out []float32
	for i := start+jump; i <= finish; i += jump {
		out = append(out, i)
	}
	return out
}
