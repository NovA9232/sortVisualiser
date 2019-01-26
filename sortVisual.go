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
	SCREEN_HEIGHT float32 = 1000

	//violet rl.Color = NewColor(61, 38, 69, 255)
	//raspberry rl.Color = NewColor(131, 33, 97, 255)
	//coral rl.Color = NewColor(218, 65, 103, 255)

	SORT_SLEEP = time.Millisecond
	SHUFFLE_SLEEP = time.Millisecond
)

type AnimArr struct {
	Data			 []float32
	sortedData []int
	lineNum			 int
	lineWidth		 int
	Active			 int		// Index of current element being operated on.
	Active2			 int   // Secondary active, for swapping elements.
	PivotInd		 int   // For highlighting pivot when doing quickSort.
	Sorted			bool
	Shuffling		bool
	linear			bool
}

func (a *AnimArr) Init(lineWidth int) {
	a.lineWidth = lineWidth
	a.lineNum = int(math.Floor(float64(SCREEN_WIDTH/float32(a.lineWidth))))
	a.Active		= -1
	a.Active2		= -1
	a.PivotInd	= -1
	a.Sorted		= false
	a.Shuffling = false
	a.linear		= true
	//a.Data			= a.Generate(a.lineNum, a.lineNum*2)
	a.Data			= a.GenerateLinear(0, SCREEN_HEIGHT, SCREEN_HEIGHT/float32(a.lineNum))

	//a.sortedData = make([]int, len(a.Data))
	//copy(a.sortedData, a.Data)
	//a.sortedData = regularQuickSort(a.sortedData)
}

func regularQuickSort(arr []int) []int {  // Just a quick sort to sort the array for comparison (if needed)
	if len(arr) < 2 { return arr }
	var left []int
	var middle []int
	var right []int
	var pivot int = arr[len(arr)/2]

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
	return SCREEN_HEIGHT-((float32(val)/float32(a.lineNum*2))*SCREEN_HEIGHT)
}

func (a *AnimArr) drawLine(i int, colour rl.Color) {  // English spelling
	var x = float32(i*a.lineWidth)
	var y float32
	if a.linear {
		y = SCREEN_HEIGHT-a.Data[i]
	} else {
		y = a.getLineY(a.Data[i])
	}
	rl.DrawLineEx(rl.NewVector2(x, SCREEN_HEIGHT), rl.NewVector2(x, y), float32(a.lineWidth), colour)
}

func (a *AnimArr) Draw(dt float32) {
	for i := 0; i < a.lineNum; i++ {
		if i == a.Active {
			a.drawLine(i, rl.Green)
		} else if i == a.Active2 {
			a.drawLine(i, rl.Red)
		} else if i == a.PivotInd {
			a.drawLine(i, rl.Yellow)
		} else if a.Sorted {
			a.drawLine(i, rl.Lime)
		} else {
			a.drawLine(i, rl.DarkBlue)
		}
	}
}

func (a *AnimArr) Generate(num, max int) []float32 {	// Generates array
	var out []float32
	for i := 0; i < num; i++ {
		out = append(out, float32(rand.Intn(max)))
	}
	return out
}

func (a *AnimArr) GenerateLinear(start, finish, jump float32) []float32 {
	fmt.Println("Oof:", start, finish, jump)
	var out []float32
	for i := start; i < finish; i += jump {
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
	fmt.Println("Shuffling...")
	var max int = len(a.Data)
	for i := 0; i < times; i++ {
		for j := 0; j < len(a.Data); j++ {
			a.Active	= rand.Intn(max)
			a.Active2 = rand.Intn(max)
			a.swapElements(a.Active, a.Active2)
			if sleep { time.Sleep(SHUFFLE_SLEEP) }
		}
		fmt.Println("Done:", i, "lots.")
	}
	fmt.Println("Finished shuffling.")
	a.Shuffling = false
	a.Active	= -1
	a.Active2 = -1
}

func (a *AnimArr) changeDataBetween(start, end int, newSlice []float32, sleep bool) {
	for i := start; i < end; i++ {
		a.Active = i
		a.Data[i] = newSlice[i-start]
		if sleep { time.Sleep(SORT_SLEEP) }
	}
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

//func (a *AnimArr) MergeSort(start, end int) {

//}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Egg")
	rl.SetTargetFPS(60)

	anim := AnimArr{}
	anim.Init(1)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(82) && !anim.Shuffling {  // When 'r' is pressed, shuffle the array.
			go func() {
				anim.Shuffle(4, true)
			}()
		}

		if rl.IsKeyPressed(83) {  // When 's' is pressed, sort the array.
			go func() {
				anim.Sorted = false
				anim.QuickSort(0, len(anim.Data))
				fmt.Println("Finished sort.")
				anim.Active = -1
				anim.Sorted = true
			}()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		anim.Draw(0)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
