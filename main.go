package main

import (
	"github.com/gen2brain/raylib-go/raylib"

	"animatedArr"
	"helpMenu"
)

const (
	WIN_SIZE_CHECK_DELAY = 0.1
	DEFAULT_LINE_WIDTH = 10
	LINE_WIDTH_INCREMENT = 2
)

var (
	screenWidth  int = 1600
	screenHeight int = 800

	currLineWidth int = DEFAULT_LINE_WIDTH

	//violet rl.Color = NewColor(61, 38, 69, 255)
	//raspberry rl.Color = NewColor(131, 33, 97, 255)
	//coral rl.Color = NewColor(218, 65, 103, 255)
)

func checkScreenSizeChange(a *animatedArr.AnimArr) {
	w, h := rl.GetScreenWidth(), rl.GetScreenHeight()
	if w != screenWidth || h != screenHeight {
		screenWidth = w
		screenHeight = h
		println("Window changed size")
		a.Init(float32(screenWidth), float32(screenHeight), currLineWidth, a.Linear, a.ColorOnly, 2)
	}
}

func changeLineWidth(a *animatedArr.AnimArr, amount int) {
	newWidth := currLineWidth + amount
	if newWidth > 0 && newWidth < screenWidth {
		currLineWidth = newWidth
		a.Init(float32(screenWidth), float32(screenHeight), currLineWidth, a.Linear, a.ColorOnly, 2)
	}
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Sort Visualiser")
	rl.SetTargetFPS(60)

	anim := &animatedArr.AnimArr{}
	anim.Init(float32(screenWidth), float32(screenHeight), DEFAULT_LINE_WIDTH, true, false, 2)  // Input line thickness, if it is linear, and if it is color only here

	var checkTimer float32 = 0

	helpM := helpMenu.NewHelpMenu()
	sortKeybindM := helpMenu.NewSortsKeyBindMenu()

	for !rl.WindowShouldClose() {
		if !anim.Sorting && !anim.Shuffling && !anim.Showcase {
			if rl.IsKeyPressed(rl.KeyEqual) || rl.IsKeyPressed(rl.KeyKpAdd) { // Treat "=" like plus sign
				changeLineWidth(anim, LINE_WIDTH_INCREMENT)
			} else if rl.IsKeyPressed(rl.KeyMinus) || rl.IsKeyPressed(rl.KeyKpSubtract) {
				changeLineWidth(anim, -LINE_WIDTH_INCREMENT)
			}

			checkTimer += rl.GetFrameTime()
			if checkTimer >= WIN_SIZE_CHECK_DELAY {
				checkScreenSizeChange(anim)
				checkTimer = 0
			}
		}

		anim.Update()

		if rl.IsKeyPressed(rl.KeyH) { // Open H
			helpM.Open = !helpM.Open
			if sortKeybindM.Open {
				sortKeybindM.Open = false
			}
		}

		if rl.IsKeyPressed(rl.KeyK) { // Open sort binds help menu
			sortKeybindM.Open = !sortKeybindM.Open
			if helpM.Open {
				helpM.Open = false
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		anim.Draw()

		if helpM.Open {
			helpM.Draw()
		}

		if sortKeybindM.Open {
			sortKeybindM.Draw()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
