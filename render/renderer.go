package render

import (
	"github.com/go-gl/gl/v2.1/gl"
)

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Render(va *VertexArray, ib *IndexBuffer, shader *Program) {
	va.Bind()
	ib.Bind()
	shader.Bind()

	gl.DrawElements(gl.TRIANGLES, ib.count, gl.UNSIGNED_INT, gl.PtrOffset(0))
}
