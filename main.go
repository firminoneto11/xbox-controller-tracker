package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"sync"
)

// DOCS: Keys can be found at:
// https://github.com/go-vgo/robotgo/blob/master/docs/keys.md

func pressKey(key string, group *sync.WaitGroup) {
	group.Add(1)
	robotgo.KeyDown(key)
	robotgo.KeyUp(key)
	group.Done()
}

func listenXboxJoyStick() error {
	os.Setenv("SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS", "1")

	var (
		joysticks []*sdl.Joystick
		group     sync.WaitGroup
	)
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	sdl.JoystickEventState(sdl.ENABLE)

	fmt.Println("Capturing joystick events...")

	const (
		XboxGuideButton = 10
		key             = "insert"
	)
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch polledEvent := event.(type) {

			case *sdl.QuitEvent:
				running = false

			case *sdl.JoyButtonEvent:
				if polledEvent.State == sdl.PRESSED && polledEvent.Button == XboxGuideButton {
					go pressKey(key, &group)
					fmt.Printf("'%s' was pressed!\n", key)
				}

			case *sdl.JoyDeviceAddedEvent:
				joystick := sdl.JoystickOpen(int(polledEvent.Which))
				if joystick != nil {
					joysticks = append(joysticks, joystick)
					fmt.Println("Joystick", polledEvent.Which, "connected")
				}

			case *sdl.JoyDeviceRemovedEvent:
				joystick := joysticks[int(polledEvent.Which)]
				if joystick != nil {
					joystick.Close()
				}
				fmt.Println("Joystick", polledEvent.Which, "disconnected")
			}
		}
		// Waiting 10 milliseconds before polling the events again
		sdl.Delay(10)
	}

	group.Wait()

	return err
}

func main() {
	errorWhileExecuting := listenXboxJoyStick()
	if errorWhileExecuting != nil {
		os.Exit(1)
	}
}
