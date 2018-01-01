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

var eventTypeToStringMap = map[EventType]string{
	EventFirstEvent:                     "first event",
	EventQuit:                           "quit",
	EventApplicationTerminating:         "application terminating",
	EventApplicationLowMemory:           "application low memory",
	EventApplicationWillEnterBackground: "application will enter background",
	EventApplicationDidEnterBackground:  "application did enter background",
	EventApplicationWillEnterForeground: "application will enter foreground",
	EventApplicationDidEnterForeground:  "application did enter foreground",
	EventWindowEvent:                    "window event",
	EventSystemWindowManagerEvent:       "system window manager event",
	EventKeyDown:                        "key down",
	EventKeyUp:                          "key up",
	EventTextEditing:                    "text editing",
	EventTextInput:                      "text input",
	EventKeyMapChanged:                  "key map changed",
	EventMouseMotion:                    "mouse motion",
	EventMouseButtonDown:                "mouse button down",
	EventMouseButtonUp:                  "mouse button up",
	EventMouseWheel:                     "mouse wheel",
	EventJoystickAxisMotion:             "joystick axis motion",
	EventJoystickTrackballMotion:        "joystick trackball motion",
	EventJoystickHatMotion:              "joystick hat motion",
	EventJoystickButtonDown:             "joystick button down",
	EventJoystickButtonUp:               "joystick button up",
	EventJoystickDeviceAdded:            "joystick device added",
	EventJoystickDeviceRemoved:          "joystick device removed",
	EventControllerAxisMotion:           "controller axis motion",
	EventControllerButtonDown:           "controller button down",
	EventControllerButtonUp:             "controller button up",
	EventControllerDeviceAdded:          "controller device added",
	EventControllerDeviceRemoved:        "controller device removed",
	EventControllerDeviceRemapped:       "controller device remapped",
	EventFingerDown:                     "finger down",
	EventFingerUp:                       "finger up",
	EventFingerMotion:                   "finger motion",
	EventDollarGesture:                  "dollar gesture",
	EventDollarRecord:                   "dollar record",
	EventMultiGesture:                   "multi gesture",
	EventClipboardUpdate:                "clipboard update",
	EventDropFile:                       "drop file",
	EventDropText:                       "drop text",
	EventDropBegin:                      "drop begin",
	EventDropComplete:                   "drop complete",
	EventAudioDeviceAdded:               "audio device added",
	EventAudioDeviceRemoved:             "audio device removed",
	EventRenderTargetsReset:             "render targets reset",
	EventRenderDeviceReset:              "render device reset",
}

func (t EventType) String() string {
	s, ok := eventTypeToStringMap[t]
	if !ok {
		if t >= EventUserEvent && t <= EventLastEvent {
			return "user event"
		}
		return "unknown event"
	}
	return s
}

type Event interface {
	eventFunc()
}

type CommonEvent struct {
	Type      EventType
	Timestamp time.Time
	Event     Event
}

type KeyboardEvent struct {
	Type      EventType
	Timestamp time.Time
	WindowID  int
	Pressed   bool
	Repeat    uint
	KeySymbol KeySymbol
}

type AudioDeviceEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     int
	IsCapture bool
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
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			WindowID:  int(binary.LittleEndian.Uint32(e[8:12])),
			Pressed:   bool(e[12] != 0),
			Repeat:    uint(e[13]),
			KeySymbol: KeySymbol{
				Scancode:  ScanCode(binary.LittleEndian.Uint32(e[16:20])),
				Keycode:   KeyCode(binary.LittleEndian.Uint32(e[20:24])),
				Modifiers: KeyModifiers(binary.LittleEndian.Uint16(e[24:26])),
			},
		}
	case EventAudioDeviceAdded, EventAudioDeviceRemoved:
		wrapper.Event = AudioDeviceEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Which:     int(binary.LittleEndian.Uint32(e[8:12])),
			IsCapture: bool(e[12] != 0),
		}
	}

	return wrapper
}

func (KeyboardEvent) eventFunc()    {}
func (AudioDeviceEvent) eventFunc() {}
