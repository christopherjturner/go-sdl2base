package main

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenX   int32  = 1024
	screenY   int32  = 1024
	TICK      uint64 = 1000 / 60
	GameTitle string = "SDL BASE APP"
)

var (
	render   *sdl.Renderer
	running  bool = true
	mousePos sdl.Point
	textures map[string]*sdl.Texture = make(map[string]*sdl.Texture)
)

func main() {

	// SDL setup stuff
	runtime.LockOSThread()

	// sdl setup
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	// window setup
	window, err := sdl.CreateWindow(GameTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenX, screenY, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// renderer setup
	render, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	render.RenderSetVSync(true)
	defer render.Destroy()

	// load assets
	LoadTextures("./images", textures)
	defer UnloadTextures(textures)

	// setup is done
	fmt.Println("starting game")

	var lastFrameTime uint64 = 0

	for running {
		handleInput()
		update()
		drawGame()

		if frameTime := sdl.GetTicks64(); frameTime-lastFrameTime < TICK {
			sdl.Delay(uint32(TICK - (frameTime - lastFrameTime)))
		}

		lastFrameTime = sdl.GetTicks64()
	}
}

func update() {
	mousePos.X, mousePos.Y, _ = sdl.GetMouseState()
}

func drawGame() {
	render.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	render.SetDrawColor(0x6C, 0x93, 0x75, 255)
	render.Clear()

	// draw stuff here

	render.Present()
}

func handleInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			// quit game
			if t.Keysym.Sym == sdl.K_ESCAPE {
				running = false
			}

		case *sdl.MouseButtonEvent:
			if t.Type == sdl.MOUSEBUTTONDOWN && t.Button == sdl.BUTTON_LEFT {
				// handle mouse down
			}
			if t.Type == sdl.MOUSEBUTTONUP && t.Button == sdl.BUTTON_LEFT {
				// handle mouse up
			}
		}
	}
}
