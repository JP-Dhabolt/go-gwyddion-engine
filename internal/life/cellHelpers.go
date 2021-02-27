package life

import (
	"math"
	"math/rand"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/draw"
)

// DrawableFuncs is an interface to allow for testing of draw functions
type DrawableFuncs interface {
	CreateDrawable([]float32) draw.Drawable
}

// CellHelper helps manage the cells
type CellHelper struct {
	Columns   int
	Rows      int
	Threshold float64
	Funcs     DrawableFuncs
}

// MakeCells instantiates the cells
func (ch *CellHelper) MakeCells() [][]*Cell {
	cells := make([][]*Cell, int(math.Max(float64(ch.Columns), float64(ch.Rows))), int(math.Max(float64(ch.Columns), float64(ch.Rows))))
	for x := 0; x < ch.Columns; x++ {
		for y := 0; y < ch.Rows; y++ {
			c := ch.newCell(x, y)

			c.alive = rand.Float64() < ch.Threshold
			c.aliveNext = c.alive
			c.justDied = false
			c.generationsAlive = 0

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

func (ch *CellHelper) newCell(x, y int) *Cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(ch.Columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(ch.Rows)
			position = float32(y) * size
		default:
			continue
		}
		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &Cell{
		drawable: ch.Funcs.CreateDrawable(points),
		x:        x,
		y:        y,
	}
}
