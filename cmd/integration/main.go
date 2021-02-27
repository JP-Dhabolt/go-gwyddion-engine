package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/color"
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/draw"
	gamesEngine "github.com/JP-Dhabolt/go-gwyddion-engine/pkg/engine"
)

func init() {
	runtime.LockOSThread()
}

var colors = []color.Color{
	color.RED,
	color.ORANGE,
	color.YELLOW,
	color.GREEN,
	color.BLUE,
	color.INDIGO,
	color.VIOLET,
}

func main() {
	rand.Seed(123)

	initOptions := gamesEngine.FactoryOptions{
		Width:      600,
		Height:     400,
		Title:      "Integration",
		Resizeable: true,
	}

	factory := gamesEngine.NewFactory(initOptions)
	options := gamesEngine.Options{
		Fps: 60,
	}

	engine := factory.CreateEngine(options)
	defer engine.Close()
	program := integrationProgram{}

	engine.RegisterProgram(&program).Start()
}

type integrationProgram struct {
	iteration int
}

func determineScale(iteration int) float32 {
	return float32(len(colors)-iteration) / float32(len(colors))
}

func (p *integrationProgram) Draw(functions *gamesEngine.DrawFunctions) {
	shape := createSquare(determineScale(p.iteration))
	drawable := draw.CreateDrawable(shape)
	functions.SetColor(colors[p.iteration])
	draw.WriteDrawable(drawable, int32(len(shape)/3))
	functions.SetColor(color.BLACK)
}

func (p *integrationProgram) Tick(info *gamesEngine.TickInfo) {
	timeSinceStart := time.Since(info.StartTime)
	wholeSecondsSinceStart := int(timeSinceStart.Seconds())
	p.iteration = wholeSecondsSinceStart % len(colors)
}

func createSquare(scale float32) []float32 {
	size := scale
	return []float32{
		-size, size, 0,
		-size, -size, 0,
		size, -size, 0,
		-size, size, 0,
		size, size, 0,
		size, -size, 0,
	}
}
