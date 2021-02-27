package life

import (
	"math"

	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/color"
)

const (
	maxGenerationCount float64 = 50
)

func generateColor(generationCount int) color.Color {
	var green = 0 + float32(math.Min(float64(generationCount)/maxGenerationCount, 1))
	var red = 1.0 - float32(math.Min(float64(generationCount)/maxGenerationCount, 1))
	var blue float32 = 0
	var alpha float32 = 0.7
	return color.Color{Red: red, Green: green, Blue: blue, Alpha: alpha}
}
