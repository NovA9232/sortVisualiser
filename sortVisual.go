package main

import (
	"math"
	"math/rand"
	"time"
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_WIDTH  float32 = 1600
	SCREEN_HEIGHT float32 = 800

	//violet rl.Color = NewColor(61, 38, 69, 255)
	//raspberry rl.Color = NewColor(131, 33, 97, 255)
	//coral rl.Color = NewColor(218, 65, 103, 255)
	audioStream rl.AudioStream

	SORT_SLEEP = time.Millisecond
	BBL_SLEEP = time.Microsecond
	SHUFFLE_SLEEP = time.Millisecond/4
)

const (
	maxSamples = 22050
	maxSamplesPerUpdate = 4096
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
	maxValue			float32
	Sorted				bool
	Sorting				bool
	Shuffling			bool
	linear				bool
	colorOnly			bool // Do not show height if true
}

func (a *AnimArr) Init(lineWidth int, linear, colorOnly bool, nonLinVarianceMult int) {  // nonLinVarianceMult is a multiplier for how variant the data is if linear is false
	a.lineWidth = lineWidth
	a.lineNum = int(math.Floor(float64(SCREEN_WIDTH/float32(a.lineWidth))))
	a.Active		= -1
	a.Active2		= -1
	a.PivotInd	= -1
	a.Shuffling = false
	a.linear		= linear
	a.nonLinearMult = nonLinVarianceMult
	a.colorOnly = colorOnly
	a.Sorted		= a.linear
	a.Sorting   = false

	if a.linear {
		a.Data = a.GenerateLinear(0, SCREEN_HEIGHT, SCREEN_HEIGHT/float32(a.lineNum))
	} else {
		a.Data = a.Generate(a.lineNum, a.lineNum*a.nonLinearMult)
	}
}

func regularQuickSort(arr []float32) []float32 {  // Just a quick sort to sort the array for comparison (if needed)
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

	return append(append(regularQuickSort(left), middle...), regularQuickSort(right)...)
}

func (a *AnimArr) getLineY(val float32) float32 {   // Lower case incase I want to have this as a package.
	return SCREEN_HEIGHT-((float32(val)/float32(a.lineNum*a.nonLinearMult))*SCREEN_HEIGHT)
}

func (a *AnimArr) drawLine(i int, colour rl.Color) {  // English spelling
	var x = float32(i*a.lineWidth)
	var y float32
	if a.colorOnly {
		y = 0
	} else if a.linear {
		y = SCREEN_HEIGHT-a.Data[i]
	} else {
		y = a.getLineY(a.Data[i])
	}
	rl.DrawLineEx(rl.NewVector2(x, SCREEN_HEIGHT), rl.NewVector2(x, y), float32(a.lineWidth), colour)
}

func (a *AnimArr) Draw(dt float32) {
	var clr rl.Color
	for i := 0; i < a.lineNum; i++ {
		if i == a.Active {
			clr = rl.Green
		} else if i == a.Active2 {
			clr = rl.Red
		} else if i == a.PivotInd {
			clr = rl.Yellow
		} else if a.Sorted && !a.colorOnly {
			clr = rl.Lime
		} else {
			normal := uint8((a.Data[i]/a.maxValue)*255)  // Value normalised to 255
			clr = rl.NewColor(normal, normal, normal, 255)
		}
		a.drawLine(i, clr)
	}
}

func (a *AnimArr) playSound(index int) {

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
	for i := start; i <= finish; i += jump {
		out = append(out, i)
	}
	return out
}

func (a *AnimArr) swapElements(i1, i2 int) {
	a.Data[i1], a.Data[i2] = a.Data[i2], a.Data[i1]
}

func (a *AnimArr) Shuffle(times int, sleep bool) {
	a.Sorted = false
	a.Shuffling = true
	//fmt.Println("Shuffling...")
	var max int = len(a.Data)
	for i := 0; i < times; i++ {
		for j := 0; j < len(a.Data); j++ {
			a.Active	= rand.Intn(max)
			a.Active2 = rand.Intn(max)
			a.swapElements(a.Active, a.Active2)
			if sleep { time.Sleep(SHUFFLE_SLEEP) }
		}
		//fmt.Println("Done:", i, "lots.")
	}
	//fmt.Println("Finished shuffling.")
	a.Shuffling = false
	a.Active	= -1
	a.Active2 = -1
}

func (a *AnimArr) changeDataBetween(start, end int, newSlice []float32, sleep bool) {
	for i := start; i < end; i++ {
		a.Active = i
		a.Data[i] = newSlice[i-start]
		if sleep { time.Sleep(SORT_SLEEP/2) }  // Sleep for half the time cause this bit should be quick
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
	sorted = regularQuickSort(sorted)

	for !a.cmpArrayWithData(sorted) {
		SHUFFLE_SLEEP = time.Millisecond * 5 // Only changed by this sort, so no point in making it an argument to a.Shuffle
		a.Shuffle(1, true)
	}
	fmt.Println("Finally sorted.")
	SHUFFLE_SLEEP = time.Millisecond/2
	a.Sorted = true
	a.Active = -1
}

func (a *AnimArr) QuickSort(start, end int) {   // Start and end of part of array to sort.
	if end-start > 1 {
		var left []float32
		var middle []float32
		var right []float32
		a.PivotInd = int(math.Floor(float64(start+end)/2))
		var pivot float32 = a.Data[a.PivotInd]

		for i := start; i < end; i++ {
			a.Active = i
			if a.Data[i] < pivot {
				left = append(left, a.Data[i])
			} else if a.Data[i] > pivot {
				right = append(right, a.Data[i])
			} else {
				middle = append(middle, a.Data[i])
			}
			time.Sleep(SORT_SLEEP)
		}

		a.PivotInd = -1
		a.changeDataBetween(start, end, append(append(left, middle...), right...), true)
		a.QuickSort(start, start+len(left))
		a.QuickSort(start+len(left)+len(middle), start+len(left)+len(middle)+len(right))
	}
}

func (a *AnimArr) BubbleSort() {
	a.Sorted = false
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(a.Data)-1; i++ {
			a.Active = i
			a.Active2 = i+1
			if a.Data[i] > a.Data[i+1] {
				a.Data[i], a.Data[i+1] = a.Data[i+1], a.Data[i]
				sorted = false
			}
			time.Sleep(BBL_SLEEP)
		}
	}
	a.Sorted = true
	a.Active = -1
	a.Active2 = -1
}


//func (a *AnimArr) MergeSort(start, end int) {

//}

func (a *AnimArr) DoSort(sort string) {
	a.Sorting = true
	if sort == "quick" {
		go func() {
			a.Sorted = false
			a.QuickSort(0, len(a.Data))
			fmt.Println("Finished sort.")
			a.Active = -1
			a.Sorted = true
			a.Sorting = false
		}()
	} else if sort == "bogo" {
		go func() {
			a.BogoSort()
			a.Sorting = false
		}()
	} else if sort == "bubble" {
		go func() {
			a.BubbleSort()
			a.Sorting = false
		}()
	}
}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Sort Visualiser")
	rl.SetTargetFPS(144)

	//rl.InitAudioDevice()
	//audioStream = rl.InitAudioStream(22050, 32, 1)
	//rl.PlayAudioStream(audioStream)

	//audioData := make([]float32, maxSamples)

	//audioStream = rl.NewAudioStream(44000, 16, 1)
	anim := AnimArr{}
	anim.Init(2, true, true, 1)  // Input line thickness, if it is linear, and if it is color only here

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyR) && !anim.Shuffling {  // When 'r' is pressed, shuffle the array.
			go func() {
				anim.Shuffle(4, true)
			}()
		}

		if rl.IsKeyPressed(rl.KeyS) && !anim.Sorting {  // When 's' is pressed, sort the array.
			anim.DoSort("quick")
		}

		//if rl.IsAudioBufferProcessed(audioStream) {  // Refil audio stream
		//}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		anim.Draw(0)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
