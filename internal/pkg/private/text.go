package private

import (
	"os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/gltext"
)

// loadFont loads the specified font at the given scale.
func loadFont(file string, scale int32) (*gltext.Font, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return gltext.LoadTruetype(fd, scale, 32, 127, gltext.LeftToRight)
}

// drawString draws the same string for each loaded font.
func (engine *gameEngine) drawString(x, y float32, str string) error {
	if engine.font != nil {
		gl.Color4f(1, 1, 1, 1)
		err := engine.font.Printf(0, 450, str)
		if err != nil {
			return err
		}
	}
	return nil
}
