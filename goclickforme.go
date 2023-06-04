package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Total duration until the next mouse move.
	totalDuration := 45
	keyToPress := "alt"

	for {
		// Get the screen size.
		screenWidth, screenHeight := robotgo.GetScreenSize()

		// Generate a random x and y value within the bounds of the screen size.
		randomX := rand.Intn(screenWidth)
		randomY := rand.Intn(screenHeight)

		// Move the mouse to the random position.
		robotgo.MoveMouse(randomX, randomY)

		// Press the "f" key.
		robotgo.KeyPress(keyToPress)

		// Countdown each second until the next move.
		for i := totalDuration; i > 0; i-- {
			fmt.Println("Next moving event in", i, "seconds...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println(fmt.Sprintf("Moved mouse randomly and pressed %s key.", keyToPress))
	}
}
