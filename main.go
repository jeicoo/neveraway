package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Print("--- neveraway started ---\n")
	fmt.Print("move your mouse up or down to exit")

	robotgo.MouseSleep = 7
	sx, sy := robotgo.GetScreenSize()
	sx = sx / 2
	sy = sy / 2

	robotgo.Move(sx, sy)
	x, y := sx, sy

	ix := 1
	for ok := true; ok; ok = (y == sy) {
		x = x + ix
		robotgo.Move(x, y)
		x, y = robotgo.GetMousePos()
		switch x {
		case sx + 40:
			ix = -1
		case sx - 40:
			ix = 1
		}
	}
}
