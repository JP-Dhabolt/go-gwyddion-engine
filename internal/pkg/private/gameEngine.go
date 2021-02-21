package private

import (
	"time"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/color"
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/public"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gltext"
)

type gameEngine struct {
	window    *glfw.Window
	glProg    uint32
	myLoc     int32
	options   public.EngineOptions
	program   public.Program
	font      *gltext.Font
	drawFuncs *public.DrawFunctions
}

func (engine *gameEngine) setColor(color color.Color) {
	gl.ProgramUniform4f(engine.glProg, engine.myLoc, color.Red, color.Green, color.Blue, color.Alpha)
}

func (engine *gameEngine) Start() {
	tickInfo := public.TickInfo{StartTime: time.Now(), CurrentTime: time.Now(), TickNumber: 1}
	for !engine.window.ShouldClose() {
		engine.tick(&tickInfo)
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

func (engine *gameEngine) tick(info *public.TickInfo) {
	engine.program.Tick(info)
}

func (engine *gameEngine) draw() {
	gl.UseProgram(engine.glProg)

	engine.program.Draw(engine.drawFuncs)

	glfw.PollEvents()
	engine.window.SwapBuffers()
}

func (engine *gameEngine) clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (engine *gameEngine) drawString(x, y float32, str string) error {
	if engine.font != nil {
		err := engine.font.Printf(x, y, str)
		if err != nil {
			return err
		}
	}
	return nil
}

func (engine *gameEngine) Close() {
	glfw.Terminate()
}
