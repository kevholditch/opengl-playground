package graf

import (
	"github.com/go-gl/gl/v2.1/gl"
	"io/ioutil"
)

type Shader struct {
	Handle uint32
	Type   uint32
}

func NewShaderFromFile(file string, sType uint32) (*Shader, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	handle := gl.CreateShader(sType)
	glSrc, freeFn := gl.Strs(string(src) + "\x00")
	defer freeFn()
	gl.ShaderSource(handle, 1, glSrc, nil)
	gl.CompileShader(handle)
	err = getGlError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::"+file)
	if err != nil {
		return nil, err
	}
	return &Shader{Handle: handle, Type: sType}, nil
}

type Program struct {
	Handle uint32
}

func NewProgram(shaders ...*Shader) (*Program, error) {
	handle := gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(handle, shader.Handle)
	}

	gl.LinkProgram(handle)
	gl.ValidateProgram(handle)

	for _, shader := range shaders {
		gl.DeleteShader(shader.Handle)
	}

	return &Program{Handle: handle}, nil
}

func (p *Program) Bind() {
	gl.UseProgram(p.Handle)
}

func (p *Program) UnBind() {
	gl.UseProgram(0)
}

func (p *Program) getUniformLocation(name string) int32 {
	return gl.GetUniformLocation(p.Handle, gl.Str(name+"\x00"))
}

func (p *Program) SetUniformValue(name string, v1, v2, v3, v4 float32) {
	gl.Uniform4f(p.getUniformLocation(name), v1, v2, v3, v4)
}
