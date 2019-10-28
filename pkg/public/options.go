package public

type InitOptions struct {
	Resizeable bool
	Width      int
	Height     int
	Title      string
}

type EngineOptions struct {
	Fps                    int
	VertexShaderLocation   string
	FragmentShaderLocation string
}
