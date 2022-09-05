package sdl

//#include <SDL2/SDL_rect.h>
import "C"
import (
	"unsafe"

	"code.witches.io/go/sdl2/internal"
)

type Rect struct {
	X, Y, W, H int
}

func (r *Rect) fromInternal(rect internal.Rect) {
	r.X = int(rect.X)
	r.Y = int(rect.Y)
	r.W = int(rect.W)
	r.H = int(rect.H)
}

func (r *Rect) toInternal() (rect *internal.Rect) {
	if r == nil {
		return nil
	}
	rect = new(internal.Rect)
	rect.X = int32(r.X)
	rect.Y = int32(r.Y)
	rect.W = int32(r.W)
	rect.H = int32(r.H)
	return rect
}

type Point struct {
	X, Y int
}

func (p *Point) fromInternal(point internal.Point) {
	p.X = int(point.X)
	p.Y = int(point.Y)
}

func (p *Point) toInternal() (point *internal.Point) {
	if p == nil {
		return nil
	}
	point = new(internal.Point)
	point.X = int32(p.X)
	point.Y = int32(p.Y)
	return point
}

func PointInRect(point Point, rect Rect) bool {
	return C.SDL_PointInRect((*C.struct_SDL_Point)(unsafe.Pointer(point.toInternal())), (*C.struct_SDL_Rect)(unsafe.Pointer(rect.toInternal()))) == C.SDL_TRUE
}

func HasIntersection(a, b Rect) bool {
	return C.SDL_HasIntersection((*C.struct_SDL_Rect)(unsafe.Pointer(a.toInternal())), (*C.struct_SDL_Rect)(unsafe.Pointer(b.toInternal()))) == C.SDL_TRUE
}

func IntersectRect(a, b Rect) (Rect, bool) {
	var internalRect internal.Rect
	r := C.SDL_IntersectRect((*C.struct_SDL_Rect)(unsafe.Pointer(a.toInternal())), (*C.struct_SDL_Rect)(unsafe.Pointer(b.toInternal())), (*C.struct_SDL_Rect)(unsafe.Pointer(&internalRect))) == C.SDL_TRUE
	if !r {
		return Rect{}, false
	}
	var rect Rect
	rect.fromInternal(internalRect)
	return rect, true
}
