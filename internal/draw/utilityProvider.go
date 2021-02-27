package draw

import (
	"github.com/go-gl/gl/v2.1/gl"
)

type UtilityProvider struct{}
type Drawable uint32

func (provider UtilityProvider) CreateDrawable(points []float32) Drawable {
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

	return Drawable(vao)
}

func Triangles(drawable Drawable, length int32) {
	gl.BindVertexArray(uint32(drawable))
	gl.DrawArrays(gl.TRIANGLES, 0, length)
}
