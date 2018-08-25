package sdl

// #include <SDL2/SDL.h>
import "C"
import (
	"github.com/SHyx0rmZ/go-sdl2/internal"
	"unsafe"
)

type RendererFlags uint32

const (
	RendererSoftware RendererFlags = 1 << iota
	RendererAccelerated
	RendererPresentVSync
	RendererTargetTexture
)

type TextureAccess uint32

const (
	TextureAccessStatic TextureAccess = iota
	TextureAccessStreaming
	TextureAccessTarget
)

type TextureModulate uint32

const (
	TextureModulateNone TextureModulate = iota
	TextureModulateColor
	TextureModulateAlpha
)

type RendererFlip uint32

const (
	FlipNone RendererFlip = iota
	FlipHorizontal
	FlipVertical
)

type Renderer C.struct_SDL_Renderer

func CreateRenderer(window *Window, index int, flags RendererFlags) (*Renderer, error) {
	nativeRenderer := C.SDL_CreateRenderer((*C.struct_SDL_Window)(window), C.int(index), C.Uint32(flags))
	if nativeRenderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(nativeRenderer), nil
}

func (r *Renderer) Destroy() {
	C.SDL_DestroyRenderer((*C.struct_SDL_Renderer)(r))
}

func (r *Renderer) GetDrawColor() (red, green, blue, alpha uint8, err error) {
	if C.SDL_GetRenderDrawColor((*C.struct_SDL_Renderer)(r), (*C.Uint8)(unsafe.Pointer(&red)), (*C.Uint8)(unsafe.Pointer(&green)), (*C.Uint8)(unsafe.Pointer(&blue)), (*C.Uint8)(unsafe.Pointer(&alpha))) != 0 {
		return 0, 0, 0, 0, GetError()
	}
	return red, green, blue, alpha, nil
}

func (r *Renderer) Clear() error {
	if C.SDL_RenderClear((*C.struct_SDL_Renderer)(r)) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) Copy(texture *Texture, srcRect, dstRect *Rect) error {
	srcR := srcRect.toInternal()
	dstR := dstRect.toInternal()
	if C.SDL_RenderCopy((*C.struct_SDL_Renderer)(r), (*C.struct_SDL_Texture)(texture), (*C.struct_SDL_Rect)(unsafe.Pointer(srcR)), (*C.struct_SDL_Rect)(unsafe.Pointer(dstR))) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) CopyEx(texture *Texture, srcRect, dstRect *Rect, angle float64, center *Point, flip RendererFlip) error {
	nativeRenderer := (*C.struct_SDL_Renderer)(r)
	nativeTexture := (*C.struct_SDL_Texture)(texture)
	nativeSrcRect := (*C.struct_SDL_Rect)(unsafe.Pointer(srcRect.toInternal()))
	nativeDstRect := (*C.struct_SDL_Rect)(unsafe.Pointer(dstRect.toInternal()))
	nativeAngle := C.double(angle)
	nativeCenter := (*C.struct_SDL_Point)(unsafe.Pointer(center.toInternal()))
	nativeFlip := (C.SDL_RendererFlip)(flip)
	if C.SDL_RenderCopyEx(nativeRenderer, nativeTexture, nativeSrcRect, nativeDstRect, nativeAngle, nativeCenter, nativeFlip) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) DrawLine(x1, y1, x2, y2 int) error {
	if C.SDL_RenderDrawLine((*C.struct_SDL_Renderer)(r), C.int(x1), C.int(y1), C.int(x2), C.int(y2)) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) DrawLines(points []*Point) error {
	p := make([]*internal.Point, len(points))
	for i, point := range points {
		p[i] = point.toInternal()
	}
	if C.SDL_RenderDrawLines((*C.struct_SDL_Renderer)(r), (*C.struct_SDL_Point)(unsafe.Pointer(&p[0])), C.int(len(p))) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) DrawPoint(x, y int) error {
	if C.SDL_RenderDrawPoint((*C.struct_SDL_Renderer)(r), C.int(x), C.int(y)) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) DrawPoints(points []*Point) error {
	p := make([]*internal.Point, len(points))
	for i, point := range points {
		p[i] = point.toInternal()
	}
	if C.SDL_RenderDrawPoints((*C.struct_SDL_Renderer)(r), (*C.struct_SDL_Point)(unsafe.Pointer(&p[0])), C.int(len(p))) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) DrawRect(rect *Rect) error {
	if C.SDL_RenderDrawRect((*C.struct_SDL_Renderer)(r), (*C.struct_SDL_Rect)(unsafe.Pointer(rect.toInternal()))) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) DrawRects(rects []*Rect) error {
	rs := make([]*internal.Rect, len(rects))
	for i, rect := range rects {
		rs[i] = rect.toInternal()
	}
	if C.SDL_RenderDrawRects((*C.struct_SDL_Renderer)(r), (*C.struct_SDL_Rect)(unsafe.Pointer(&rs[0])), C.int(len(rs))) != 0 {
		return GetError()
	}
	return nil
}

func (r *Renderer) Present() {
	C.SDL_RenderPresent((*C.struct_SDL_Renderer)(r))
}

func (r *Renderer) SetDrawColor(red, green, blue, alpha uint8) error {
	if C.SDL_SetRenderDrawColor((*C.struct_SDL_Renderer)(r), C.Uint8(red), C.Uint8(green), C.Uint8(blue), C.Uint8(alpha)) != 0 {
		return GetError()
	}
	return nil
}

type Texture C.struct_SDL_Texture

func (r *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	nativeTexture := C.SDL_CreateTextureFromSurface((*C.struct_SDL_Renderer)(r), (*C.struct_SDL_Surface)(surface))
	if nativeTexture != nil {
		return nil, GetError()
	}
	return (*Texture)(nativeTexture), nil
}

func (t *Texture) Destroy() {
	C.SDL_DestroyTexture((*C.struct_SDL_Texture)(t))
}
