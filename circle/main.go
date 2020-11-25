package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/opengl-playground/render"
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"strings"
)

const width, height = 800, 600
const sizeOfFloat32 = 4

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

	if err := gl.Init(); err != nil {
		panic(err)
	}

	triangleAmount := float32(60)
	twicePi := float32(2.0) * math.Pi

	var positions []float32
	x := float32(200)
	y := float32(200)
	radius := float32(20)
	for i := float32(0); i <= triangleAmount; i++ {
		x1 := x + (radius * float32(math.Cos(float64(i*twicePi/triangleAmount))))
		y1 := y + (radius * float32(math.Sin(float64(i*twicePi/triangleAmount))))
		positions = append(positions, x1, y1)
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions)*sizeOfFloat32, gl.Ptr(positions), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, sizeOfFloat32*2, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	shader := createShader("./circle/vertex.shader", "./circle/fragment.shader")

	location := gl.GetUniformLocation(shader, gl.Str("u_MVP"+"\x00"))

	gl.UseProgram(shader)
	proj := mgl32.Ortho(0, width, 0, height, -1.0, 1.0)
	gl.UniformMatrix4fv(location, 1, false, &proj[0])


	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(positions)))
		render.CheckErrors()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}


func Circle(x, y, radius, r, g, b, a float32) {
	triangleAmount := float32(20)
	twicePi := float32(2.0) * math.Pi

	gl.Begin(gl.TRIANGLE_FAN)
	gl.Color4f(r, g, b, a)

	gl.Vertex2f(x, y)
	for i := float32(0); i <= triangleAmount; i++ {
		x1 := x + (radius * float32(math.Cos(float64(i*twicePi/triangleAmount))))
		y1 := y + (radius * float32(math.Sin(float64(i*twicePi/triangleAmount))))
		gl.Vertex2f(
			x1,
			y1)
	}
	gl.End()
}

func Circle2(x, y, radius, r, g, b, a float32) {
	triangleAmount := float32(20)
	twicePi := float32(2.0) * math.Pi

	gl.Begin(gl.TRIANGLE_FAN)
	gl.Color4f(r, g, b, a)

	gl.Vertex2f(x, y)
	for i := float32(0); i <= triangleAmount; i++ {
		x1 := x + (radius * float32(math.Cos(float64(i*twicePi/triangleAmount))))
		y1 := y + (radius * float32(math.Sin(float64(i*twicePi/triangleAmount))))
		gl.Vertex2f(
			x1,
			y1)
	}
	gl.End()
}

func DrawCircle( cx,  cy,  r float32,  num_segments int) {
	gl.Begin(gl.LINE_LOOP)
	for ii := 0; ii < num_segments; ii++ {

		theta := float32(2.0) * 3.1415926 * float32(ii) / float32(num_segments)

		x := r * float32(math.Cos(float64(theta))) //calculate the x component
		y := r * float32(math.Sin(float64(theta))) //calculate the y component

		gl.Vertex2f(x + cx, y + cy) //output vertex

	}
	gl.End()
}