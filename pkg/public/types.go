package public

import (
	"time"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/color"
)

type TickInfo struct {
	StartTime   time.Time
	CurrentTime time.Time
	TickNumber  int
}

type DrawFunctions struct {
	SetColor      func(color.Color)
	DrawTriangles func(Drawable, int32)
	drawText      func(float32, float32, string) error
	Clear         func()
}

type Drawable uint32
