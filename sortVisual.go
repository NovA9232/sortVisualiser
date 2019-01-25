package main

import (
	"math"
	"math/rand"
	"time"
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_WIDTH  float32 = 1200
	SCREEN_HEIGHT float32 = 450

	SORT_SLEEP_TIME = time.Millisecond
)

type AnimArr struct {
	Data []int
	lineNum int
	lineWidth int
	Active int		// Index of current element being operated on.
	Sorted bool
}

func (a *AnimArr) Init(lineWidth int) {
	a.lineWidth = lineWidth
	a.Sorted = true
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
		} else {
			a.drawLine(i, rl.Violet)
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

func (a *AnimArr) changeDataBetween(start, end int, newSlice []int) {
	for i := start; i < end; i++ {
		a.Data[i] = newSlice[i-start]
	}
}

func (a *AnimArr) QuickSort(start, end int) {   // Start and end of part of array to sort
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
			time.Sleep(SORT_SLEEP_TIME)
		}
		a.changeDataBetween(start, end, append(append(left, middle...), right...))
		a.QuickSort(start, start+len(left))
		a.QuickSort(start+len(left)+len(middle), start+len(left)+len(middle)+len(right))
	}
}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Egg")
	rl.SetTargetFPS(60)

	anim := AnimArr{}
	anim.Init(1)
	go func() {
		anim.QuickSort(0, len(anim.Data))
		fmt.Println("Finished quick sort.")
		anim.Active = 0
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		//rl.DrawText("Egg", 190, 200, 20, rl.LightGray)
		//rl.DrawLine(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, rl.Green)
		//rl.DrawLineEx(rl.NewVector2(0, 0), rl.NewVector2(float32(SCREEN_WIDTH), float32(SCREEN_HEIGHT)), 4, rl.Green)
		anim.Draw(0)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
