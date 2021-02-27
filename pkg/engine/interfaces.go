package engine

/*
Program interface is what is expected by the game engine to operate.

Tick function gets executed every frame, and is intended to modify game state
Draw function gets executed every frame, after Tick, and is intended to draw the current frame to the screen
*/
type Program interface {
	Tick(*TickInfo)
	Draw(*DrawFunctions)
}
