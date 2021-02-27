package life

import (
	"fmt"
	"testing"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/draw"
)

func TestCell_checkState_should_set_blinker(t *testing.T) {
	ch := CellHelper{
		Columns:   5,
		Rows:      5,
		Threshold: 0,
		Funcs:     MockDrawFuncs{},
	}
	cells := ch.MakeCells()
	cells[2][1].alive = true
	cells[2][1].aliveNext = true
	cells[2][2].alive = true
	cells[2][2].aliveNext = true
	cells[2][3].alive = true
	cells[2][3].aliveNext = true

	fmt.Printf("Initial setup\n%s", formatCellsString(cells))

	setupCells(cells)
	cycleCells(cells)
	if !isBlinkerHorizontal(cells) {
		t.Fatalf("Expected horizontal blinker, actually \n%s", formatCellsString(cells))
	}
	fmt.Printf("After first cycle\n%s", formatCellsString(cells))
	setupCells(cells)
	cycleCells(cells)
	if !isBlinkerVertical(cells) {
		t.Fatalf("Expected vertical blinker, actually \n%s", formatCellsString(cells))
	}

}

func setupCells(cells [][]*Cell) {
	for x := range cells {
		for _, c := range cells[x] {
			c.setupState(cells)
		}
	}
}

func cycleCells(cells [][]*Cell) {
	for x := range cells {
		for _, c := range cells[x] {
			c.transitionState()
		}
	}
}

func formatCellsString(cells [][]*Cell) string {
	formatString := ""
	for y := range cells[0] {
		for x := range cells {
			nextVal := ""
			if cells[x][y].aliveNext {
				nextVal = "*"
			} else {
				nextVal = "."
			}
			formatString += nextVal
		}
		formatString += "\n"
	}
	return formatString
}

func isBlinkerHorizontal(cells [][]*Cell) bool {
	return cells[1][2].aliveNext &&
		cells[2][2].aliveNext &&
		cells[3][2].aliveNext
}

func isBlinkerVertical(cells [][]*Cell) bool {
	return cells[2][1].aliveNext &&
		cells[2][2].aliveNext &&
		cells[2][3].aliveNext
}

type MockDrawFuncs struct{}

func (MockDrawFuncs) CreateDrawable([]float32) draw.Drawable {
	return 0
}
