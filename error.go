package sdl

// #cgo windows LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
import "C"
import "fmt"

func GetError() error {
	ptr := C.SDL_GetError()
	if ptr == nil {
		return nil
	}
	err := C.GoString(ptr)
	if len(err) == 0 {
		return nil
	}
	return fmt.Errorf("sdl: %s", err)
}
