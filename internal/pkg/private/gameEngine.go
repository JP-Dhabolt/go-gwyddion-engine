package private

import (
	"github.com/4ydx/gltext/v4.1"
	"github.com/GwyddionGames/go-gwyddion-engine/pkg/public"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"time"
)

type gameEngine struct {
	window  *glfw.Window
	glProg  uint32
	myLoc   int32
	options public.EngineOptions
	program public.Program
	font    *v41.Font
}

func (engine *gameEngine) setColor(color public.Color) {
	gl.ProgramUniform4f(engine.glProg, engine.myLoc, color.Red, color.Green, color.Blue, color.Alpha)
}

func (engine *gameEngine) Start() {
	tickInfo := public.TickInfo{StartTime: time.Now(), CurrentTime: time.Now(), TickNumber: 1}
	for !engine.window.ShouldClose() {
		engine.tick(tickInfo)
		engine.draw()

		time.Sleep(time.Second/time.Duration(engine.options.Fps) - time.Since(tickInfo.CurrentTime))
		tickInfo.TickNumber++
		tickInfo.CurrentTime = time.Now()
	}
}

func (engine *gameEngine) RegisterProgram(program public.Program) public.Starter {
	engine.program = program
	return engine
}

func (engine *gameEngine) tick(info public.TickInfo) {
	engine.program.Tick(info)
}

func (engine *gameEngine) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(engine.glProg)

	engine.program.Draw(public.DrawFunctions{SetColor: engine.setColor, DrawTriangles: drawTriangles})

	_ = engine.drawString(10, 10, "Game of Life") // TODO: Handle error (or don't return error))

	glfw.PollEvents()
	engine.window.SwapBuffers()
}

func (engine *gameEngine) Close() {
	glfw.Terminate()
}
