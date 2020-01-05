package sdl

// #include <SDL2/SDL.h>
import "C"
import (
	"image"
	"image/color"
	"unsafe"

	"code.witches.io/go/sdl2/internal"
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

func (s *Surface) AlphaMod() (a uint8, err error) {
	if C.SDL_GetSurfaceAlphaMod((*C.struct_SDL_Surface)(s), (*C.Uint8)(unsafe.Pointer(&a))) != 0 {
		return 0, GetError()
	}
	return a, nil
}

func (s *Surface) ColorMod() (r, g, b uint8, err error) {
	if C.SDL_GetSurfaceColorMod((*C.struct_SDL_Surface)(s), (*C.Uint8)(unsafe.Pointer(&r)), (*C.Uint8)(unsafe.Pointer(&g)), (*C.Uint8)(unsafe.Pointer(&b))) != 0 {
		return 0, 0, 0, GetError()
	}
	return r, g, b, nil
}

func (s *Surface) SetAlphaMod(a uint8) error {
	if C.SDL_SetSurfaceAlphaMod((*C.struct_SDL_Surface)(s), C.Uint8(a)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) SetColorMod(r, g, b uint8) error {
	if C.SDL_SetSurfaceColorMod((*C.struct_SDL_Surface)(s), C.Uint8(r), C.Uint8(g), C.Uint8(b)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) ClipRect() Rect {
	var r internal.Rect
	C.SDL_GetClipRect((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(&r)))
	var rect Rect
	rect.fromInternal(r)
	return rect
}

func (s *Surface) SetClipRect(r *Rect) bool {
	return C.SDL_SetClipRect((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(r.toInternal()))) == C.SDL_TRUE
}

func (s *Surface) Lock() error {
	if C.SDL_LockSurface((*C.struct_SDL_Surface)(s)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) Unlock() {
	C.SDL_UnlockSurface((*C.struct_SDL_Surface)(s))
}

func (s *Surface) ColorModel() color.Model {
	return color.RGBAModel
}

func (s *Surface) Bounds() image.Rectangle {
	r := s.ClipRect()
	return image.Rect(r.X, r.Y, r.X+r.W, r.Y+r.H)
}

func (s *Surface) At(x, y int) color.Color {
	_surface := (*C.struct_SDL_Surface)(s)
	_format := (*PixelFormatS)(unsafe.Pointer(_surface.format))
	_pixel := *(*uint32)(unsafe.Pointer(uintptr(_surface.pixels) + uintptr(y)*uintptr(_surface.pitch) + uintptr(x)*uintptr(_format.bytesPerPixel)))
	r := uint8(((_pixel & _format.Rmask) >> _format.rShift) << _format.rLoss)
	g := uint8(((_pixel & _format.Gmask) >> _format.gShift) << _format.gLoss)
	b := uint8(((_pixel & _format.Bmask) >> _format.bShift) << _format.bLoss)
	a := uint8(((_pixel & _format.Amask) >> _format.aShift) << _format.aLoss)
	return color.RGBA{
		R: uint8(uint16(r) * uint16(a) / 255),
		G: uint8(uint16(g) * uint16(a) / 255),
		B: uint8(uint16(b) * uint16(a) / 255),
		A: a,
	}
}

func (s Surface) Format() PixelFormatS {
	return *(*PixelFormatS)(unsafe.Pointer(s.format))
}

func (s *Surface) Pixels() unsafe.Pointer {
	return unsafe.Pointer(s.pixels)
}
