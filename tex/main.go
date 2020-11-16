package main

import (
	"fmt"
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
		Title:        "Texture Demo - Kevin Holditch",
		SwapInterval: 1,
	})

	if err != nil {
		panic(err)
	}

	render.UseDefaultBlending()

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

	x := float32(0)
	y := float32(0)

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

	increment := float32(5)
	w.OnKeyPress(func(key int) {
		fmt.Printf("%v\n", key)
		switch key {
			case 70: x += increment
			case 65: x -= increment
			case 83: y -= increment
			case 68: y += increment
		}
	})

	for !w.ShouldClose() {

		render.Clear()

		program.Bind()
		m := mgl32.Ident4().Mul4(mgl32.Translate3D(x, y, 0))
		mvp := proj.Mul4(m)
		program.SetUniformMat4f("u_MVP", mvp)

		render.Render(va, ib, program)

		w.SwapBuffers()
		glfw.PollEvents()
	}
}
