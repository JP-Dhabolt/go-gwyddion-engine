package life

import "github.com/JP-Dhabolt/go-gwyddion-engine/pkg/draw"

// DrawFuncs implements the DrawableFuncs interface
type DrawFuncs struct{}

// CreateDrawable simply wraps the draw.CreateDrawable call
func (d *DrawFuncs) CreateDrawable(points []float32) draw.Drawable {
	return draw.CreateDrawable(points)
}
