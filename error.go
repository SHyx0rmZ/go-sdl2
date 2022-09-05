package sdl

// #include <SDL2/SDL_error.h>
//
// int setError(char *s) {
//   return SDL_SetError("%s", s);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

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
	nativeError := C.CString(fmt.Sprintf(format, a...))
	defer C.free(unsafe.Pointer(nativeError))

	C.setError(nativeError)
}
