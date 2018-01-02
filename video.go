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
	WindowFullscreen WindowFlags = 1 << iota
	WindowOpenGL
	WindowShown
	WindowHidden
	WindowBorderless
	WindowResizable
	WindowMinimized
	WindowMaximized
	WindowInputGrabbed
	WindowInputFocus
	WindowMouseFocus
	WindowForeign
	WindowFullscreenDesktop WindowFlags = 1<<iota + 1
	WindowAllowHighDPI      WindowFlags = 1 << iota
	WindowMouseCapture
	WindowAlwaysOnTop
	WindowSkipTaskbar
	WindowUtility
	WindowTooltip
	WindowPopupMenu
	WindowVulkan WindowFlags = 1 << 28
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

type WindowEventID uint8

const (
	WindowEventNone WindowEventID = iota
	WindowEventShown
	WindowEventHidden
	WindowEventExposed
	WindowEventMoved
	WindowEventResized
	WindowEventSizeChanged
	WindowEventMinimized
	WindowEventMaximized
	WindowEventRestored
	WindowEventMouseEntered
	WindowEventMouseLeft
	WindowEventFocusGained
	WindowEventFocusLost
	WindowEventClose
	WindowEventTakeFocus
	WindowEventHitTest
)

var windowEventToStringMap = map[WindowEventID]string{
	WindowEventNone:         "none",
	WindowEventShown:        "shown",
	WindowEventHidden:       "hidden",
	WindowEventExposed:      "exposed",
	WindowEventMoved:        "moved",
	WindowEventResized:      "resized",
	WindowEventSizeChanged:  "size changed",
	WindowEventMinimized:    "minimized",
	WindowEventMaximized:    "maximized",
	WindowEventRestored:     "restored",
	WindowEventMouseEntered: "mouse entered",
	WindowEventMouseLeft:    "mouse left",
	WindowEventFocusGained:  "focus gained",
	WindowEventFocusLost:    "focus lost",
	WindowEventClose:        "close",
	WindowEventTakeFocus:    "take focus",
	WindowEventHitTest:      "hit test",
}

func (i WindowEventID) String() string {
	s, ok := windowEventToStringMap[i]
	if !ok {
		return "unknown window event id"
	}
	return s
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
