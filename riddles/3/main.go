package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"syscall/js"
)

type ButtonStruct struct {
	btnx  float64
	btny  float64
	btnsx float64
	btnsy float64
}

func main() {
	done := make(chan struct{})

	doc := js.Global().Get("document")
	window := js.Global().Get("window")
	canvas := doc.Call("getElementById", "1")
	ctx := canvas.Call("getContext", "2d")

	var clientwidth float64

	getwidth := func() {
		clientwidth = doc.Get("body").Get("clientWidth").Float()

		canvas.Set("width", clientwidth)
		canvas.Set("height", math.Round((clientwidth/1903)*9000))
	}

	getwidth()

	green := js.Global().Get("Image").New()
	green.Set("src", "riddles/correct.png")

	img := js.Global().Get("Image").New()

	onload_img := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx.Call("drawImage", img, 0, 0, canvas.Get("width"), canvas.Get("height"))
		return nil
	})
	defer onload_img.Release()

	img.Set("onload", onload_img)
	img.Set("src", "riddles/rng.png")

	draw := func() {
		ctx.Call("reset")
		ctx.Call("drawImage", img, 0, 0, canvas.Get("width"), canvas.Get("height"))
	}

	buttons := make([]ButtonStruct, 46)

	calc := func() {
		// 1
		one_buttonWidth := math.Round(83.0 * (clientwidth / 1903))
		one_buttonHeight := math.Round(56.0 * (clientwidth / 1903))

		one_buttonRows := [4]float64{math.Round(1765.0 * (clientwidth / 1903)), math.Round(1859.0 * (clientwidth / 1903)), math.Round(1954.0 * (clientwidth / 1903)), math.Round(2048.0 * (clientwidth / 1903))}

		one_buttonColumns := [4]float64{math.Round(864.0 * (clientwidth / 1903)), math.Round(956.0 * (clientwidth / 1903)), math.Round(1052.0 * (clientwidth / 1903)), math.Round(1148.0 * (clientwidth / 1903))}

		// 2
		two_buttonWidth := math.Round(628.0 * (clientwidth / 1903))
		two_buttonHeight := math.Round(354.0 * (clientwidth / 1903))

		two_buttonRows := [2]float64{math.Round(2920.0 * (clientwidth / 1903)), math.Round(3290.0 * (clientwidth / 1903))}

		two_buttonColumns := [3]float64{math.Round(8.0 * (clientwidth / 1903)), math.Round(639.0 * (clientwidth / 1903)), math.Round(1271.0 * (clientwidth / 1903))}

		// 3
		three_buttonWidth := math.Round(618.0 * (clientwidth / 1903))
		three_buttonHeight := math.Round(348.0 * (clientwidth / 1903))

		three_buttonRows := [1]float64{math.Round(4343.0 * (clientwidth / 1903))}

		three_buttonColumns := [3]float64{math.Round(20.0 * (clientwidth / 1903)), math.Round(642.0 * (clientwidth / 1903)), math.Round(1269.0 * (clientwidth / 1903))}

		// 4
		four_buttonWidth := math.Round(579.0 * (clientwidth / 1903))
		four_buttonHeight := math.Round(652.0 * (clientwidth / 1903))

		four_buttonRows := [1]float64{math.Round(5353.0 * (clientwidth / 1903))}

		four_buttonColumns := [2]float64{math.Round(376.0 * (clientwidth / 1903)), math.Round(955.0 * (clientwidth / 1903))}

		// 5
		five_buttonWidth := math.Round(208.0 * (clientwidth / 1903))
		five_buttonHeight := math.Round(73.0 * (clientwidth / 1903))

		five_buttonRows := [4]float64{math.Round(6647.0 * (clientwidth / 1903)), math.Round(6742.0 * (clientwidth / 1903)), math.Round(6831.0 * (clientwidth / 1903)), math.Round(6915.0 * (clientwidth / 1903))}

		five_buttonColumns := [4]float64{math.Round(104.0 * (clientwidth / 1903)), math.Round(559.0 * (clientwidth / 1903)), math.Round(1086.0 * (clientwidth / 1903)), math.Round(1569.0 * (clientwidth / 1903))}

		// 6
		six_buttonWidth := math.Round(440.0 * (clientwidth / 1903))
		six_buttonHeight := math.Round(195.0 * (clientwidth / 1903))

		six_buttonRows := [3]float64{math.Round(7637.0 * (clientwidth / 1903)), math.Round(7853.0 * (clientwidth / 1903)), math.Round(8069.0 * (clientwidth / 1903))}

		six_buttonColumns := [1]float64{math.Round(763.0 * (clientwidth / 1903))}

		buttonstemp := make([]ButtonStruct, 46)

		temp := 0
		for i := range one_buttonRows {
			for u := range one_buttonColumns {
				buttonstemp[temp].btnx = one_buttonColumns[u]
				buttonstemp[temp].btny = one_buttonRows[i]
				buttonstemp[temp].btnsx = one_buttonWidth
				buttonstemp[temp].btnsy = one_buttonHeight
				temp++
			}
		}
		for i := range two_buttonRows {
			for u := range two_buttonColumns {
				buttonstemp[temp].btnx = two_buttonColumns[u]
				buttonstemp[temp].btny = two_buttonRows[i]
				buttonstemp[temp].btnsx = two_buttonWidth
				buttonstemp[temp].btnsy = two_buttonHeight
				temp++
			}
		}
		for i := range three_buttonRows {
			for u := range three_buttonColumns {
				buttonstemp[temp].btnx = three_buttonColumns[u]
				buttonstemp[temp].btny = three_buttonRows[i]
				buttonstemp[temp].btnsx = three_buttonWidth
				buttonstemp[temp].btnsy = three_buttonHeight
				temp++
			}
		}
		for i := range four_buttonRows {
			for u := range four_buttonColumns {
				buttonstemp[temp].btnx = four_buttonColumns[u]
				buttonstemp[temp].btny = four_buttonRows[i]
				buttonstemp[temp].btnsx = four_buttonWidth
				buttonstemp[temp].btnsy = four_buttonHeight
				temp++
			}
		}
		for u := range five_buttonColumns {
			for i := range five_buttonRows {
				buttonstemp[temp].btnx = five_buttonColumns[u]
				buttonstemp[temp].btny = five_buttonRows[i]
				buttonstemp[temp].btnsx = five_buttonWidth
				buttonstemp[temp].btnsy = five_buttonHeight
				temp++
			}
		}
		for i := range six_buttonRows {
			for u := range six_buttonColumns {
				buttonstemp[temp].btnx = six_buttonColumns[u]
				buttonstemp[temp].btny = six_buttonRows[i]
				buttonstemp[temp].btnsx = six_buttonWidth
				buttonstemp[temp].btnsy = six_buttonHeight
				temp++
			}
		}

		buttons = buttonstemp

	}

	calc()

	one_correct := []int{rand.Intn(4), rand.Intn(4), rand.Intn(4), rand.Intn(4)}
	two_correct := []int{}

	for {
		temp3 := rand.Intn(6)
		if !(slices.Contains(two_correct, temp3)) {
			two_correct = append(two_correct, temp3)
		}
		if len(two_correct) == 4 {
			break
		}
	}

	three_correct := rand.Intn(3)
	four_correct := rand.Intn(2)
	five_state := 0
	six_correct := rand.Intn(3)

	result_state := 0

	pressed := []int{}

	clickHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
	defer clickHandler.Release()

	finish := func() {
		size := fmt.Sprintf("%fpx Arial", math.Round(108.0*(clientwidth/1903)))
		ctx.Set("textRendering", "geometricPrecision")
		ctx.Set("fillStyle", "#FFFFFF")
		ctx.Set("font", size)
		const code string = "<3--LQb5o3-->"
		ctx.Call("fillText", code, math.Round(650.0*(clientwidth/1903)), math.Round(8750.0*(clientwidth/1903)))
	}

	afterclick := func(q bool, n int) {
		if q {
			pressed = append(pressed, n)
			sound := js.Global().Get("Audio").New("riddles/correct.ogg")
			sound.Call("play")
			ctx.Call("drawImage", green, buttons[n].btnx, buttons[n].btny, buttons[n].btnsx, buttons[n].btnsy)
			if result_state == 11 {
				canvas.Call("removeEventListener", "click", clickHandler)
				finish()
				sound1 := js.Global().Get("Audio").New("riddles/finished.ogg")
				sound1.Call("play")
			}
		} else {
			pressed = nil
			sound2 := js.Global().Get("Audio").New("riddles/wrong.ogg")
			sound2.Call("play")
			five_state = 0
			result_state = 0
			draw()
		}
	}

	check := func(n int) {
		switch n {
		case 0, 1, 2, 3:
			if (one_correct[0] == n) && (result_state == 0) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 4, 5, 6, 7:
			if (one_correct[1] == (n - 4)) && (result_state == 1) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 8, 9, 10, 11:
			if (one_correct[2] == (n - 8)) && (result_state == 2) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 12, 13, 14, 15:
			if (one_correct[3] == (n - 12)) && (result_state == 3) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 16, 17, 18, 19, 20, 21:
			if (slices.Contains(two_correct, (n - 16))) && (result_state == 4 || result_state == 5 || result_state == 6 || result_state == 7) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 22, 23, 24:
			if (three_correct == (n - 22)) && (result_state == 8) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 25, 26:
			if (four_correct == (n - 25)) && (result_state == 9) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42:
			if (five_state == ((n - 26) - 1)) && (result_state == 10) {
				five_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		case 43, 44, 45:
			if (six_correct == (n - 43)) && (slices.Contains(pressed, 42)) {
				result_state++
				afterclick(true, n)
			} else {
				afterclick(false, n)
			}
		default:
			return
		}
	}

	click := func(x float64, y float64) int {
		for i := range buttons {
			if (x >= buttons[i].btnx && x <= buttons[i].btnx+buttons[i].btnsx && y >= buttons[i].btny && y <= buttons[i].btny+buttons[i].btnsy) && !(slices.Contains(pressed, i)) {
				return i
			}
		}
		return 99
	}

	clickHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		rect := canvas.Call("getBoundingClientRect")
		x := event.Get("clientX").Float() - rect.Get("left").Float()
		y := event.Get("clientY").Float() - rect.Get("top").Float()

		check(click(x, y))

		return nil
	})
	defer clickHandler.Release()

	screenchange := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		getwidth()
		draw()
		calc()
		for i := range pressed {
			ctx.Call("drawImage", green, buttons[pressed[i]].btnx, buttons[pressed[i]].btny, buttons[pressed[i]].btnsx, buttons[pressed[i]].btnsy)
		}
		if result_state == 11 {
			finish()
		}
		return nil
	})
	defer screenchange.Release()

	window.Call("addEventListener", "resize", screenchange)
	canvas.Call("addEventListener", "click", clickHandler)

	<-done
}
