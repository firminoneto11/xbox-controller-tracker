package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/sys/windows"
	"os"
	"sync"
	"unsafe"
)

var (
	user32DLL = windows.NewLazyDLL("user32.dll")
	SendInput = user32DLL.NewProc("SendInput")
)

type KeyBoardInput struct {
	windowsVirtualKey uint16
	wScan             uint16
	dwFlags           uint32
	time              uint32
	dwExtraInfo       uint64
}

type Input struct {
	inputType     uint32
	keyboardInput KeyBoardInput
	padding       uint64
}

func pressKey(group *sync.WaitGroup) {
	group.Add(1)
	defer group.Done()

	var input Input
	input.inputType = 1
	input.keyboardInput.windowsVirtualKey = 0xD2

	_, _, err := SendInput.Call(uintptr(1), uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	if err != nil {
		fmt.Println(err)
	}
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

	const XboxGuideButton = 10
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch polledEvent := event.(type) {

			case *sdl.QuitEvent:
				running = false

			case *sdl.JoyButtonEvent:
				if polledEvent.State == sdl.PRESSED && polledEvent.Button == XboxGuideButton {
					fmt.Println("Joystick", polledEvent.Which, "button", polledEvent.Button, "pressed")
					go pressKey(&group)
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
		// Waiting 16 milliseconds before polling the events again
		sdl.Delay(16)
	}

	return err
}

func main() {
	errorWhileExecuting := listenXboxJoyStick()
	if errorWhileExecuting != nil {
		os.Exit(1)
	}
}
