package engine

type EngineFactory interface {
	CreateEngine(Options) GameEngine
	CreateUtils() EngineUtilityProvider
}

type EngineUtilityProvider interface {
	CreateDrawable(points []float32) Drawable
}

type Starter interface {
	Start()
}

type GameEngine interface {
	RegisterProgram(Program) Starter
	Close()
}

type Program interface {
	Tick(*TickInfo)
	Draw(*DrawFunctions)
}
