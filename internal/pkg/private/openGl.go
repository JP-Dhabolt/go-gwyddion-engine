package private

import (
	"dev.azure.com/gwyddiongames/_git/go-gwyddion-engine.git/pkg/public"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

type openGlInfo struct {
	Prog uint32
	Loc  int32
}

func initGlfw(options public.InitOptions) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	var resizable int

	if options.Resizeable {
		resizable = glfw.True
	} else {
		resizable = glfw.False
	}

	glfw.WindowHint(glfw.Resizable, resizable)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(options.Width, options.Height, options.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(onKey)
	window.SetSizeCallback(onResize)

	return window
}

func initOpenGl(options public.EngineOptions) openGlInfo {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader := getShader(options.VertexShaderLocation, gl.VERTEX_SHADER)
	fragmentShader := getShader(options.FragmentShaderLocation, gl.FRAGMENT_SHADER)

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	myLoc := gl.GetUniformLocation(prog, gl.Str("color\x00"))

	return openGlInfo{prog, myLoc}
}

var onResize glfw.SizeCallback = func(window *glfw.Window, width int, height int) {
	if width < 1 {
		width = 1
	}

	if height < 1 {
		height = 1
	}

	gl.Viewport(0, 0, int32(width), int32(height))
}

var onKey glfw.KeyCallback = func(window *glfw.Window, key glfw.Key, _ int, _ glfw.Action, _ glfw.ModifierKey) {
	if key == glfw.KeyEscape {
		window.SetShouldClose(true)
	}
}
