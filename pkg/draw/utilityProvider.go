package draw

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Drawable is a reference to an OpenGL registered Vertex Array Object
type Drawable uint32

/*
CreateDrawable registers an array of points with OpenGL and returns a reference to a drawable

points is an array of float32 numbers that represent the x, y, and z positions of a desired shape.
Currently these points are expected to make up a series of triangles to draw to the screen, requiring the caller
to have appropriately massaged them.  This interface will likely break in an upcoming release, as this massaging
will be handled by the engine itself
*/
func CreateDrawable(points []float32) Drawable {
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

/*
WriteDrawable takes a Drawable created with CreateDrawable and writes it to the screen
*/
func WriteDrawable(drawable Drawable, length int32) {
	gl.BindVertexArray(uint32(drawable))
	gl.DrawArrays(gl.TRIANGLES, 0, length)
}
