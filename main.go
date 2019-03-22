package main

import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"

	"animatedArr"
	"helpMenu"
)

const (
	helpW int32 = 533
	WIN_SIZE_CHECK_DELAY = 0.1
	DEFAULT_LINE_WIDTH = 4
)

var (
	screenWidth  int = 1600
	screenHeight int = 800
	helpH int32 = int32(math.Floor(float64(screenHeight)/3))

	//violet rl.Color = NewColor(61, 38, 69, 255)
	//raspberry rl.Color = NewColor(131, 33, 97, 255)
	//coral rl.Color = NewColor(218, 65, 103, 255)
)

const (
  maxSamples = 22050
)

func checkScreenSizeChange(a *animatedArr.AnimArr) {
	w, h := rl.GetScreenWidth(), rl.GetScreenHeight()
	if w != screenWidth || h != screenHeight {
		screenWidth = w
		screenHeight = h
		println("Window changed size")
		a.Init(float32(screenWidth), float32(screenHeight), DEFAULT_LINE_WIDTH, true, false, 2)  // Input line thickness, if it is linear, and if it is color only here
	}
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Sort Visualiser")
	rl.SetTargetFPS(60)

	anim := &animatedArr.AnimArr{}
	anim.Init(float32(screenWidth), float32(screenHeight), 2, true, false, 2)  // Input line thickness, if it is linear, and if it is color only here

	var checkTimer float32 = 0

	helpOpen := false

	for !rl.WindowShouldClose() {
		checkTimer += rl.GetFrameTime()
		if checkTimer >= WIN_SIZE_CHECK_DELAY {
			checkScreenSizeChange(anim)
			checkTimer = 0
		}

		anim.Update()

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
