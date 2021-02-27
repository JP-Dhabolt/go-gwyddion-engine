package font

import (
	"os"

	"github.com/go-gl/gltext"
)

// Load loads the specified font at the given scale.
func Load(file string, scale int32) (*gltext.Font, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return gltext.LoadTruetype(fd, scale, 32, 127, gltext.LeftToRight)
}
