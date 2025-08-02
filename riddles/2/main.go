package main

import (
	"fmt"
	"math"
	"syscall/js"
)

func main() {
	done := make(chan struct{})

	doc := js.Global().Get("document")
	window := js.Global().Get("window")
	canvas := doc.Call("getElementById", "1")
	ctx := canvas.Call("getContext", "2d")

	var clientwidth float64
	var buttonX float64
	var buttonY float64
	var buttonWidth float64
	var buttonHeight float64

	calc := func() {
		clientwidth = doc.Get("body").Get("clientWidth").Float()
		buttonX = math.Round(1212.0 * (clientwidth / 1903))
		buttonY = math.Round(1308.0 * (clientwidth / 1903))
		buttonWidth = math.Round(66.0 * (clientwidth / 1903))
		buttonHeight = math.Round(14.0 * (clientwidth / 1903))
	}
	calc()

	clicked := false

	const code string = "<2--R2MqiY-->"

	finish := func() {
		size := fmt.Sprintf("%fpx Arial", math.Round(14.0*(clientwidth/1903)))
		ctx.Set("textRendering", "geometricPrecision")
		ctx.Set("fillStyle", "#CCCCCC")
		ctx.Set("font", size)
		ctx.Call("fillText", code, math.Round(774.0*(clientwidth/1903)), math.Round(2700.0*(clientwidth/1903)))
	}

	clickHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		rect := canvas.Call("getBoundingClientRect")
		x := event.Get("clientX").Float() - rect.Get("left").Float()
		y := event.Get("clientY").Float() - rect.Get("top").Float()

		if x >= buttonX && x <= buttonX+buttonWidth && y >= buttonY && y <= buttonY+buttonHeight {
			sound := js.Global().Get("Audio").New("riddles/found.ogg")
			sound.Call("play")
			if !clicked {
				finish()
				clicked = true
			}
		}
		return nil
	})
	defer clickHandler.Release()

	mousein := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		rect := canvas.Call("getBoundingClientRect")
		x := event.Get("clientX").Float() - rect.Get("left").Float()
		y := event.Get("clientY").Float() - rect.Get("top").Float()

		if x >= buttonX && x <= buttonX+buttonWidth && y >= buttonY && y <= buttonY+buttonHeight {
			doc.Get("body").Get("style").Set("cursor", "pointer")
		} else {
			doc.Get("body").Get("style").Set("cursor", "default")
		}
		return nil
	})
	defer mousein.Release()

	screenchange := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		calc()
		if clicked {
			finish()
		}
		return nil
	})
	defer screenchange.Release()

	canvas.Call("addEventListener", "click", clickHandler)
	canvas.Call("addEventListener", "mousemove", mousein)
	window.Call("addEventListener", "resize", screenchange)

	<-done
}
