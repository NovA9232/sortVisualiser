package main

import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"

	"animatedArr"
	"helpMenu"
)

var (
	SCREEN_WIDTH  float32 = 1600
	SCREEN_HEIGHT float32 = 800

	helpW int32 = 533
	helpH int32 = int32(math.Floor(float64(SCREEN_HEIGHT)/3))

	//violet rl.Color = NewColor(61, 38, 69, 255)
	//raspberry rl.Color = NewColor(131, 33, 97, 255)
	//coral rl.Color = NewColor(218, 65, 103, 255)
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Sort Visualiser")
	rl.SetTargetFPS(60)

	anim := animatedArr.AnimArr{}
	anim.Init(SCREEN_WIDTH, SCREEN_HEIGHT, 2, true, false, 2)  // Input line thickness, if it is linear, and if it is color only here

	helpOpen := false

	for !rl.WindowShouldClose() {
		if !anim.Sorting && !anim.Shuffling && !anim.Showcase {
			if rl.IsKeyPressed(rl.KeyS) {
				anim.ArrayAccesses = 0
				anim.Comparisons = 0
				go func() {
					anim.Shuffle(2, true, false)
					anim.CurrentText = ""
				}()
			} else if rl.IsKeyPressed(rl.KeyOne) {
				anim.DoSort("quick")
			} else if rl.IsKeyPressed(rl.KeyTwo) {
				anim.DoSort("bubble")
			} else if rl.IsKeyPressed(rl.KeyThree) {
				anim.DoSort("insertion")
			} else if rl.IsKeyPressed(rl.KeyFour) {
				anim.DoSort("shell")
			} else if rl.IsKeyPressed(rl.KeyFive) {
				anim.DoSort("merge")
			} else if rl.IsKeyPressed(rl.KeyNine) {
				anim.DoSort("bogo")
			} else if rl.IsKeyPressed(rl.KeyL) {
				anim.Data = animatedArr.RegularQuickSort(anim.Data)
				anim.Sorted = true
			} else if rl.IsKeyPressed(rl.KeyR) {
				go func() {
					anim.Reverse(anim.Data)
				}()
			} else if rl.IsKeyPressed(rl.KeyP) {
				go anim.RunShowcase()
			}
		}

		if rl.IsKeyPressed(rl.KeyC) {
			anim.ColorOnly = !anim.ColorOnly
		}

		if rl.IsKeyPressed(rl.KeyH) { // Open H
			helpOpen = !helpOpen
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		anim.Draw()

		if helpOpen {
			helpMenu.DisplayHelp(helpW, helpH)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
