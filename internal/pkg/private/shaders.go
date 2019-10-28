package private

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"strings"
)

func compileShader(source []byte, shaderType uint32) (uint32, error) {

	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(string(source) + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		cLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(cLog))

		return 0, fmt.Errorf("failed to compile %v: %v", source, cLog)
	}

	return shader, nil
}

func getShader(location string, shaderType uint32) uint32 {
	shaderSource, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err)
	}

	shader, err := compileShader(shaderSource, shaderType)
	if err != nil {
		panic(err)
	}

	return shader
}
