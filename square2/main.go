package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

const (
	width, height = 800, 600
	sizeOfFloat32 = 4
	sizeOfInt32   = 4
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

type getObjIv func(uint32, uint32, *int32)
type getObjInfoLog func(uint32, int32, *int32, *uint8)

func getGlError(glHandle uint32, checkTrueParam uint32, getObjIvFn getObjIv,
	getObjInfoLogFn getObjInfoLog, failMsg string) error {

	var success int32
	getObjIvFn(glHandle, checkTrueParam, &success)

	if success == gl.FALSE {
		var logLength int32
		getObjIvFn(glHandle, gl.INFO_LOG_LENGTH, &logLength)

		log := gl.Str(strings.Repeat("\x00", int(logLength)))
		getObjInfoLogFn(glHandle, logLength, nil, log)

		return fmt.Errorf("%s: %s", failMsg, gl.GoStr(log))
	}

	return nil
}

func shaderFromFile(file string, sType uint32) (*uint32, error) {
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
	return &handle, nil
}

func compileShader(src string, sType uint32) uint32 {
	handle := gl.CreateShader(sType)
	glSrcs, freeFn := gl.Strs(src + "\x00")
	defer freeFn()
	gl.ShaderSource(handle, 1, glSrcs, nil)
	gl.CompileShader(handle)

	err := getGlError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::")
	if err != nil {
		panic(err)
	}

	return handle
}

func createShader(vertexShader, fragmentShader string) uint32 {
	program := gl.CreateProgram()
	vs, err := shaderFromFile(vertexShader, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fs, err := shaderFromFile(fragmentShader, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	gl.AttachShader(program, *vs)
	gl.AttachShader(program, *fs)

	gl.LinkProgram(program)
	gl.ValidateProgram(program)

	gl.DeleteShader(*vs)
	gl.DeleteShader(*fs)

	return program
}

func checkErrors() {
	for {
		e := gl.GetError()
		if e == gl.NO_ERROR {
			break
		}
		fmt.Printf("error is :%v\n", e)
	}
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	window, err := glfw.CreateWindow(width, height, "Kevin - Demo", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	indices := []int32{
		0, 1, 2,
		0, 3, 2,
	}

	positions := []float32{
		-0.5, -0.5,
		0.5, -0.5,
		0.5, 0.5,
		-0.5, 0.5,
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions)*sizeOfFloat32, gl.Ptr(positions), gl.STATIC_DRAW)

	var ibo uint32
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*sizeOfInt32, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, sizeOfFloat32*2, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	shader := createShader("./square2/vertex.shader", "./square2/fragment.shader")

	gl.UseProgram(shader)

	location := gl.GetUniformLocation(shader, gl.Str("u_Color"+"\x00"))

	r := float32(0.0)
	increment := float32(0.05)

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Uniform4f(location, r, 0.3, 0.1, 1.0)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))

		if r > 1.0 {
			increment = -0.05
		} else if r < 0.0 {
			increment = 0.05
		}
		r += increment

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
