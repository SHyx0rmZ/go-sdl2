package sdl

// #include <SDL2/SDL_gesture.h>
import "C"
import (
	"unsafe"
)

type GestureID int64

func RecordGesture(touchID TouchID) bool {
	return C.SDL_RecordGesture(C.Sint64(touchID)) == 0
}

func LoadDollarTemplates(touchID TouchID, src *RWOps) (int, error) {
	n := int(C.SDL_LoadDollarTemplates(C.Sint64(touchID), (*C.struct_SDL_RWops)(unsafe.Pointer(src))))
	if n < 0 {
		return 0, GetError()
	}
	return n, nil
}

func SaveAllDollarTemplates(dst *RWOps) (int, error) {
	n := int(C.SDL_SaveAllDollarTemplates((*C.struct_SDL_RWops)(unsafe.Pointer(dst))))
	if n < 0 {
		return 0, GetError()
	}
	return n, nil
}

func SaveDollarTemplate(gestureID GestureID, dst *RWOps) error {
	if C.SDL_SaveDollarTemplate(C.Sint64(gestureID), (*C.struct_SDL_RWops)(unsafe.Pointer(dst))) != 1 {
		return GetError()
	}
	return nil
}
