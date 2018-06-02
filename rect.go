package sdl

import "github.com/SHyx0rmZ/go-sdl2/internal"

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
