package private

import (
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine.git/pkg/public"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

type engineFactory struct {
	window *glfw.Window
}

func (factory engineFactory) CreateEngine(options public.EngineOptions) public.GameEngine {
	oglInfo := initOpenGl(options)

	//font, err := loadFont("../gwyddionGamesEngine/luxisr.ttf", 1)
	font, err := loadFont("luxisr", 1)
	if err != nil {
		panic(err)
	}

	log.Println("Opened font")

	return &gameEngine{
		window:  factory.window,
		glProg:  oglInfo.Prog,
		myLoc:   oglInfo.Loc,
		options: options,
		font:    font,
	}
}

func (factory engineFactory) CreateUtils() public.EngineUtilityProvider {
	return createUtilityProvider()
}

func EngineFactory(options public.InitOptions) public.EngineFactory {
	window := initGlfw(options)
	return engineFactory{window}
}
