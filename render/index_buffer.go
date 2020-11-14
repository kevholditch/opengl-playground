package render

import (
	"github.com/go-gl/gl/v2.1/gl"
)

const (
	sizeOfFloat32 = 4
	sizeOfInt32   = 4
)

type IndexBuffer struct {
	handle uint32
	count  int32
}

func NewIndexBuffer(indices []int32) *IndexBuffer {
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*sizeOfInt32, gl.Ptr(indices), gl.STATIC_DRAW)

	return &IndexBuffer{handle: ibo, count: int32(len(indices))}
}

func (ib *IndexBuffer) GetCount() int32 {
	return ib.count
}

func (ib *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.handle)
}

func (ib *IndexBuffer) UnBind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
