package main

import (
	"math/rand"
	"runtime"

	gamesEngine "github.com/JP-Dhabolt/go-gwyddion-engine/pkg/public"
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/start"
)

var (
	square = []float32{
		-10, 10, 0,
		-10, -10, 0,
		10, -10, 0,
		-10, 10, 0,
		10, 10, 0,
		10, -10, 0,
	}
)

func init() {
	runtime.LockOSThread()
}

func main() {
	rand.Seed(123)

	initOptions := gamesEngine.InitOptions{
		Width: 600,
		Height: 400,
		Title: "Integration",
		Resizeable: true,
	}

	factory := start.Init(initOptions)
	utils := factory.CreateUtils()

	options := gamesEngine.EngineOptions{
		Fps: 60,
		VertexShaderLocation: "./shaders/vertexShader.glsl",
		FragmentShaderLocation: "./shaders/fragmentShader.glsl",
	}

	engine := factory.CreateEngine(options)
	defer engine.Close()
	program := integrationProgram{
		utils,
	}

	engine.RegisterProgram(program).Start()
}

type integrationProgram struct {
	utils gamesEngine.EngineUtilityProvider
}

func (p integrationProgram) Draw(functions gamesEngine.DrawFunctions) {
	functions.SetColor(gamesEngine.Color{Alpha: 1})
	drawable := p.utils.CreateDrawable(square)
	functions.DrawTriangles(drawable, int32(len(square)/3))
}

func (p integrationProgram) Tick(info gamesEngine.TickInfo) {

}
