package ttf

// #cgo windows LDFLAGS: -lSDL2_ttf -lSDL2
// #cgo linux freebsd darwin pkg-config: SDL2_ttf
// #include <SDL2/SDL_ttf.h>
import "C"
import (
	"fmt"
	"unsafe"

	"code.witches.io/go/sdl2"
)

type Destroyable interface {
	Destroy()
}

type Font C.struct_TTF_Font

func Init() error {
	if C.TTF_Init() == -1 {
		return GetError()
	}
	return nil
}

func WasInit() bool {
	result := C.TTF_WasInit()
	if result == 0 {
		return false
	}
	return true
}

func Quit() {
	C.TTF_Quit()
}

func OpenFont(file string, size int) (*Font, error) {
	nativeFile := C.CString(file)
	defer C.free(unsafe.Pointer(nativeFile))

	font := (*Font)(C.TTF_OpenFont(nativeFile, C.int(size)))
	if font == nil {
		return nil, GetError()
	}

	return font, nil
}

func OpenFontRW(src *sdl.RWOps, freeSrc bool, size int) (*Font, error) {
	var free int

	if freeSrc {
		free = 1
	}

	font := (*Font)(C.TTF_OpenFontRW((*C.struct_SDL_RWops)(unsafe.Pointer(src)), C.int(free), C.int(size)))
	if font == nil {
		return nil, GetError()
	}

	return font, nil
}

func OpenFontIndex(file string, size int, index int) (*Font, error) {
	nativeFile := C.CString(file)
	defer C.free(unsafe.Pointer(nativeFile))

	font := (*Font)(C.TTF_OpenFontIndex(nativeFile, C.int(size), C.long(index)))
	if font == nil {
		return nil, GetError()
	}

	return font, nil
}

func GetError() error {
	ptr := C.TTF_GetError()
	if ptr == nil {
		return nil
	}
	err := C.GoString(ptr)
	if len(err) == 0 {
		return nil
	}
	return fmt.Errorf("ttf: %s", err)
}

func (f *Font) Close() {
	C.TTF_CloseFont((*C.struct__TTF_Font)(f))
}
