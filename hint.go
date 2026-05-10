package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

func SetHint(name, value string) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	result := C.SDL_SetHint(_name, _value)
	C.free(unsafe.Pointer(_value))
	C.free(unsafe.Pointer(_name))
	return result == C.SDL_TRUE
}
