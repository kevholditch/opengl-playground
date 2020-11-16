package render

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)


func Initialise() func() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	return func() { glfw.Terminate()}
}