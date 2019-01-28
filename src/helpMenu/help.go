package helpMenu

import (
  "github.com/gen2brain/raylib-go/raylib"
)

func DisplayHelp(width, height int32) {
	rl.DrawRectangle(20, 60, width, height, rl.LightGray)

	v1, v2, v3, v4 := rl.NewVector2(20, 60),
                    rl.NewVector2(float32(width)+20, 60),
                    rl.NewVector2(20, float32(height)+60),
                    rl.NewVector2(float32(width)+20, float32(height)+60) // All 4 verticies of square

	// BORDER
	rl.DrawLineEx(v1, v2, 2, rl.Red) // Top line
	rl.DrawLineEx(v3, v4, 2, rl.Red) // Bottom line
	rl.DrawLineEx(v1, v3, 2, rl.Red)  // Left line
	rl.DrawLineEx(v2, v4, 2, rl.Red)

	// Information
	var (
		g1x int32 = 30
		g2x int32 = width/2  // Guidelines
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
