package color

// Color represents a color using rgba
type Color struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

// NewColor returns a new Color object using the provided red, green, blue values
// If alpha is omitted, it will default to the value 1.0
// If alpha is included, only the first value will be used
func NewColor(red, green, blue float32, alpha ...float32) Color {
	actualAlpha := float32(1)
	if len(alpha) > 0 {
		actualAlpha = alpha[0]
	}
	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: actualAlpha,
	}
}

var (
	// BLACK is the color
	BLACK = NewColor(0, 0, 0)
	// RED is the color
	RED = NewColor(1, 0, 0)
	// ORANGE is the color
	ORANGE = NewColor(1, 0.5, 0)
	// YELLOW is the color
	YELLOW = NewColor(1, 1, 0)
	// GREEN is the color
	GREEN = NewColor(0, 1, 0)
	// BLUE is the color
	BLUE = NewColor(0, 0, 1)
	// INDIGO is the color
	INDIGO = NewColor(0.25, 0, 0.5)
	// VIOLET is the color
	VIOLET = NewColor(0.5, 0, 1)
	// WHITE is the color
	WHITE = NewColor(1, 1, 1)
)
