package sdl

// #include <SDL2/SDL_joystick.h>
import "C"
import (
	"unsafe"
)

type JoystickID int32

type Joystick = C.SDL_Joystick

func JoystickOpen(deviceIndex int) (*Joystick, error) {
	ptr := C.SDL_JoystickOpen(C.int(deviceIndex))
	if ptr == nil {
		return nil, GetError()
	}
	return (*Joystick)(unsafe.Pointer(ptr)), nil
}

func JoystickClose(joystick *Joystick) {
	C.SDL_JoystickClose(joystick)
}

func JoystickInstanceID(joystick *Joystick) (JoystickID, error) {
	id := C.SDL_JoystickInstanceID(joystick)
	if id < 0 {
		return 0, GetError()
	}
	return JoystickID(id), nil
}
