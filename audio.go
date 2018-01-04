package sdl

// #include <SDL2/SDL.h>
import "C"

func GetAudioDeviceName(device int, isCapture bool) (string, error) {
	var capture int
	if isCapture {
		capture = 1
	}
	ptr := C.SDL_GetAudioDeviceName(C.int(device), C.int(capture))
	if ptr == nil {
		return "", GetError()
	}
	return C.GoString(ptr), nil
}
