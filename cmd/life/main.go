package main

import (
	"flag"
	"math/rand"
	"runtime"
	"time"

	"github.com/JP-Dhabolt/go-gwyddion-engine/internal/life"
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/engine"
)

const (
	width  = 1024 // TODO: Get this info programmatically based on screen size
	height = 768
)

var (
	rows      = 100
	columns   = 100
	seed      = time.Now().UnixNano()
	threshold = 0.15
	fps       = 10
)

func init() {
	flag.IntVar(&columns, "columns", columns, "Sets the number of columns")
	flag.IntVar(&rows, "rows", rows, "Sets the number of rows.")
	flag.Int64Var(&seed, "seed", seed, "Sets the starting seed of the game, used to randomize the initial state.")
	flag.Float64Var(&threshold, "threshold", threshold, "A percentage between 0 and 1 used in conjunction with -seed to determine if the cell starts alive or dead.")
	flag.IntVar(&fps, "fps", fps, "Sets the frames-per-second, used to set the speed of the simulation")
	flag.Parse()

	runtime.LockOSThread()
}

func main() {
	rand.Seed(seed)

	factoryOptions := engine.FactoryOptions{
		Width:      width,
		Height:     height,
		Title:      "Conway's Game of Life",
		Resizeable: true,
	}

	factory := engine.NewFactory(factoryOptions)

	options := engine.Options{
		Fps: fps,
	}

	eng := factory.CreateEngine(options)
	defer eng.Close()

	cellHelper := life.CellHelper{
		Columns:   columns,
		Rows:      rows,
		Threshold: threshold,
		Funcs:     &life.DrawFuncs{},
	}

	gameOfLife := life.GameOfLife{Cells: cellHelper.MakeCells()}

	eng.RegisterProgram(gameOfLife).Start()
}
