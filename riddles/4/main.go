package main

import (
	"fmt"
	"math"
	"math/rand"
	"syscall/js"
	"time"
)

const code string = "<4--wLQkSS-->"

type ButtonStruct struct {
	btnx  float64
	btny  float64
	btnsx float64
	btnsy float64
}

func setTimeout(d time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		close(ch)
		return nil
	})
	js.Global().Call("setTimeout", callback, d.Milliseconds())
	return ch
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
		canvas.Set("height", math.Round((clientwidth/1905)*700))
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
	img.Set("src", "riddles/keypad.png")

	draw := func() {
		ctx.Call("reset")
		ctx.Call("drawImage", img, 0, 0, canvas.Get("width"), canvas.Get("height"))
	}

	finish := func() {
		size := fmt.Sprintf("%fpx Arial", math.Round(64.0*(clientwidth/1905)))
		ctx.Set("textRendering", "geometricPrecision")
		ctx.Set("fillStyle", "#FFFFFF")
		ctx.Set("font", size)
		ctx.Call("fillText", code, math.Round(1.0*(clientwidth/1905)), math.Round(300.0*(clientwidth/1905)))
	}

	buttons := make([]ButtonStruct, 30)
	pressed := []int{}

	calc := func() {
		buttons = nil

		one_buttonWidth := math.Round(67.0 * (clientwidth / 1905))
		one_buttonHeight := math.Round(44.0 * (clientwidth / 1905))

		one_buttonRows := [5]float64{
			math.Round(804.0 * (clientwidth / 1905)),
			math.Round(901.0 * (clientwidth / 1905)),
			math.Round(993.0 * (clientwidth / 1905)),
			math.Round(1085.0 * (clientwidth / 1905)),
			math.Round(1171.0 * (clientwidth / 1905)),
		}

		one_buttonColumns := [6]float64{
			math.Round(241.0 * (clientwidth / 1905)),
			math.Round(323.0 * (clientwidth / 1905)),
			math.Round(396.0 * (clientwidth / 1905)),
			math.Round(465.0 * (clientwidth / 1905)),
			math.Round(530.0 * (clientwidth / 1905)),
			math.Round(597.0 * (clientwidth / 1905)),
		}

		buttonstemp := make([]ButtonStruct, 30)

		temp := 0
		for u := range one_buttonColumns {
			for i := range one_buttonRows {
				buttonstemp[temp].btnx = one_buttonRows[i]
				buttonstemp[temp].btny = one_buttonColumns[u]
				buttonstemp[temp].btnsx = one_buttonWidth
				buttonstemp[temp].btnsy = one_buttonHeight
				temp++
			}
		}

		buttons = buttonstemp
	}

	calc()

	one_correct := []int{}

	rng := func() {
		one_correct = []int{rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30), rand.Intn(30)}
	}

	rng()

	result_state := 0

	windowHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
	defer windowHandler.Release()

	clickHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
	defer clickHandler.Release()

	afterclick := func(q bool, n int) {
		if q {
			pressed = append(pressed, n)
			sound := js.Global().Get("Audio").New("riddles/correct.ogg")
			sound.Call("play")
			ctx.Call("drawImage", green, buttons[n].btnx, buttons[n].btny, buttons[n].btnsx, buttons[n].btnsy)
			result_state++
			if result_state == 10 {
				canvas.Call("removeEventListener", "click", clickHandler)
				finish()
				sound1 := js.Global().Get("Audio").New("riddles/finished.ogg")
				sound1.Call("play")
			}
		} else {
			canvas.Call("removeEventListener", "click", clickHandler)
			pressed = nil
			result_state = 0
			sound2 := js.Global().Get("Audio").New("riddles/wrong.ogg")
			sound2.Call("play")
			draw()
			for i := range one_correct {
				ctx.Call("drawImage", green, buttons[one_correct[i]].btnx, buttons[one_correct[i]].btny, buttons[one_correct[i]].btnsx, buttons[one_correct[i]].btnsy)
			}
			<-setTimeout(3 * time.Second)
			rng()
			draw()
			window.Call("addEventListener", "click", windowHandler)
		}
	}

	check := func(n int) {
		if one_correct[result_state] == n {
			go afterclick(true, n)
		} else if n == 99 {
			return
		} else {
			go afterclick(false, n)
		}
	}

	click := func(x float64, y float64) int {
		for i := range buttons {
			if x >= buttons[i].btnx && x <= buttons[i].btnx+buttons[i].btnsx && y >= buttons[i].btny && y <= buttons[i].btny+buttons[i].btnsy {
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

	start := func() {
		<-setTimeout(1 * time.Second)
		for i := range one_correct {
			sound4 := js.Global().Get("Audio").New("riddles/beep.ogg")
			sound4.Call("play")
			ctx.Call("drawImage", green, buttons[one_correct[i]].btnx, buttons[one_correct[i]].btny, buttons[one_correct[i]].btnsx, buttons[one_correct[i]].btnsy)
			<-setTimeout(1500 * time.Millisecond)
			draw()
			<-setTimeout(500 * time.Millisecond)
		}
		canvas.Call("addEventListener", "click", clickHandler)
	}

	windowHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		window.Call("removeEventListener", "click", windowHandler)
		go start()
		return nil
	})
	defer windowHandler.Release()

	screenchange := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		getwidth()
		draw()
		calc()
		for i := range pressed {
			ctx.Call("drawImage", green, buttons[pressed[i]].btnx, buttons[pressed[i]].btny, buttons[pressed[i]].btnsx, buttons[pressed[i]].btnsy)
		}
		if result_state == 10 {
			finish()
		}
		return nil
	})
	defer screenchange.Release()

	window.Call("addEventListener", "resize", screenchange)
	window.Call("addEventListener", "click", windowHandler)

	<-done
}
