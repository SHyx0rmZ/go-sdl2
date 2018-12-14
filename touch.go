package sdl

// #include <SDL2/SDL.h>
import "C"

type TouchID int64
type FingerID int64

type Finger struct {
	ID FingerID
	X float32
	Y float32
	Pressure float32
}

func GetNumTouchDevices() int {
	return int(C.SDL_GetNumTouchDevices())
}
