package animatedArr

import (
  "math"
  "fmt"
  "time"
  "github.com/gen2brain/raylib-go/raylib"
)

const NON_LINEAR_VARIANCE = 10

var (
  // Base speeds (usually~ in time per comparison/change)
  QS_SLEEP time.Duration
  CHANGE_SLEEP time.Duration
  MS_SLEEP time.Duration
  BBL_SLEEP time.Duration
  INST_SLEEP time.Duration
  SHL_SLEEP time.Duration
  CCT_SLEEP time.Duration
  COMB_SLEEP time.Duration

  SHUFFLE_SLEEP = time.Microsecond * 500

  ScreenWidth *int
  ScreenHeight *int
)

type AnimArr struct {
	Data					[]float32
	lineNum				int
	LineWidth			int
	Active				int		// Index of current element being operated on.
	Active2				int   // Secondary active, for swapping elements.
	PivotInd				int   // For highlighting pivot when doing quickSort.
	nonLinearMult		int
	ArrayAccesses		int
	Comparisons			int
	W             float32
	H             float32
	beepSound	 rl.Sound
	lastPitchShift	float32
	lastActive			int
	maxValue		  float32
	CurrentText		string
	Sorted			  bool
	Sorting			  bool
	Shuffling		  bool
	Linear			  bool
	ColorOnly		  bool // Do not show height if true
	Dots				  bool // Draw with dots
	Showcase			  bool  // If showcase is running
	SleepMultiplier float32
	totalSleepTime  float64
	totalTime		 float64
}

func (a *AnimArr) Init(width, height float32, LineWidth int, linear, colorOnly, dots bool, nonLinVarianceMult int) {  // nonLinVarianceMult is a multiplier for how variant the data is if linear is false
	a.W, a.H = width, height
	a.LineWidth = LineWidth
	a.lineNum = int(math.Floor(float64(a.W/float32(a.LineWidth))))

	//a.beepSound = rl.LoadSound("src/sounds/boop.wav")
	a.lastPitchShift = 1
	a.lastActive = -1

	a.Active		= -1
	a.Active2		= -1
	a.PivotInd	= -1
	a.Shuffling = false
	a.CurrentText = ""
	a.Linear		= linear
	a.nonLinearMult = nonLinVarianceMult
	a.ColorOnly = colorOnly
	a.Dots = dots
	a.Sorted		= a.Linear
	a.Sorting   = false
	a.SleepMultiplier = 10
	a.totalSleepTime = 0
	a.totalTime = 0

	oNlogN := float32(float64(a.lineNum) * math.Log(float64(a.lineNum)))
	oNSqrd := float32(math.Pow(float64(a.lineNum), 2))

	QS_SLEEP = time.Duration(float32(time.Second) * a.SleepMultiplier / oNlogN)          // O(n log n)
	CHANGE_SLEEP = QS_SLEEP
	MS_SLEEP = time.Duration(float32(time.Second) * 2 * a.SleepMultiplier / oNlogN)  // O(n log n)
	BBL_SLEEP = time.Duration(float32(time.Second) * 2 * a.SleepMultiplier / oNSqrd)   // O(n^2)
	INST_SLEEP = time.Duration(float32(time.Second) * 2 * a.SleepMultiplier / oNSqrd)	 // O(n^2)
	SHL_SLEEP = time.Duration(float32(time.Second) * 2 * a.SleepMultiplier / float32(math.Pow(float64(a.lineNum), 1.5)))   // O(n^(3/2)) 
	CCT_SLEEP = time.Duration(float32(time.Second) * 2 * a.SleepMultiplier / oNSqrd)  // O(n^2)
	COMB_SLEEP = time.Duration(float32(time.Second) * a.SleepMultiplier / oNlogN)  // O(n log n)

	if a.Linear {
		a.Data = a.GenerateLinear(0, a.H, a.H/float32(a.lineNum))
	} else {
		a.Data = a.Generate(a.lineNum, a.lineNum*a.nonLinearMult)
	}
}

func (a *AnimArr) getLineY(val float32) float32 {   // Lower case so not exported
	return a.H-((float32(val)/float32(a.lineNum*a.nonLinearMult))*a.H)
}

func (a *AnimArr) drawLineOrDot(i int, colour rl.Color) {  // English spelling
	var x = float32((i*a.LineWidth)+(a.LineWidth/2))
	var y float32
	if a.ColorOnly {
		y = 0
	} else if a.Linear {
		y = a.H-a.Data[i]
	} else {
		y = a.getLineY(a.Data[i])
	}
	if a.Dots && !a.ColorOnly {
		radius := float32(a.LineWidth)/2
		rl.DrawCircle(int32(x), int32(y + radius), radius, colour)
	} else {
		rl.DrawLineEx(rl.NewVector2(x, a.H), rl.NewVector2(x, y), float32(a.LineWidth), colour)
	}
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
			//clr = rl.Lime
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

		a.drawLineOrDot(i, clr)
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
		if a.totalTime > 0 {
			rl.DrawText(fmt.Sprintf("Our time: %f", a.totalTime), 10, 110, 20, rl.LightGray)
			rl.DrawText(fmt.Sprintf("Real time: %f", a.totalTime-a.totalSleepTime), 10, 130, 20, rl.LightGray)
		}
		if a.totalSleepTime > 0 {
			rl.DrawText(fmt.Sprintf("Total sleep time: %f", a.totalSleepTime), 10, 150, 20, rl.LightGray)
		}
	}
}

func (a *AnimArr) changeLineWidth(amount int) {
	newWidth := a.LineWidth + amount
	scrW := *ScreenWidth
	scrH := *ScreenHeight
	if newWidth > 0 && newWidth < scrW {
		a.LineWidth = newWidth
		a.Init(float32(scrW), float32(scrH), a.LineWidth, a.Linear, a.ColorOnly, a.Dots, NON_LINEAR_VARIANCE)
	}
}

func (a *AnimArr) Update() {
	if rl.IsKeyPressed(rl.KeyC) {
		a.ColorOnly = !a.ColorOnly
		//a.Dots = false
	}

	if rl.IsKeyPressed(rl.KeyD) {
		a.Dots = !a.Dots
		if a.ColorOnly {
			a.ColorOnly = false
		}
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		a.Sorted = true
		println("Stopping...")
	}

	/*
	if a.lastActive != a.Active {
		if a.Active > -1 {
			rl.SetSoundPitch(a.beepSound, 1/a.lastPitchShift)
			pitchShift := 0.5 + (a.Data[a.Active]/a.maxValue)
			rl.SetSoundPitch(a.beepSound, pitchShift)
			rl.PlaySound(a.beepSound)
			a.lastPitchShift = pitchShift
		}
		a.lastActive = a.Active
	}
	*/

	if !a.Sorting && !a.Shuffling && !a.Showcase {
		mouseMv := int(rl.GetMouseWheelMove())
		if math.Abs(float64(mouseMv)) > 0 {
			a.changeLineWidth(mouseMv)
		}

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
		} else if rl.IsKeyPressed(rl.KeySeven) {
			a.DoSort("comb")
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
	} else {
		a.totalTime += float64(rl.GetFrameTime())
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
	a.totalSleepTime = 0
	a.totalTime = 0
	if sort == "quick" {
		a.CurrentText = "Quick Sort"
		go func() {
			startTime := time.Now()
			a.QuickSort(0, len(a.Data))
			println("AAA:", time.Since(startTime).Seconds())
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
	} else if sort == "comb" {
		a.CurrentText = "Comb Sort"
		go func() {
			a.CombSort()
			a.resetVals()
		}()
	} else {
		panic("Invalid sort: "+sort)
	}
}
