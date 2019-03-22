package helpMenu

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
	HELP_W int32 = 533
	HELP_H int32 = 320
	HELP_X = 10
	HELP_Y = 110
)

type HelpMenu struct {
	pos rl.Vector2
	textPositions []int32
	g1x int32  // guidelines
	g2x int32
	Open bool
}

func NewHelpMenu() *HelpMenu {
	h := &HelpMenu {
		pos: rl.NewVector2(HELP_X, HELP_Y),
		g1x: HELP_X+10,
		g2x: HELP_X + (HELP_W/2),  // Guidelines
		Open: false,
	}
	h.genTextPositions()
	return h
}

func (h *HelpMenu) genTextPositions() {
	h.textPositions = []int32{}
	for i := int32(0); i < 10; i++ {
		h.textPositions = append(h.textPositions, int32(h.pos.Y+10) + (i*30))
	}
}

func (h *HelpMenu) Draw() {
	rl.DrawRectangle(int32(h.pos.X), int32(h.pos.Y), HELP_W, HELP_H, rl.LightGray)

	v1, v2, v3, v4 := rl.NewVector2(h.pos.X, h.pos.Y),
                    rl.NewVector2(float32(HELP_W)+h.pos.X, h.pos.Y),
                    rl.NewVector2(h.pos.X, float32(HELP_H)+h.pos.Y),
                    rl.NewVector2(float32(HELP_W)+h.pos.X, float32(HELP_H)+h.pos.Y) // All 4 verticies of square

	// BORDER
	rl.DrawLineEx(v1, v2, 2, rl.Red) // Top line
	rl.DrawLineEx(v3, v4, 2, rl.Red) // Bottom line
	rl.DrawLineEx(v1, v3, 2, rl.Red)  // Left line
	rl.DrawLineEx(v2, v4, 2, rl.Red)

	// Information
	var (
		fontSize int32 = 20
	)
	rl.DrawText("Key", h.g1x, h.textPositions[0], fontSize+2, rl.Black)     // Headers
	rl.DrawText("Function", h.g2x, h.textPositions[0], fontSize+2, rl.Black)

	rl.DrawText("1,2...9", h.g1x, h.textPositions[1], fontSize, rl.Black)
	rl.DrawText("Do various sorts.", h.g2x, h.textPositions[1], fontSize, rl.Black)

	rl.DrawText("k", h.g1x, h.textPositions[2], fontSize, rl.Black)
	rl.DrawText("View sort keybinds (1..9).", h.g2x, h.textPositions[2], fontSize, rl.Black)

	rl.DrawText("'s'", h.g1x, h.textPositions[3], fontSize, rl.Black)
	rl.DrawText("Shuffle the data.", h.g2x, h.textPositions[3], fontSize, rl.Black)

	rl.DrawText("'r'", h.g1x, h.textPositions[4], fontSize, rl.Black)
	rl.DrawText("Reverse the data.", h.g2x, h.textPositions[4], fontSize, rl.Black)

	rl.DrawText("'p'", h.g1x, h.textPositions[5], fontSize, rl.Black)
	rl.DrawText("Run the showcase.", h.g2x, h.textPositions[5], fontSize, rl.Black)

	rl.DrawText("'l'", h.g1x, h.textPositions[6], fontSize, rl.Black)
	rl.DrawText("Sort the data instantly.", h.g2x, h.textPositions[6], fontSize, rl.Black)

	rl.DrawText("'c'", h.g1x, h.textPositions[7], fontSize, rl.Black)
	rl.DrawText("Toggle colour only mode.", h.g2x, h.textPositions[7], fontSize, rl.Black)

	rl.DrawText("'q'", h.g1x, h.textPositions[8], fontSize, rl.Black)
	rl.DrawText("Stop current sort ASAP.", h.g2x, h.textPositions[8], fontSize, rl.Black)

	rl.DrawText("'+' / '-'", h.g1x, h.textPositions[9], fontSize, rl.Black)
	rl.DrawText("Change line width.", h.g2x, h.textPositions[9], fontSize, rl.Black)
}
