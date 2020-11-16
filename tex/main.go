package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/opengl-playground/render"
	"log"
	"runtime"
)

const (
	width, height = 1024, 768
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
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

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	indices := []int32{
		0, 1, 2,
		0, 3, 2,
	}

	positions := []float32{
		200, 200, 0.0, 0.0,
		500, 200, 1.0, 0.0,
		500, 500, 1.0, 1.0,
		200, 500, 0.0, 1.0,
	}

	va := render.NewVertexArray()
	ib := render.NewIndexBuffer(indices)

	proj := mgl32.Ortho(0, width, 0, height, -1.0, 1.0)

	va.AddBuffer(render.NewVertexBuffer(positions), render.NewVertexBufferLayout().AddLayout(2).AddLayout(2))

	vs, err := render.NewShaderFromFile("./tex/vertex.shader", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fs, err := render.NewShaderFromFile("./tex/fragment.shader", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program, err := render.NewProgram(vs, fs)
	if err != nil {
		panic(err)
	}

	texture, err := render.NewTextureFromFile("./tex/form3.png")
	if err != nil {
		panic(err)
	}
	texture.Bind(0)
	program.SetUniformI1("u_Texture", 0)
	program.SetUniformMat4f("u_MVP", proj)

	va.UnBind()
	ib.UnBind()
	program.UnBind()

	r := float32(0.0)
	increment := float32(0.05)

	for !window.ShouldClose() {

		render.Clear()

		program.Bind()
		program.SetUniformMat4f("u_MVP", proj)

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
