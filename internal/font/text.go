package font

import (
	"bytes"
	"io"
	"os"

	"github.com/go-gl/gltext"
)

// Load loads the specified font at the given scale.
func LoadFromFile(file string, scale int32) (*gltext.Font, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return Load(fd, scale)
}

func Load(reader io.Reader, scale int32) (*gltext.Font, error) {
	return gltext.LoadTruetype(reader, scale, 32, 127, gltext.LeftToRight)
}

func LoadDefault(scale int32) (*gltext.Font, error) {
	fontBytes, err := defaultFont()
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(fontBytes)
	return Load(reader, scale)
}
