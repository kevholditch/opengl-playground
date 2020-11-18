package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/opengl-playground/render"
	"runtime"
)

const (
	width, height = 1024, 768
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	cleanUp := render.Initialise()
	defer cleanUp()

	w, err := render.NewWindow(render.Config{
		MajorVersion: 3,
		MinorVersion: 2,
		Width:        width,
		Height:       height,
		Title:        "Batch Rendering Demo - Kevin Holditch",
		SwapInterval: 1,
	})

	if err != nil {
		panic(err)
	}

	render.UseDefaultBlending()

	indices := []int32{
		0, 1, 2,
		0, 3, 2,
		4, 5, 6,
		4, 7, 6,
	}

	buffer := []float32{
		200, 200, 0.4, 0.3, 0.2, 1.0,
		500, 200, 0.4, 0.3, 0.2, 1.0,
		500, 500, 0.4, 0.3, 0.2, 1.0,
		200, 500, 0.4, 0.3, 0.2, 1.0,
		600, 200, 0.8, 0.2, 0.2, 1.0,
		900, 200, 0.2, 0.8, 0.2, 1.0,
		900, 500, 0.8, 0.2, 0.2, 1.0,
		600, 500, 0.2, 0.8, 0.2, 1.0,
	}

	va := render.NewVertexArray()
	ib := render.NewIndexBuffer(indices)

	proj := mgl32.Ortho(0, width, 0, height, -1.0, 1.0)

	va.AddBuffer(render.NewVertexBuffer(buffer), render.NewVertexBufferLayout().AddLayout(2).AddLayout(4))

	vs, err := render.NewShaderFromFile("./batchrendering/vertex.shader", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fs, err := render.NewShaderFromFile("./batchrendering/fragment.shader", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program, err := render.NewProgram(vs, fs)
	if err != nil {
		panic(err)
	}

	program.SetUniformMat4f("u_MVP", proj)

	va.UnBind()
	ib.UnBind()
	program.UnBind()


	for !w.ShouldClose() {

		render.Clear()

		render.CheckErrors()

		program.Bind()
		program.SetUniformMat4f("u_MVP", proj)

		render.Render(va, ib, program)

		w.SwapBuffers()
		glfw.PollEvents()
	}
}
