package sdl

// #include <SDL2/SDL.h>
import "C"

type Surface C.struct_SDL_Surface

func (s *Surface) Free() {
	if s != nil {
		C.SDL_FreeSurface((*C.struct_SDL_Surface)(s))
	}
}
