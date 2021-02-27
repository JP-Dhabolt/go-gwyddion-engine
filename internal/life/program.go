package life

import "github.com/JP-Dhabolt/go-gwyddion-engine/pkg/engine"

// GameOfLife implements the Program interface
type GameOfLife struct {
	Cells [][]*Cell
}

// Draw draws the current frame of cells
func (game GameOfLife) Draw(functions *engine.DrawFunctions) {
	functions.Clear()
	for x := range game.Cells {
		for _, c := range game.Cells[x] {
			c.transitionState()
			c.draw(functions)
		}
	}
}

// Tick sets up the current state
func (game GameOfLife) Tick(info *engine.TickInfo) {
	for x := range game.Cells {
		for _, c := range game.Cells[x] {
			c.setupState(game.Cells)
		}
	}
}
