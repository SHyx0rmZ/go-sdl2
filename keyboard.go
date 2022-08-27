package sdl

// #include <SDL2/SDL.h>
import "C"
import (
	"unsafe"
)

type KeySymbol struct {
	Scancode  ScanCode
	Keycode   KeyCode
	Modifiers KeyModifiers
}

func GetKeyboardState() []uint8 {
	var numkeys int32
	ptr := C.SDL_GetKeyboardState((*C.int)(unsafe.Pointer(&numkeys)))
	ary := make([]uint8, numkeys)
	copy(ary, unsafe.Slice((*uint8)(unsafe.Pointer(ptr)), numkeys))
	return ary
}
