package engine

import (
	"log"

	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/font"
	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/opengl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Factory is used to enable OpenGL functiality before creating the engine itself
type Factory struct {
	window *glfw.Window
}

/*
Options are settings for the game engine.

Fps indicates the desired frames-per-second that the engine will try to maintain.
VertexShaderLocation is an optional path to a custom Vertex Shader file in glsl format.
FragmentShaderLocation is an optional path to a custom Fragment Shader file in glsl format.

If a shader location is not provided, a simple default shader will be used.
*/
type Options = opengl.OpenGlOptions

/*
FactoryOptions are settings for intializing the engine factory, which includes the windows

Resizable indicates whether the created window should be able to be resized.
Width is the number of pixels wide the created window should be.
Height is the number of pixels tall the creaed window should be.
Title is the window title.
*/
type FactoryOptions = opengl.InitOptions

// CreateEngine returns a new GameEngine object created with the provided options
func (f Factory) CreateEngine(options Options) *GameEngine {
	oglInfo := opengl.Initialize(options)

	font, err := font.LoadDefault(1)
	if err != nil {
		panic(err)
	}

	log.Println("Opened font")

	engine := GameEngine{
		window:  f.window,
		glProg:  oglInfo.Prog,
		myLoc:   oglInfo.Loc,
		options: options,
		font:    font,
	}
	engine.drawFuncs = &DrawFunctions{
		SetColor: engine.setColor,
		Clear:    engine.clear,
	}
	return &engine
}

// NewFactory creates a new engine factory instance with the provided options.
func NewFactory(options FactoryOptions) Factory {
	window := opengl.InitGlfw(options)
	return Factory{window}
}
