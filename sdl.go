package sdl

// #cgo windows LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
import "C"

type Destroyable interface {
	Destroy()
}
