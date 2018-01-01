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
