package helpMenu

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
	SORT_HELP_W int32 = 533
	SORT_HELP_H int32 = 320
)

type SortKeybindMenu struct {
	HelpMenu
}

func NewSortsKeyBindMenu() *SortKeybindMenu {
	s := &SortKeybindMenu {
		HelpMenu: *NewHelpMenu(),
	}
	s.genTextPositions()

	return s
}

func (h *SortKeybindMenu) genTextPositions() {
	h.textPositions = []int32{}
	for i := int32(0); i < 9; i++ {
		h.textPositions = append(h.textPositions, int32(h.pos.Y+10) + (i*30))
	}
}

func (h *SortKeybindMenu) Draw() {
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

	rl.DrawText("1", h.g1x, h.textPositions[1], fontSize, rl.Black)
	rl.DrawText("Quick Sort.", h.g2x, h.textPositions[1], fontSize, rl.Black)

	rl.DrawText("2", h.g1x, h.textPositions[2], fontSize, rl.Black)
	rl.DrawText("Bubble Sort.", h.g2x, h.textPositions[2], fontSize, rl.Black)

	rl.DrawText("3", h.g1x, h.textPositions[3], fontSize, rl.Black)
	rl.DrawText("Insertion Sort.", h.g2x, h.textPositions[3], fontSize, rl.Black)

	rl.DrawText("4", h.g1x, h.textPositions[4], fontSize, rl.Black)
	rl.DrawText("Shell Sort.", h.g2x, h.textPositions[4], fontSize, rl.Black)

	rl.DrawText("5", h.g1x, h.textPositions[5], fontSize, rl.Black)
	rl.DrawText("Merge Sort.", h.g2x, h.textPositions[5], fontSize, rl.Black)

	rl.DrawText("6", h.g1x, h.textPositions[6], fontSize, rl.Black)
	rl.DrawText("Cocktail Shaker Sort.", h.g2x, h.textPositions[6], fontSize, rl.Black)

	rl.DrawText("7", h.g1x, h.textPositions[7], fontSize, rl.Black)
	rl.DrawText("Comb Sort.", h.g2x, h.textPositions[7], fontSize, rl.Black)

	rl.DrawText("9", h.g1x, h.textPositions[8], fontSize, rl.Black)
	rl.DrawText("Bogo Sort.", h.g2x, h.textPositions[8], fontSize, rl.Black)

}
