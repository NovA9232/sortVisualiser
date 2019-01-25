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

	SORT_SLEEP = time.Millisecond * 10
	SHUFFLE_SLEEP = time.Millisecond/2
)

type AnimArr struct {
	Data []int
	lineNum int
	lineWidth int
	Active int		// Index of current element being operated on.
	Active2 int   // Secondary active, for swapping elements.
	Sorted bool
	Shuffling bool
}

func (a *AnimArr) Init(lineWidth int) {
	a.lineWidth = lineWidth
	a.Sorted		= false
	a.Shuffling = false
	a.lineNum = int(math.Floor(float64(SCREEN_WIDTH/float32(a.lineWidth))))
	a.Data = a.Generate(a.lineNum, a.lineNum*2)

}

func (a *AnimArr) getLineY(val int) float32 {   // Lower case incase I want to have this as a package.
	return SCREEN_HEIGHT-((float32(val)/float32(a.lineNum*2))*SCREEN_HEIGHT)
}

func (a *AnimArr) drawLine(i int, colour rl.Color) {  // English spelling
	x := float32(i*a.lineWidth)
	rl.DrawLineEx(rl.NewVector2(x, SCREEN_HEIGHT), rl.NewVector2(x, a.getLineY(a.Data[i])), float32(a.lineWidth), colour)
}

func (a *AnimArr) Draw(dt float32) {
	for i := 0; i < a.lineNum; i++ {
		if i == a.Active {
			a.drawLine(i, rl.Green)
		} else if i == a.Active2 {
			a.drawLine(i, rl.Red)
		} else if a.Sorted {
			a.drawLine(i, rl.Lime)
		} else {
			a.drawLine(i, rl.DarkBlue)
		}
	}
}

func (a *AnimArr) Generate(num, max int) []int {	// Generates array
	var out []int
	for i := 0; i < num; i++ {
		out = append(out, rand.Intn(max))
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

func (a *AnimArr) changeDataBetween(start, end int, newSlice []int, sleep bool) {
	for i := start; i < end; i++ {
		a.Active = i
		a.Data[i] = newSlice[i-start]
		if sleep { time.Sleep(SORT_SLEEP) }
	}
}

func (a *AnimArr) QuickSort(start, end int, sleep bool) {   // Start and end of part of array to sort
	if end-start > 1 {
		var left []int
		var middle []int
		var right []int
		var pivot int = a.Data[int(math.Floor(float64(start+end)/2))]

		for i := start; i < end; i++ {
			a.Active = i
			if a.Data[i] < pivot {
				left = append(left, a.Data[i])
			} else if a.Data[i] > pivot {
				right = append(right, a.Data[i])
			} else {
				middle = append(middle, a.Data[i])
			}
			if sleep { time.Sleep(SORT_SLEEP) }
		}
		a.changeDataBetween(start, end, append(append(left, middle...), right...), true)
		a.QuickSort(start, start+len(left), sleep)
		a.QuickSort(start+len(left)+len(middle), start+len(left)+len(middle)+len(right), sleep)
	}
}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Egg")
	rl.SetTargetFPS(144)

	anim := AnimArr{}
	anim.Init(10)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(82) && !anim.Shuffling {  // When 'r' is pressed, shuffle the array.
			go func() {
				anim.Shuffle(4, true)
			}()
		}

		if rl.IsKeyPressed(83) {  // When 's' is pressed, sort the array.
			go func() {
				anim.Sorted = false
				anim.QuickSort(0, len(anim.Data), true)
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
