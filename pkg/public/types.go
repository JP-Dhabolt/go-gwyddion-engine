package public

import "time"

type TickInfo struct {
	StartTime   time.Time
	CurrentTime time.Time
	TickNumber  int
}

type DrawFunctions struct {
	SetColor      func(Color)
	DrawTriangles func(Drawable, int32)
}

type Drawable uint32
