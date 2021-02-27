package engine

import (
	"log"

	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/draw"
	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/font"
	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/opengl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type engineFactory struct {
	window *glfw.Window
}

type Options = opengl.OpenGlOptions
type FactoryOptions = opengl.InitOptions

func (factory engineFactory) CreateEngine(options Options) GameEngine {
	oglInfo := opengl.Initialize(options)

	font, err := font.LoadDefault(1)
	if err != nil {
		panic(err)
	}

	log.Println("Opened font")

	engine := gameEngine{
		window:  factory.window,
		glProg:  oglInfo.Prog,
		myLoc:   oglInfo.Loc,
		options: options,
		font:    font,
	}
	engine.drawFuncs = &DrawFunctions{
		SetColor:      engine.setColor,
		DrawTriangles: draw.Triangles,
		Clear:         engine.clear,
	}
	return &engine
}

func (factory engineFactory) CreateUtils() EngineUtilityProvider {
	return draw.UtilityProvider{}
}

// NewFactory creates a new Engine Factory instance
func NewFactory(options FactoryOptions) EngineFactory {
	window := opengl.InitGlfw(options)
	return engineFactory{window}
}
