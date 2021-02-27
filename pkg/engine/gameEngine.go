package engine

import (
	"time"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/color"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gltext"
)

// GameEngine is the workhorse of the project
type GameEngine struct {
	window    *glfw.Window
	glProg    uint32
	myLoc     int32
	options   Options
	program   Program
	font      *gltext.Font
	drawFuncs *DrawFunctions
}

// Starter is used to start the engine
type Starter struct {
	engine *GameEngine
}

func (e *GameEngine) setColor(color color.Color) {
	gl.ProgramUniform4f(e.glProg, e.myLoc, color.Red, color.Green, color.Blue, color.Alpha)
}

// Start starts the game engine until the window is closed
func (s *Starter) Start() {
	tickInfo := TickInfo{StartTime: time.Now(), CurrentTime: time.Now(), TickNumber: 1}
	for !s.engine.window.ShouldClose() {
		s.engine.tick(&tickInfo)
		s.engine.draw()

		time.Sleep(time.Second/time.Duration(s.engine.options.Fps) - time.Since(tickInfo.CurrentTime))
		tickInfo.TickNumber++
		tickInfo.CurrentTime = time.Now()
	}
}

// RegisterProgram registers a ProgramInterface and returns a Starter to start the application
func (e *GameEngine) RegisterProgram(program Program) *Starter {
	e.program = program
	return &Starter{
		engine: e,
	}
}

func (e *GameEngine) tick(info *TickInfo) {
	e.program.Tick(info)
}

func (e *GameEngine) draw() {
	gl.UseProgram(e.glProg)

	e.program.Draw(e.drawFuncs)

	glfw.PollEvents()
	e.window.SwapBuffers()
}

func (e *GameEngine) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (e *GameEngine) drawString(x, y float32, str string) error {
	if e.font != nil {
		err := e.font.Printf(x, y, str)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close is used to gracefully terminate the GameEngine.
// As soon as you have a reference to the GameEngine, call defer engine.Close()
func (e *GameEngine) Close() {
	glfw.Terminate()
}
