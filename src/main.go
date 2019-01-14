package main

import (
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 640, 640),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	snake := NewSnake()
	snake.initPositions(4, 4)
	snakeMap := NewSnakeMap(snake.frameSize, 10)

	for !win.Closed() {
		snakeMap.dt = time.Since(snakeMap.last).Seconds()
		snake.initMatrix()

		// Detect move actions
		snakeMap.handleKeys(win)

		// Turn management
		if snakeMap.dt > 0.5 {
			snakeMap.index += snakeMap.move
			snake.moveSnake(snakeMap) // change positions and sprites of the snake
			snakeMap.last = time.Now()
		}

		// win.Clear(colornames.Greenyellow)
		win.Clear(colornames.Firebrick)
		snake.Draw(snakeMap, win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
