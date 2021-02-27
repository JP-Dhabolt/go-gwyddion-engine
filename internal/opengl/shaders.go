package opengl

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
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

func getShaderFromFile(location string, shaderType uint32) uint32 {
	shaderSource, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err)
	}
	return getShader(shaderSource, shaderType)
}

func getShader(shaderSource []byte, shaderType uint32) uint32 {
	shader, err := compileShader(shaderSource, shaderType)
	if err != nil {
		panic(err)
	}

	return shader
}

func getDefaultFragmentShader() uint32 {
	return getShader([]byte(`
#version 140

uniform vec4 color;

out vec4 frag_colour;

void main() {
    frag_colour = color;
}
	`), gl.FRAGMENT_SHADER)
}

func getDefaultVertexShader() uint32 {
	return getShader([]byte(`
#version 140
in vec3 position;
in vec4 color;

void main() {
    gl_Position = vec4(position, 1.0);
}
	`), gl.VERTEX_SHADER)
}
