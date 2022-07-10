package sdl

// #include <SDL2/SDL.h>
// #include <SDL2/SDL_syswm.h>
import "C"
import (
	"unsafe"

	"code.witches.io/go/sdl2/internal"
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

func (w *Window) Surface() (*Surface, error) {
	nativeSurface := C.SDL_GetWindowSurface((*C.struct_SDL_Window)(w))
	if nativeSurface == nil {
		return nil, GetError()
	}
	return (*Surface)(nativeSurface), nil
}

func (w *Window) UpdateSurface() error {
	if C.SDL_UpdateWindowSurface((*C.struct_SDL_Window)(w)) != 0 {
		return GetError()
	}
	return nil
}

func (w *Window) PixelFormat() (PixelFormat, error) {
	format := PixelFormat(C.SDL_GetWindowPixelFormat((*C.struct_SDL_Window)(w)))
	if format == PixelFormatUnknown {
		return 0, GetError()
	}
	return format, nil
}

type SubsystemType C.SDL_SYSWM_TYPE

const (
	SubsystemUnknown  SubsystemType = C.SDL_SYSWM_UNKNOWN
	SubsystemWindows  SubsystemType = C.SDL_SYSWM_WINDOWS
	SubsystemX11      SubsystemType = C.SDL_SYSWM_X11
	SubsystemDirectFB SubsystemType = C.SDL_SYSWM_DIRECTFB
	SubsystemCocoa    SubsystemType = C.SDL_SYSWM_COCOA
	SubsystemUIKit    SubsystemType = C.SDL_SYSWM_UIKIT
	SubsystemWayland  SubsystemType = C.SDL_SYSWM_WAYLAND
	SubsystemMir      SubsystemType = C.SDL_SYSWM_MIR
	SubsystemWinRT    SubsystemType = C.SDL_SYSWM_WINRT
	SubsystemAndroid  SubsystemType = C.SDL_SYSWM_ANDROID
	SubsystemVivante  SubsystemType = C.SDL_SYSWM_VIVANTE
)

func (s SubsystemType) String() string {
	return map[SubsystemType]string{
		SubsystemWindows:  "Windows",
		SubsystemX11:      "X Window System",
		SubsystemDirectFB: "DirectFB",
		SubsystemCocoa:    "Apple Mac OS X",
		SubsystemUIKit:    "Apple iOS",
		SubsystemWayland:  "Wayland",
		SubsystemMir:      "Mir",
		SubsystemWinRT:    "WinRT",
		SubsystemAndroid:  "Android",
		SubsystemVivante:  "Vivante",
	}[s]
}

type WMInfo struct {
	Version   Version
	Subsystem SubsystemType
	_         [unsafe.Sizeof(C.struct_SDL_SysWMinfo{})]uint8
}

type WMInfoXlib struct {
	Version   Version
	Subsystem SubsystemType
	Display   uintptr
	Window    uintptr
}

type WMInfoCocoa struct {
	Version   Version
	Subsystem SubsystemType
	Window    uintptr
}

type WMInfoWin32 struct {
	Version       Version
	Subsystem     SubsystemType
	Window        uintptr
	DeviceContext uintptr
	Instance      uintptr
}

func (w *Window) GetWMInfo() (WMInfo, error) {
	info := WMInfo{
		Version: GetVersion(),
	}
	if C.SDL_GetWindowWMInfo((*C.struct_SDL_Window)(w), (*C.struct_SDL_SysWMinfo)(unsafe.Pointer(&info))) == 0 {
		return WMInfo{}, GetError()
	}
	return info, nil
}

func GetDisplayBounds(displayIndex int, rect *Rect) error {
	var r internal.Rect

	if C.SDL_GetDisplayBounds(C.int(displayIndex), (*C.struct_SDL_Rect)(unsafe.Pointer(&r))) != 0 {
		return GetError()
	}

	rect.fromInternal(r)

	return nil
}

// GetDisplayUsableBounds gets the usable desktop area represented by a display.
//
// The primary display (`displayIndex` zero) is always located at 0,0.
//
// This is the same area as GetDisplayBounds reports, but with portions
// reserved by the system removed. For example, on Apple's macOS, this
// subtracts the area occupied by the menu bar and dock.
//
// Setting a window to be fullscreen generally bypasses these unusable areas,
// so these are good guidelines for the maximum space available to a
// non-fullscreen window.
//
// The parameter `rect` is ignored if it is nil.
//
// This function also returns -1 if the parameter `displayIndex` is out of
// range.
//
// displayIndex is the index of the display to query the usable bounds from.
// rect is the Rect structure filled in with the display bounds
//
// This function is available since SDL 2.0.5.
func GetDisplayUsableBounds(displayIndex int, rect *Rect) error {
	var r internal.Rect
	if C.SDL_GetDisplayBounds(C.int(displayIndex), (*C.struct_SDL_Rect)(unsafe.Pointer(&r))) != 0 {
		return GetError()
	}
	rect.fromInternal(r)
	return nil
}

func GetNumVideoDisplays() (int, error) {
	num := int(C.SDL_GetNumVideoDisplays())
	if num < 0 {
		return 0, GetError()
	}
	return num, nil
}

func GetNumVideoDrivers() (int, error) {
	num := int(C.SDL_GetNumVideoDrivers())
	if num < 0 {
		return 0, GetError()
	}
	return num, nil
}

func GetVideoDriver(index int) string {
	return C.GoString(C.SDL_GetVideoDriver(C.int(index)))
}
