package animatedArr

import (
  "math"
  "fmt"
  "time"
  "github.com/gen2brain/raylib-go/raylib"
)

var (
  // Base speeds (usually~ in time per comparison/change)
  QS_SLEEP time.Duration
  CHANGE_SLEEP time.Duration
  MS_SLEEP time.Duration
  BBL_SLEEP time.Duration
  INST_SLEEP time.Duration
  SHL_SLEEP time.Duration
  CCT_SLEEP time.Duration

  SHUFFLE_SLEEP = time.Microsecond * 500
)

type AnimArr struct {
	Data					[]float32
	sortedData		[]int
	lineNum				int
	lineWidth			int
	Active				int		// Index of current element being operated on.
	Active2				int   // Secondary active, for swapping elements.
	PivotInd			int   // For highlighting pivot when doing quickSort.
	nonLinearMult int
	ArrayAccesses int
	Comparisons		int
  W             float32
  H             float32
	maxValue			float32
	CurrentText   string
	Sorted				bool
	Sorting				bool
	Shuffling			bool
	linear				bool
	ColorOnly			bool // Do not show height if true
	Showcase			bool  // If showcase is running
}

func (a *AnimArr) Init(width, height float32, lineWidth int, linear, colorOnly bool, nonLinVarianceMult int) {  // nonLinVarianceMult is a multiplier for how variant the data is if linear is false
  a.W, a.H = width, height
	a.lineWidth = lineWidth
	a.lineNum = int(math.Floor(float64(a.W/float32(a.lineWidth))))

	a.Active		= -1
	a.Active2		= -1
	a.PivotInd	= -1
	a.Shuffling = false
	a.CurrentText = ""
	a.linear		= linear
	a.nonLinearMult = nonLinVarianceMult
	a.ColorOnly = colorOnly
	a.Sorted		= a.linear
	a.Sorting   = false

	oNlogN := time.Duration(float64(a.lineNum) * math.Log(float64(a.lineNum)))
	oNSqrd := time.Duration(math.Pow(float64(a.lineNum), 2))

  QS_SLEEP = time.Second * 8 / oNlogN          // O(n log n)
  CHANGE_SLEEP = QS_SLEEP
  MS_SLEEP = time.Second * 2 * 10 / oNlogN  // O(n log n)
  BBL_SLEEP = time.Second * 2 * 6 / oNSqrd   // O(n^2)
  INST_SLEEP = time.Second * 2 * 6 / oNSqrd	 // O(n^2)
  SHL_SLEEP = time.Second * 2 * 10 / time.Duration(math.Pow(float64(a.lineNum), 1.5))   // O(n^(3/2)) 
  CCT_SLEEP = time.Second * 2 * 10 / oNSqrd  // O(n^2)

	if a.linear {
		a.Data = a.GenerateLinear(0, a.H, a.H/float32(a.lineNum))
	} else {
		a.Data = a.Generate(a.lineNum, a.lineNum*a.nonLinearMult)
	}
}

func (a *AnimArr) getLineY(val float32) float32 {   // Lower case incase I want to have this as a package.
	return a.H-((float32(val)/float32(a.lineNum*a.nonLinearMult))*a.H)
}

func (a *AnimArr) drawLine(i int, colour rl.Color) {  // English spelling
	var x = float32((i*a.lineWidth)+(a.lineWidth/2))
	var y float32
	if a.ColorOnly {
		y = 0
	} else if a.linear {
		y = a.H-a.Data[i]
	} else {
		y = a.getLineY(a.Data[i])
	}
	rl.DrawLineEx(rl.NewVector2(x, a.H), rl.NewVector2(x, y), float32(a.lineWidth), colour)
}

func (a *AnimArr) Draw() {
	var clr rl.Color
	for i := 0; i < len(a.Data); i++ {
		if i == a.Active {
			clr = rl.Green
		} else if i == a.Active2 {
			clr = rl.Red
		} else if i == a.PivotInd {
			clr = rl.Yellow
		//} else if a.Sorted && !a.ColorOnly {   // Remove this to prevent the view going green when sorted.
		//	clr = rl.Lime
		} else {
			normal := uint8((a.Data[i]/a.maxValue)*255)  // Value normalised to 255
			//clr = rl.NewColor((normal/2)+127, (normal), (normal/3)+70, 255)  // Off yellow + coral
			//clr = rl.NewColor((normal/2)+127, (normal), (normal/3)+50, 255)  // Fire
			//clr = rl.NewColor(normal, normal, normal, 255)  // Grayscale
			//clr = rl.NewColor(normal, (normal/2)+127, normal/3, 255)  // Zesty (green --> yellow)
			clr = rl.NewColor(normal, (normal/3), (normal/2)+127, 255)  // Twilight/Vapourwave
      //clr = rl.NewColor(128-(normal/2), 191-(normal/4), normal, 255)  // Sea
      //clr = rl.NewColor(((normal)/3)+85, 128-(normal/2), 170-(normal/3), 255)  // Soft Vapourwave
		}
		a.drawLine(i, clr)
	}

	rl.DrawText(a.CurrentText, 10, 10, 30, rl.LightGray)

	if a.ArrayAccesses+a.Comparisons > 0 {
		rl.DrawText(fmt.Sprintf("Total length of array: %d", len(a.Data)), 10, 80, 20, rl.LightGray)
		if a.ArrayAccesses > 0 {
			rl.DrawText(fmt.Sprintf("Array accesses: %d", a.ArrayAccesses), 10, 40, 20, rl.LightGray)
		}
		if a.Comparisons > 0 {
			rl.DrawText(fmt.Sprintf("Comparisons: %d", a.Comparisons), 10, 60, 20, rl.LightGray)
		}
	}
}

func (a *AnimArr) Update() {
	if rl.IsKeyPressed(rl.KeyC) {
		a.ColorOnly = !a.ColorOnly
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		a.Sorted = true
		println("Stopping...")
	}

	if !a.Sorting && !a.Shuffling && !a.Showcase {
		if rl.IsKeyPressed(rl.KeyS) {
			a.ArrayAccesses = 0
			a.Comparisons = 0
			go func() {
				a.Shuffle(2, true, false)
				a.CurrentText = ""
			}()
		} else if rl.IsKeyPressed(rl.KeyOne) {
			a.DoSort("quick")
		} else if rl.IsKeyPressed(rl.KeyTwo) {
			a.DoSort("bubble")
		} else if rl.IsKeyPressed(rl.KeyThree) {
			a.DoSort("insertion")
		} else if rl.IsKeyPressed(rl.KeyFour) {
			a.DoSort("shell")
		} else if rl.IsKeyPressed(rl.KeyFive) {
			a.DoSort("merge")
		} else if rl.IsKeyPressed(rl.KeySix) {
			a.DoSort("shaker")
		} else if rl.IsKeyPressed(rl.KeyNine) {
			a.DoSort("bogo")
		} else if rl.IsKeyPressed(rl.KeyL) {
			a.Data = RegularQuickSort(a.Data)
			a.Sorted = true
		} else if rl.IsKeyPressed(rl.KeyR) {
			go func() {
				a.Reverse(a.Data)
			}()
		} else if rl.IsKeyPressed(rl.KeyP) {
			go a.RunShowcase()
		}
	}
}

func (a *AnimArr) resetVals() {
	a.Sorting = false
	a.Active = -1
	a.Active2 = -1
	a.PivotInd = -1
	a.Sorted = true
}

func (a *AnimArr) DoSort(sort string) {
	a.Sorting = true
	a.Sorted = false
	a.ArrayAccesses = 0
	a.Comparisons = 0
	if sort == "quick" {
		a.CurrentText = "Quick Sort"
		go func() {
			a.QuickSort(0, len(a.Data))
			a.resetVals()
		}()
	} else if sort == "bogo" {
		a.CurrentText = "Bogo Sort"
		go func() {
			a.BogoSort()
			a.resetVals()
		}()
	} else if sort == "bubble" {
		a.CurrentText = "Bubble Sort"
		go func() {
			a.BubbleSort()
			a.resetVals()
		}()
	} else if sort == "insertion" {
		a.CurrentText = "Insertion Sort"
		go func() {
			a.InsertionSort()
			a.resetVals()
		}()
	} else if sort == "shell" {
		a.CurrentText = "Shell Sort"
		go func() {
			a.ShellSort()
			a.resetVals()
		}()
	} else if sort == "merge" {
		a.CurrentText = "Merge Sort"
		go func() {
			a.MergeSort(0, len(a.Data))
			a.resetVals()
		}()
  } else if sort == "shaker" {
    a.CurrentText = "Cocktail Shaker Sort"
    go func() {
      a.CocktailShakerSort()
      a.resetVals()
    }()
	} else {
		panic("Invalid sort: "+sort)
	}
}
