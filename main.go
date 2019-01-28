package main

import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"

	"AnimatedArr"
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

	rl.DrawText("'c'", g1x, 260, fontSize, rl.Black)
	rl.DrawText("Toggle colour only mode.", g2x, 260, fontSize, rl.Black)
}


func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "Sort Visualiser")
	rl.SetTargetFPS(60)

	anim := AnimatedArr.AnimArr{}
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
				anim.Data = AnimatedArr.RegularQuickSort(anim.Data)
				anim.Sorted = true
			} else if rl.IsKeyPressed(rl.KeyR) {
				go func() {
					anim.Reverse()
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
			displayHelp()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
