package private

import (
	"github.com/JP-Dhabolt/go-gwyddion-engine/pkg/public"
	"github.com/go-gl/gl/v2.1/gl"
)

type utilityProvider struct {
	BLACK public.Color
	RED   public.Color
	BLUE  public.Color
	GREEN public.Color
}

func createUtilityProvider() utilityProvider {
	provider := utilityProvider{}
	provider.BLACK = public.Color{Alpha: 1}
	provider.RED = public.Color{Red: 1, Alpha: 1}
	provider.BLUE = public.Color{Blue: 1, Alpha: 1}
	provider.GREEN = public.Color{Green: 1, Alpha: 1}

	return provider
}

func (provider utilityProvider) CreateDrawable(points []float32) public.Drawable {
	var vbo uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return public.Drawable(vao)
}

func drawTriangles(drawable public.Drawable, length int32) {
	gl.BindVertexArray(uint32(drawable))
	gl.DrawArrays(gl.TRIANGLES, 0, length)
}
