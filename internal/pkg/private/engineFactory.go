package private

import (
	"log"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/public"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type engineFactory struct {
	window *glfw.Window
}

func (factory engineFactory) CreateEngine(options public.EngineOptions) public.GameEngine {
	oglInfo := initOpenGl(options)

	font, err := loadFont("luxisr.ttf", 1)
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
	engine.drawFuncs = &public.DrawFunctions{
		SetColor: engine.setColor,
		DrawTriangles: drawTriangles,
		Clear: engine.clear,
	}
	return &engine
}

func (factory engineFactory) CreateUtils() public.EngineUtilityProvider {
	return utilityProvider{}
}

func EngineFactory(options public.InitOptions) public.EngineFactory {
	window := initGlfw(options)
	return engineFactory{window}
}
