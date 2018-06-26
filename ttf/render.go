package ttf

// #include <SDL2/SDL_ttf.h>
import "C"
import (
	"github.com/SHyx0rmZ/go-sdl2"
	"unsafe"
)

func RenderUTF8Blended(font *Font, text string, fg sdl.Color) (*sdl.Surface, error) {
	nativeText := C.CString(text)
	defer C.free(unsafe.Pointer(nativeText))

	color := C.struct_SDL_Color{
		r: C.Uint8(fg.R),
		g: C.Uint8(fg.G),
		b: C.Uint8(fg.B),
		a: C.Uint8(fg.A),
	}

	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended((*C.struct__TTF_Font)(font), nativeText, color)))
	if surface == nil {
		return nil, GetError()
	}

	return surface, nil
}
