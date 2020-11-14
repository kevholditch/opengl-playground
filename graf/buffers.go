package graf

import (
	"github.com/go-gl/gl/v2.1/gl"
)


const (
	sizeOfFloat32 = 4
	sizeOfInt32   = 4
)


type VertexArray struct {
	handle uint32
}

type IndexBuffer struct {
	handle uint32
}

type VertexBuffer struct {
	handle uint32
}

func NewVertexArray() *VertexArray {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	return &VertexArray{handle: vao}
}

func (v *VertexArray) Bind(){
	gl.BindVertexArray(v.handle)
}

func (v *VertexArray) UnBind(){
	gl.BindVertexArray(0)
}

func NewVertexBuffer(values []float32) *VertexBuffer{

	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(values)*sizeOfFloat32, gl.Ptr(values), gl.STATIC_DRAW)

	return &VertexBuffer{handle: buffer}
}

func (v *VertexBuffer) Bind()  {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.handle)
}

func (v *VertexBuffer) Unbind()  {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func NewIndexBuffer(indices [] int32) *IndexBuffer {
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*sizeOfInt32, gl.Ptr(indices), gl.STATIC_DRAW)

	return &IndexBuffer{handle: ibo}
}

func (ib *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.handle)
}

func (ib *IndexBuffer) UnBind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

