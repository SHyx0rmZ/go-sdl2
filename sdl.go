package sdl

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: --static sdl2
// #include <SDL2/SDL.h>
import "C"

type Destroyable interface {
	Destroy()
}
