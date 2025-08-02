package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	done := make(chan struct{})

	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "159")

	buttonX := 256.0
	buttonY := 128.0
	buttonWidth := 15.0
	buttonHeight := 10.0

	clickHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		rect := canvas.Call("getBoundingClientRect")
		x := event.Get("clientX").Float() - rect.Get("left").Float()
		y := event.Get("clientY").Float() - rect.Get("top").Float()

		if x >= buttonX && x <= buttonX+buttonWidth && y >= buttonY && y <= buttonY+buttonHeight {
			sound := js.Global().Get("Audio").New("riddles/found.ogg")
			sound.Call("play")
			elems := doc.Call("getElementsByTagName", "p")
			const code string = "<1--AtQ2jx-->"
			var answer string = fmt.Sprintf(`<p align="left"><font size="3" face="Arial Black" color="#FFFFFF">Now, we need help.&nbsp; The pictures below are after one weekend of work by Bob Giese and crew.&nbsp; This was a former furniture store with plenty of partitions and fake inner walls.&nbsp;%sIt also needs serious upgrades of insulation, lighting and facilities.&nbsp; We have a very tight budget, volunteers are needed for another couple of weekends (or nights) to whip this into shape.</font></p>`, code)
			elems.Get("24").Set("outerHTML", answer)
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

	mouseon := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		canvas.Call("addEventListener", "mousemove", mousein)
		return nil
	})
	defer mouseon.Release()

	mouseoff := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		canvas.Call("removeEventListener", "mousemove", mousein)
		doc.Get("body").Get("style").Set("cursor", "default")
		return nil
	})
	defer mouseoff.Release()

	canvas.Call("addEventListener", "click", clickHandler)
	canvas.Call("addEventListener", "mouseenter", mouseon)
	canvas.Call("addEventListener", "mouseleave", mouseoff)

	<-done
}
