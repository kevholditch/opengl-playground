package render

import (
	"github.com/go-gl/gl/v2.1/gl"
)

type VertexArray struct {
	handle uint32
}

func NewVertexArray() *VertexArray {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	return &VertexArray{handle: vao}
}

func (v *VertexArray) AddBuffer(vb *VertexBuffer, layout *VertexBufferLayout) {
	vb.Bind()

	offset := 0
	for i := uint32(0); i < uint32(len(layout.elements)); i++ {
		gl.EnableVertexAttribArray(i)
		gl.VertexAttribPointer(i, layout.elements[i].getCount(), gl.FLOAT, false, layout.getStride(), gl.PtrOffset(offset))
		offset += int(layout.elements[i].getSize())
	}

}

func (v *VertexArray) Bind() {
	gl.BindVertexArray(v.handle)
}

func (v *VertexArray) UnBind() {
	gl.BindVertexArray(0)
}
