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

	helpW int32 = 533
	helpH int32 = int32(math.Floor(float64(SCREEN_HEIGHT)/4))

	//violet rl.Color = NewColor(61, 38, 69, 255)
	//raspberry rl.Color = NewColor(131, 33, 97, 255)
	//coral rl.Color = NewColor(218, 65, 103, 255)

	// Base speeds (~ in time per comparison/change)
	QS_SLEEP = time.Millisecond  // Quick sort sleep time
	CHANGE_SLEEP = time.Millisecond/2  // Time for changeDataBetween to sleep
	BBL_SLEEP = time.Microsecond // Bubble sort sleep time
	INST_SLEEP = time.Microsecond * 2

	SHUFFLE_SLEEP = time.Microsecond * 500
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
	CurrentText   string
	Sorted				bool
	Sorting				bool
	Shuffling			bool
	linear				bool
	colorOnly			bool // Do not show height if true
	Showcase			bool  // If showcase is running
}

func (a *AnimArr) Init(lineWidth int, linear, colorOnly bool, nonLinVarianceMult int) {  // nonLinVarianceMult is a multiplier for how variant the data is if linear is false
	a.lineWidth = lineWidth
	a.lineNum = int(math.Floor(float64(SCREEN_WIDTH/float32(a.lineWidth))))

	a.Active		= -1
	a.Active2		= -1
	a.PivotInd	= -1
	a.Shuffling = false
	a.CurrentText = ""
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
	var x = float32((i*a.lineWidth)+(a.lineWidth/2))
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

func (a *AnimArr) Draw() {
	var clr rl.Color
	for i := 0; i < a.lineNum; i++ {
		if i == a.Active {
			clr = rl.Green
		} else if i == a.Active2 {
			clr = rl.Red
		} else if i == a.PivotInd {
			clr = rl.Yellow
		//} else if a.Sorted && !a.colorOnly {   // Remove this to prevent the view going green when sorted.
		//	clr = rl.Lime
		} else {
			normal := uint8((a.Data[i]/a.maxValue)*255)  // Value normalised to 255
			//clr = rl.NewColor((normal/2)+127, (normal), (normal/3)+70, 255)  // Off yellow + coral
			//clr = rl.NewColor((normal/2)+127, (normal), (normal/3)+50, 255)  // Fire
			//clr = rl.NewColor(normal, normal, normal, 255)  // Grayscale
			//clr = rl.NewColor(normal, (normal/2)+127, normal/3, 255)  // Zesty (green --> yellow)
			//clr = rl.NewColor(normal, (normal/3), (normal/2)+127, 255)  // Vapourwave
			clr = rl.NewColor(normal, (normal/3), (normal/2)+127, 255)
		}
		a.drawLine(i, clr)
	}

	rl.DrawText(a.CurrentText, 10, 10, 30, rl.LightGray)
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

func (a *AnimArr) Reverse() {
	for i, j := 0, len(a.Data)-1; i < j; i, j = i+1, j-1 {
		a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
	}
}

func (a *AnimArr) swapElements(i1, i2 int) {
	a.Data[i1], a.Data[i2] = a.Data[i2], a.Data[i1]
}

func (a *AnimArr) Shuffle(times int, sleep bool) {
	a.Sorted = false
	a.Shuffling = true

	var max int = len(a.Data)
	for i := 0; i < times; i++ {
		a.CurrentText = fmt.Sprintf("Shuffling, round: %d", i+1)
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
	a.CurrentText = ""
}

func (a *AnimArr) changeDataBetween(start, end int, newSlice []float32, sleep bool) {
	for i := start; i < end; i++ {
		a.Active = i
		a.Data[i] = newSlice[i-start]
		if sleep { time.Sleep(CHANGE_SLEEP) }  // Sleep for half the time cause this bit should be quick
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
			time.Sleep(QS_SLEEP)
		}

		a.PivotInd = -1
		a.changeDataBetween(start, end, append(append(left, middle...), right...), true)
		a.QuickSort(start, start+len(left))
		a.QuickSort(start+len(left)+len(middle), start+len(left)+len(middle)+len(right))
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

func (a *AnimArr) InsertionSort() {
	for i := 1; i < len(a.Data); i++ {
		a.PivotInd = i
		for j := i; j > 0 && a.Data[j-1] > a.Data[j]; j-- {
			a.Active = j
			a.Active2 = j-1
			a.Data[j], a.Data[j-1] = a.Data[j-1], a.Data[j]
			time.Sleep(INST_SLEEP)
		}
	}
	a.Active = -1
	a.Active2 = -1
	a.PivotInd = -1
	a.Sorted = true
}


//func (a *AnimArr) MergeSort(start, end int) {

//}

func (a *AnimArr) DoQuickSort() {
	a.CurrentText = "Quick Sort"
	a.QuickSort(0, len(a.Data))
	fmt.Println("Finished sort.")
	a.Active = -1
	a.Sorted = true
	a.Sorting = false
	a.CurrentText = ""
}

func (a *AnimArr) DoSort(sort string) {
	a.Sorting = true
	a.Sorted = false
	if sort == "quick" {
		go a.DoQuickSort()
	} else if sort == "bogo" {
		a.CurrentText = "Bogo Sort"
		go func() {
			a.BogoSort()
			a.Sorting = false
			a.CurrentText = ""
		}()
	} else if sort == "bubble" {
		a.CurrentText = "Bubble Sort"
		go func() {
			a.BubbleSort()
			a.Sorting = false
			a.CurrentText = ""
		}()
	} else if sort == "insertion" {
		a.CurrentText = "Insertion Sort"
		go func() {
			a.InsertionSort()
			a.Sorting = false
			a.CurrentText = ""
		}()
	} else {
		panic("Invalid sort: "+sort)
	}
}


func (a *AnimArr) RunShowcase() {
	a.Showcase = true
	a.Sorting = true
	a.Sorted = false
	a.Shuffle(2, true)
	time.Sleep(time.Second)
	a.CurrentText = "Quick Sort"
	a.DoQuickSort()
	a.CurrentText = ""

	time.Sleep(time.Second)
	a.Sorting = true
	a.Sorted = false
	a.Shuffle(2, true)
	time.Sleep(time.Second)
	a.CurrentText = "Bubble Sort"
	a.BubbleSort()
	a.Sorting = false
	a.CurrentText = ""

	time.Sleep(time.Second)
	a.Sorting = true
	a.Sorted = false
	a.Shuffle(2, true)
	time.Sleep(time.Second)
	a.CurrentText = "Insertion Sort"
	a.InsertionSort()
	a.Sorting = false
	a.Showcase = false
	a.CurrentText = ""
}


func displayHelp() {
	rl.DrawRectangle(20, 60, helpW, helpH, rl.LightGray)

	v1, v2, v3, v4 := rl.NewVector2(20, 60), rl.NewVector2(float32(helpW)+20, 60), rl.NewVector2(20, float32(helpH)+60), rl.NewVector2(float32(helpW)+20, float32(helpH)+60) // All 4 verticies of square

	// BORDER
	rl.DrawLineEx(v1, v2, 2, rl.Red) // Top line
	rl.DrawLineEx(v3, v4, 2, rl.Red) // Bottom line
	rl.DrawLineEx(v1, v3, 2, rl.Red)  // Left line
	rl.DrawLineEx(v2, v4, 2, rl.Red)

	// Information
	var (
		g1x int32 = 30
		g2x int32 = helpW/2  // Guidelines
		fontSize int32 = 20
	)
	rl.DrawText("Key", g1x, 70, fontSize+2, rl.Black)     // Headers
	rl.DrawText("Function", g2x, 70, fontSize+2, rl.Black)

	rl.DrawText("1,2...9", g1x, 110, fontSize, rl.Black)
	rl.DrawText("Do various sorts.", g2x, 110, fontSize, rl.Black)

	rl.DrawText("'s'", g1x, 140, fontSize, rl.Black)
	rl.DrawText("Shuffle the data.", g2x, 140, fontSize, rl.Black)

	rl.DrawText("'r'", g1x, 170, fontSize, rl.Black)
	rl.DrawText("Reverse the data.", g2x, 170, fontSize, rl.Black)

	rl.DrawText("'p'", g1x, 200, fontSize, rl.Black)
	rl.DrawText("Run the showcase.", g2x, 200, fontSize, rl.Black)

	rl.DrawText("'l'", g1x, 230, fontSize, rl.Black)
	rl.DrawText("Sort the data instantly.", g2x, 230, fontSize, rl.Black)
}


func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Sort Visualiser")
	rl.SetTargetFPS(144)

	anim := AnimArr{}
	anim.Init(2, true, false, 2)  // Input line thickness, if it is linear, and if it is color only here

	helpOpen := false

	for !rl.WindowShouldClose() {
		if !anim.Sorting && !anim.Shuffling && !anim.Showcase {
			if rl.IsKeyPressed(rl.KeyS) {
				go anim.Shuffle(2, true)

			} else if rl.IsKeyPressed(rl.KeyOne) {
				anim.DoSort("quick")
			} else if rl.IsKeyPressed(rl.KeyTwo) {
				anim.DoSort("bubble")
			} else if rl.IsKeyPressed(rl.KeyThree) {
				anim.DoSort("insertion")
			} else if rl.IsKeyPressed(rl.KeyNine) {
				anim.DoSort("bogo")
			} else if rl.IsKeyPressed(rl.KeyL) {
				anim.Data = regularQuickSort(anim.Data)
				anim.Sorted = true
			} else if rl.IsKeyPressed(rl.KeyR) {
				go func() {
					anim.Reverse()
				}()
			} else if rl.IsKeyPressed(rl.KeyP) {
				go anim.RunShowcase()
			}
		}

		if rl.IsKeyPressed(rl.KeyH) { // Open H
			helpOpen = !helpOpen
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		anim.Draw()

		if helpOpen {
			displayHelp()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
