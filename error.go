package sdl

// #include <SDL2/SDL.h>
//
// int setError(char *s) {
//   return SDL_SetError("%s", s);
// }
import "C"
import "fmt"

func ClearError() {
	C.SDL_ClearError()
}

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

func SetError(format string, a ...interface{}) {
	C.setError(C.CString(fmt.Sprintf(format, a...)))
}
