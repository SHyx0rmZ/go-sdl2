package sdl

// #cgo windows LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
import "C"
import (
	"unsafe"

	"github.com/SHyx0rmZ/go-sdl2/internal"
)

type Window C.struct_SDL_Window

type WindowFlags uint32

const (
	WindowFullscreen        WindowFlags = C.SDL_WINDOW_FULLSCREEN
	WindowFullscreenDesktop WindowFlags = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	WindowOpenGL            WindowFlags = C.SDL_WINDOW_OPENGL
	WindowShown             WindowFlags = C.SDL_WINDOW_SHOWN
	WindowHidden            WindowFlags = C.SDL_WINDOW_HIDDEN
	WindowBorderless        WindowFlags = C.SDL_WINDOW_BORDERLESS
	WindowResizable         WindowFlags = C.SDL_WINDOW_RESIZABLE
	WindowMinimized         WindowFlags = C.SDL_WINDOW_MINIMIZED
	WindowMaximized         WindowFlags = C.SDL_WINDOW_MAXIMIZED
	WindowInputGrabbed      WindowFlags = C.SDL_WINDOW_INPUT_GRABBED
	WindowInputFocus        WindowFlags = C.SDL_WINDOW_INPUT_FOCUS
	WindowMouseFocs         WindowFlags = C.SDL_WINDOW_MOUSE_FOCUS
	WindowForeign           WindowFlags = C.SDL_WINDOW_FOREIGN
	WindowAllowHighDPI      WindowFlags = C.SDL_WINDOW_ALLOW_HIGHDPI
	WindowMouseCapture      WindowFlags = C.SDL_WINDOW_MOUSE_CAPTURE
	WindowAlwaysOnTop       WindowFlags = C.SDL_WINDOW_ALWAYS_ON_TOP
	WindowSkipTaskbar       WindowFlags = C.SDL_WINDOW_SKIP_TASKBAR
	WindowUtility           WindowFlags = C.SDL_WINDOW_UTILITY
	WindowTooltip           WindowFlags = C.SDL_WINDOW_TOOLTIP
	WindowPopupMenu         WindowFlags = C.SDL_WINDOW_POPUP_MENU
)

const (
	WindowPositionCenteredMask  int = 0x2fff0000
	WindowPositionCentered      int = WindowPositionCenteredMask
	WindowPositionUndefinedMask int = 0x1fff0000
	WindowPositionUndefined     int = WindowPositionUndefinedMask
)

func WindowPositionCenteredDisplay(displayIndex int) int {
	return WindowPositionCenteredMask | displayIndex
}

func WindowPositionUndefinedDisplay(displayIndex int) int {
	return WindowPositionUndefinedMask | displayIndex
}

func CreateWindow(title string, x, y, w, h int, flags WindowFlags) (*Window, error) {
	nativeTitle := C.CString(title)
	defer C.free(unsafe.Pointer(nativeTitle))

	nativeWindow := C.SDL_CreateWindow(nativeTitle, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if nativeWindow == nil {
		return nil, GetError()
	}
	return (*Window)(nativeWindow), nil
}

func (w *Window) Destroy() {
	if w != nil {
		C.SDL_DestroyWindow((*C.struct_SDL_Window)(w))
	}
}

func GetDisplayBounds(displayIndex int, rect *Rect) error {
	var r internal.Rect

	if C.SDL_GetDisplayBounds(C.int(displayIndex), (*C.struct_SDL_Rect)(unsafe.Pointer(&r))) != 0 {
		return GetError()
	}

	rect.fromInternal(r)

	return nil
}
