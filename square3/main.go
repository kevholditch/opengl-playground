package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/kevholditch/opengl-playground/render"
	"log"
	"runtime"
)

const (
	width, height = 800, 600
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
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

	va := render.NewVertexArray()
	ib := render.NewIndexBuffer(indices)

	va.AddBuffer(render.NewVertexBuffer(positions), render.NewVertexBufferLayout().AddLayoutFloats(2))

	vs, err := render.NewShaderFromFile("./square3/vertex.shader", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fs, err := render.NewShaderFromFile("./square3/fragment.shader", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program, err := render.NewProgram(vs, fs)
	if err != nil {
		panic(err)
	}

	va.UnBind()
	ib.UnBind()
	program.UnBind()

	r := float32(0.0)
	increment := float32(0.05)

	for !window.ShouldClose() {

		render.Clear()

		program.Bind()
		program.SetUniformVec4("u_Color", r, 0.3, 0.8, 1.0)

		render.Render(va, ib, program)

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
