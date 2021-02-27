package engine

import (
	"time"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/color"
)

/*
TickInfo provides information about the current tick of the engine's loop

StartTime is the time that the engine started.
CurrentTime is the current time for this tick.
TickNumber is the number of ticks since the engine started.
*/
type TickInfo struct {
	StartTime   time.Time
	CurrentTime time.Time
	TickNumber  int
}

/*
DrawFunctions are a collection of drawing utility functions

SetColor sets the color being used by the engine
Clear resets the screen to a blank canvas
*/
type DrawFunctions struct {
	SetColor func(color.Color)
	Clear    func()
}
