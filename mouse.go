package sdl

// #include <SDL2/SDL.h>
import "C"

const (
	MouseButtonLeft   = 1
	MouseButtonMiddle = 2
	MouseButtonRight  = 3
	MouseButtonExtra1 = 4
	MouseButtonExtra2 = 5
)

type MouseButtons struct {
	Left   bool
	Middle bool
	Right  bool
	X1     bool
	X2     bool
}

func (b MouseButtons) String() string {
	s := []byte{'-', '-', '-', '-', '-'}
	if b.Left {
		s[0] = 'L'
	}
	if b.Middle {
		s[1] = 'M'
	}
	if b.Right {
		s[2] = 'R'
	}
	if b.X1 {
		s[3] = 'X'
	}
	if b.X2 {
		s[4] = 'X'
	}
	return string(s)
}

func GetMouseState(x, y *int) uint32 {
	var _x, _y int32
	buttons := uint32(C.SDL_GetMouseState((*C.int)(&_x), (*C.int)(&_y)))
	*x = int(_x)
	*y = int(_y)
	return buttons
	//return uint32(C.SDL_GetMouseState((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y))))
}
