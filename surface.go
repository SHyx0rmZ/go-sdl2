package sdl

// #include <SDL2/SDL.h>
import "C"
import (
	"unsafe"

	"github.com/SHyx0rmZ/go-sdl2/internal"
)

type Surface C.struct_SDL_Surface

func (s *Surface) FillRect(rect *Rect, color uint32) error {
	r := rect.toInternal()
	if C.SDL_FillRect((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(r)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) FillRects(rects []*Rect, color uint32) error {
	r := make([]*internal.Rect, len(rects))
	for i, rect := range rects {
		r[i] = rect.toInternal()
	}
	if C.SDL_FillRects((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(&r[0])), C.int(len(r)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) Free() {
	if s != nil {
		C.SDL_FreeSurface((*C.struct_SDL_Surface)(s))
	}
}

func (s *Surface) BlitSurface(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	srcR := srcRect.toInternal()
	dstR := dstRect.toInternal()
	r := int(C.SDL_BlitSurface((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(srcR)), (*C.struct_SDL_Surface)(dst), (*C.struct_SDL_Rect)(unsafe.Pointer(dstR))))
	if r != 0 {
		return GetError()
	}
	return nil
}
