package sdl

// #cgo windows LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
import "C"
import (
	"encoding/binary"
	"time"
	"unsafe"
)

type EventType uint32

const (
	EventFirstEvent EventType = iota
)

const (
	EventQuit EventType = 0x100 + iota
	EventApplicationTerminating
	EventApplicationLowMemory
	EventApplicationWillEnterBackground
	EventApplicationDidEnterBackground
	EventApplicationWillEnterForeground
	EventApplicationDidEnterForeground
)

const (
	EventWindowEvent EventType = 0x200 + iota
	EventSystemWindowManagerEvent
)

const (
	EventKeyDown EventType = 0x300 + iota
	EventKeyUp
	EventTextEditing
	EventTextInput
	EventKeyMapChanged
)

const (
	EventMouseMotion EventType = 0x400 + iota
	EventMouseButtonDown
	EventMouseButtonUp
	EventMouseWheel
)

const (
	EventJoystickAxisMotion EventType = 0x600 + iota
	EventJoystickTrackballMotion
	EventJoystickHatMotion
	EventJoystickButtonDown
	EventJoystickButtonUp
	EventJoystickDeviceAdded
	EventJoystickDeviceRemoved
)

const (
	EventControllerAxisMotion EventType = 0x650 + iota
	EventControllerButtonDown
	EventControllerButtonUp
	EventControllerDeviceAdded
	EventControllerDeviceRemoved
	EventControllerDeviceRemapped
)

const (
	EventFingerDown EventType = 0x700 + iota
	EventFingerUp
	EventFingerMotion
)

const (
	EventDollarGesture EventType = 0x800 + iota
	EventDollarRecord
	EventMultiGesture
)

const (
	EventClipboardUpdate EventType = 0x900 + iota
)

const (
	EventDropFile EventType = 0x1000 + iota
	EventDropText
	EventDropBegin
	EventDropComplete
)

const (
	EventAudioDeviceAdded EventType = 0x1100 + iota
	EventAudioDeviceRemoved
)

const (
	EventRenderTargetsReset EventType = 0x2000 + iota
	EventRenderDeviceReset
)

const (
	EventUserEvent EventType = 0x8000 + iota
)

const (
	EventLastEvent EventType = 0xffff
)

type Event interface {
	eventFunc()
}

type CommonEvent struct {
	Type      EventType
	Timestamp time.Time
	Event
}

type KeyboardEvent struct {
	WindowID  int
	Pressed   bool
	Repeat    uint
	KeySymbol KeySymbol
}

func PollEvent() *CommonEvent {
	var e C.SDL_Event

	if C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(&e))) != 1 {
		return nil
	}

	_type := EventType(binary.LittleEndian.Uint32(e[0:4]))
	_timestamp := time.Duration(binary.LittleEndian.Uint32(e[4:8]))

	wrapper := &CommonEvent{
		Type:      EventType(binary.LittleEndian.Uint32(e[0:4])),
		Timestamp: timeInit.Add(_timestamp * time.Millisecond),
	}

	switch _type {
	case EventKeyDown, EventKeyUp:
		wrapper.Event = KeyboardEvent{
			WindowID: int(binary.LittleEndian.Uint32(e[8:12])),
			Pressed:  !bool(e[12] != 0),
			Repeat:   uint(e[13]),
			KeySymbol: KeySymbol{
				Scancode:  ScanCode(binary.LittleEndian.Uint32(e[16:20])),
				Keycode:   KeyCode(binary.LittleEndian.Uint32(e[20:24])),
				Modifiers: KeyModifiers(binary.LittleEndian.Uint16(e[24:26])),
			},
		}
	}

	return wrapper
}

func (KeyboardEvent) eventFunc() {}
