package life

import (
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/draw"
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/engine"
)

// Cell represents a cell in the automation
type Cell struct {
	drawable draw.Drawable

	alive     bool
	aliveNext bool
	justDied  bool

	generationsAlive int

	x int
	y int
}

func (c *Cell) draw(functions *engine.DrawFunctions) {
	if !c.alive {
		return
	}

	functions.SetColor(generateColor(c.generationsAlive))
	draw.WriteDrawable(c.drawable, int32(len(square)/3))
}

func (c *Cell) transitionState() {
	c.alive = c.aliveNext
}

// setupState determines the state of the cell for the next tick of the game.
func (c *Cell) setupState(cells [][]*Cell) {
	if c.justDied {
		c.justDied = false
	}

	liveCount := c.liveNeighbors(cells)
	if c.alive {
		// 1. Any live cell with fewer than two live neighbors dies, as if caused by underpopulation
		if liveCount < 2 {
			c.aliveNext = false
			c.justDied = true
			c.generationsAlive = 0
		}

		// 2. Any live cell with two or three live neighbors lives on to the next generation
		if liveCount == 2 || liveCount == 3 {
			c.aliveNext = true
			c.generationsAlive++
		}

		// 3. Any live cell with more than three live neighbors dies, as if by overpopulation
		if liveCount > 3 {
			c.aliveNext = false
			c.justDied = true
			c.generationsAlive = 0
		}
	} else {
		// 4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction
		if liveCount == 3 {
			c.aliveNext = true
			c.generationsAlive++
		}
	}
}

// liveNeighbors returns the number of live neighbors for a cell
func (c *Cell) liveNeighbors(cells [][]*Cell) int {
	var liveCount int
	add := func(x, y int) {
		// If we're at an edge, check the other side of the board.
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].alive {
			liveCount++
		}
	}

	add(c.x-1, c.y)   // To the left
	add(c.x+1, c.y)   // To the right
	add(c.x, c.y+1)   // up
	add(c.x, c.y-1)   // down
	add(c.x-1, c.y+1) // top left
	add(c.x+1, c.y+1) // top right
	add(c.x-1, c.y-1) // bottom left
	add(c.x+1, c.y-1) // bottom right

	return liveCount
}
