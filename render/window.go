package render

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Config struct {
	MajorVersion int
	MinorVersion int
	Width int
	Height int
	Title string
}

type Window struct {
	handle *glfw.Window
}


func NewWindow(cfg Config) (*Window, error){

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, cfg.MajorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, cfg.MinorVersion)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	window, err := glfw.CreateWindow(cfg.Width, cfg.Height, cfg.Title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	return &Window{handle: window}, nil
}

func (w *Window) ShouldClose() bool {
	return w.handle.ShouldClose()
}


func (w *Window) SwapBuffers() {
	w.handle.SwapBuffers()
}

