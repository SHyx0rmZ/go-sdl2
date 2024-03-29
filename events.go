package sdl

// #include <SDL2/SDL_events.h>
// #include "events.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"code.witches.io/go/sdl2/internal"
)

const (
	Query   = -1
	Ignore  = 0
	Disable = 0
	Enable  = 1
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
	EventControllerTouchpadDown
	EventControllerTouchpadMotion
	EventControllerTouchpadUp
	EventControllerSensorUpdate
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
	EventControllerTouchpadDown:         "controller touchpad down",
	EventControllerTouchpadMotion:       "controller touchpad motion",
	EventControllerTouchpadUp:           "controller touchpad up",
	EventControllerSensorUpdate:         "controller sensor update",
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

type WindowEvent struct {
	Type      EventType
	Timestamp time.Time
	WindowID  int
	Event     WindowEventID
	Data1     int
	Data2     int
}

type KeyboardEvent struct {
	Type      EventType
	Timestamp time.Time
	WindowID  int
	Pressed   bool
	Repeat    uint
	KeySymbol KeySymbol
}

type MouseMotionEvent struct {
	Type      EventType
	Timestamp time.Time
	WindowID  int
	Which     int
	Buttons   MouseButtons
	X         int
	Y         int
	DeltaX    int
	DeltaY    int
}

type MouseButtonEvent struct {
	Type      EventType
	Timestamp time.Time
	WindowID  int
	Which     int
	Button    int
	Pressed   bool
	Clicks    int
	X         int
	Y         int
}

type AudioDeviceEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     int
	IsCapture bool
}

type TouchFingerEvent struct {
	Type      EventType
	Timestamp time.Time
	TouchID   TouchID
	FingerID  FingerID
	X         float32
	Y         float32
	DeltaX    float32
	DeltaY    float32
	Pressure  float32
}

type MultiGestureEvent struct {
	Type       EventType
	Timestamp  time.Time
	TouchID    TouchID
	Angle      float32
	Distance   float32
	X          float32
	Y          float32
	NumFingers uint16
	Padding    uint16
}

type DollarGestureEvent struct {
	Type       EventType
	Timestamp  time.Time
	TouchID    TouchID
	GestureID  GestureID
	NumFingers uint32
	Error      float32
	X          float32
	Y          float32
}

type ControllerAxisEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     JoystickID
	Axis      uint8
	_         [3]uint8
	Value     int16
	_         uint16
}

type ControllerButtonEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     JoystickID
	Button    uint8
	State     uint8
	_         [2]uint8
}

type ControllerDeviceEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     int
}

type ControllerTouchpadEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     JoystickID
	Touchpad  int32
	Finger    int32
	X         float32
	Y         float32
	Pressure  float32
}

type ControllerSensorEvent struct {
	Type      EventType
	Timestamp time.Time
	Which     JoystickID
	Sensor    SensorType
	Data      [3]float32
}

var events = make(chan C.SDL_Event, 16)

//export propevent
func propevent(ev *C.SDL_Event) {
	if ev == nil {
		return
	}
	goev := *ev
	select {
	case events <- goev:
	case <-quit:
	}
}

//export didQuit
func didQuit(c chan<- struct{}) {
	c <- struct{}{}
}

var evo sync.Once
var quit chan struct{} = make(chan struct{})

func PollEvent() *CommonEvent {
	evo.Do(func() {
		sc := make(chan os.Signal, 2)
		//dq := make(chan struct{})
		internal.Cleanup.Add(1)
		var collectEventsQuit uint32
		go func() {
			select {
			case <-sc:
			case <-quit:
			}
			//y := (*uint32)(unsafe.Pointer(&C.collectEventsQuit))
			atomic.StoreUint32(&collectEventsQuit, 1)
			for {
				x := atomic.LoadUint32(&collectEventsQuit)
				//fmt.Println(x)
				if x == 2 {
					break
				}
				runtime.Gosched()
			}
			internal.Cleanup.Done()
		}()
		signal.Notify(sc, os.Interrupt, os.Kill)
		go C.collectEvents((*C.Uint32)(unsafe.Pointer(&collectEventsQuit))) //C.GoChan(unsafe.Pointer(&dq)))
	})

	var e C.SDL_Event

	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()

	//if C.SDL_PollEvent((*C.SDL_Event)(unsafe.Pointer(&e))) != 1 {
	//	return nil
	//}
	select {
	case e = <-events:

	default:
		return nil
	}

	_type := EventType(binary.LittleEndian.Uint32(e[0:4]))
	_timestamp := time.Duration(binary.LittleEndian.Uint32(e[4:8]))

	wrapper := &CommonEvent{
		Type:      EventType(binary.LittleEndian.Uint32(e[0:4])),
		Timestamp: timeInit.Add(_timestamp * time.Millisecond),
	}

	switch _type {
	case EventWindowEvent:
		wrapper.Event = WindowEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			WindowID:  int(binary.LittleEndian.Uint32(e[8:12])),
			Event:     WindowEventID(e[12]),
			Data1:     int(binary.LittleEndian.Uint32(e[16:20])),
			Data2:     int(binary.LittleEndian.Uint32(e[20:24])),
		}
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
	case EventMouseMotion:
		wrapper.Event = MouseMotionEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			WindowID:  int(binary.LittleEndian.Uint32(e[8:12])),
			Which:     int(binary.LittleEndian.Uint32(e[12:16])),
			Buttons: MouseButtons{
				Left:   (e[16]>>0)&1 == 1,
				Middle: (e[16]>>1)&1 == 1,
				Right:  (e[16]>>2)&1 == 1,
				X1:     (e[16]>>3)&1 == 1,
				X2:     (e[16]>>4)&1 == 1,
			},
			X:      int(int32(binary.LittleEndian.Uint32(e[20:24]))),
			Y:      int(int32(binary.LittleEndian.Uint32(e[24:28]))),
			DeltaX: int(int32(binary.LittleEndian.Uint32(e[28:32]))),
			DeltaY: int(int32(binary.LittleEndian.Uint32(e[32:36]))),
		}
	case EventMouseButtonDown:
		wrapper.Event = MouseButtonEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			WindowID:  int(binary.LittleEndian.Uint32(e[8:12])),
			Which:     int(binary.LittleEndian.Uint32(e[12:16])),
			Button:    int(e[16]),
			Pressed:   e[17] == 1,
			Clicks:    int(e[18]),
			X:         int(binary.LittleEndian.Uint32(e[20:24])),
			Y:         int(binary.LittleEndian.Uint32(e[24:28])),
		}
	case EventMouseButtonUp:
		wrapper.Event = MouseButtonEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			WindowID:  int(binary.LittleEndian.Uint32(e[8:12])),
			Which:     int(binary.LittleEndian.Uint32(e[12:16])),
			Button:    int(e[16]),
			Pressed:   e[17] == 1,
			Clicks:    int(e[18]),
			X:         int(binary.LittleEndian.Uint32(e[20:24])),
			Y:         int(binary.LittleEndian.Uint32(e[24:28])),
		}
	case EventFingerDown, EventFingerMotion, EventFingerUp:
		wrapper.Event = TouchFingerEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			TouchID:   TouchID(binary.LittleEndian.Uint64(e[8:16])),
			FingerID:  FingerID(binary.LittleEndian.Uint64(e[16:24])),
			X:         math.Float32frombits(binary.LittleEndian.Uint32(e[24:28])),
			Y:         math.Float32frombits(binary.LittleEndian.Uint32(e[28:32])),
			DeltaX:    math.Float32frombits(binary.LittleEndian.Uint32(e[32:36])),
			DeltaY:    math.Float32frombits(binary.LittleEndian.Uint32(e[36:40])),
			Pressure:  math.Float32frombits(binary.LittleEndian.Uint32(e[40:44])),
		}
		if !VersionAtLeast(2, 0, 7) {
			// todo: normalize values
		}
	case EventMultiGesture:
		wrapper.Event = MultiGestureEvent{
			Type:       wrapper.Type,
			Timestamp:  wrapper.Timestamp,
			TouchID:    TouchID(binary.LittleEndian.Uint64(e[8:16])),
			Angle:      math.Float32frombits(binary.LittleEndian.Uint32(e[16:20])),
			Distance:   math.Float32frombits(binary.LittleEndian.Uint32(e[20:24])),
			X:          math.Float32frombits(binary.LittleEndian.Uint32(e[24:28])),
			Y:          math.Float32frombits(binary.LittleEndian.Uint32(e[28:32])),
			NumFingers: binary.LittleEndian.Uint16(e[32:34]),
			Padding:    binary.LittleEndian.Uint16(e[34:36]),
		}
	case EventDollarGesture, EventDollarRecord:
		wrapper.Event = DollarGestureEvent{
			Type:       wrapper.Type,
			Timestamp:  wrapper.Timestamp,
			TouchID:    TouchID(binary.LittleEndian.Uint64(e[8:16])),
			GestureID:  GestureID(binary.LittleEndian.Uint64(e[16:24])),
			NumFingers: binary.LittleEndian.Uint32(e[24:28]),
			Error:      math.Float32frombits(binary.LittleEndian.Uint32(e[28:32])),
			X:          math.Float32frombits(binary.LittleEndian.Uint32(e[32:36])),
			Y:          math.Float32frombits(binary.LittleEndian.Uint32(e[36:40])),
		}
	case EventControllerAxisMotion:
		wrapper.Event = ControllerAxisEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Which:     JoystickID(binary.LittleEndian.Uint32(e[8:12])),
			Axis:      e[12],
			Value:     int16(binary.LittleEndian.Uint16(e[16:18])),
		}
	case EventControllerButtonDown, EventControllerButtonUp:
		wrapper.Event = ControllerButtonEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Which:     JoystickID(binary.LittleEndian.Uint32(e[8:12])),
			Button:    e[12],
			State:     e[13],
		}
	case EventControllerDeviceAdded, EventControllerDeviceRemapped, EventControllerDeviceRemoved:
		wrapper.Event = ControllerDeviceEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Which:     int(binary.LittleEndian.Uint32(e[8:12])),
		}
	case EventControllerTouchpadDown, EventControllerTouchpadMotion, EventControllerTouchpadUp:
		wrapper.Event = ControllerTouchpadEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Which:     JoystickID(binary.LittleEndian.Uint32(e[8:12])),
			Touchpad:  int32(binary.LittleEndian.Uint32(e[12:16])),
			Finger:    int32(binary.LittleEndian.Uint32(e[16:20])),
			X:         math.Float32frombits(binary.LittleEndian.Uint32(e[20:24])),
			Y:         math.Float32frombits(binary.LittleEndian.Uint32(e[24:28])),
			Pressure:  math.Float32frombits(binary.LittleEndian.Uint32(e[28:32])),
		}
	case EventControllerSensorUpdate:
		wrapper.Event = ControllerSensorEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Which:     JoystickID(binary.LittleEndian.Uint32(e[8:12])),
			Sensor:    SensorType(binary.LittleEndian.Uint32(e[12:16])),
			Data: [3]float32{
				math.Float32frombits(binary.LittleEndian.Uint32(e[16:20])),
				math.Float32frombits(binary.LittleEndian.Uint32(e[20:24])),
				math.Float32frombits(binary.LittleEndian.Uint32(e[24:28])),
			},
		}
	default:
		wrapper.Event = CommonEvent{
			Type:      wrapper.Type,
			Timestamp: wrapper.Timestamp,
			Event:     wrapper,
		}
	}

	return wrapper
}

type EventAction uint32

const (
	AddEvent  EventAction = C.SDL_ADDEVENT
	PeekEvent EventAction = C.SDL_PEEKEVENT
	GetEvent  EventAction = C.SDL_GETEVENT
)

func PeepEvents(n uint, action EventAction, min, max EventType) ([]Event, error) {
	_events := (*C.SDL_Event)(C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(C.SDL_Event{}))))
	defer C.free(unsafe.Pointer(_events))

	result := C.SDL_PeepEvents(_events, C.int(n), C.SDL_eventaction(action), C.Uint32(min), C.Uint32(max))
	if result < 1 {
		return nil, GetError()
	}

	events := make([]Event, result)
	return events, nil
}

func PumpEvents() {
	C.SDL_PumpEvents()
}

func PushEvent(event Event) (filtered bool, err error) {
	var e C.SDL_Event

	switch event := event.(type) {
	case CommonEvent:
		binary.LittleEndian.PutUint32(e[0:4], uint32(event.Type))
		binary.LittleEndian.PutUint32(e[4:8], uint32(event.Timestamp.UnixNano()*int64(time.Millisecond)/int64(time.Nanosecond)))
	default:
		return false, fmt.Errorf("not yet supported")
	}

	switch int(C.SDL_PushEvent(&e)) {
	case 0:
		return true, nil
	case 1:
		return false, nil
	default:
		return false, GetError()
	}
}

func (CommonEvent) eventFunc()             {}
func (WindowEvent) eventFunc()             {}
func (KeyboardEvent) eventFunc()           {}
func (MouseMotionEvent) eventFunc()        {}
func (MouseButtonEvent) eventFunc()        {}
func (AudioDeviceEvent) eventFunc()        {}
func (TouchFingerEvent) eventFunc()        {}
func (MultiGestureEvent) eventFunc()       {}
func (DollarGestureEvent) eventFunc()      {}
func (ControllerAxisEvent) eventFunc()     {}
func (ControllerButtonEvent) eventFunc()   {}
func (ControllerDeviceEvent) eventFunc()   {}
func (ControllerTouchpadEvent) eventFunc() {}
func (ControllerSensorEvent) eventFunc()   {}

func (e CommonEvent) String() string {
	return fmt.Sprintf("%s", e.Type)
}

func (e WindowEvent) String() string {
	return fmt.Sprintf("%s, window: %2d, data1: %16p, data2: %16p, event: %s", e.Type, e.WindowID, e.Data1, e.Data2, e.Event)
}

func (e KeyboardEvent) String() string {
	return fmt.Sprintf("%s, window: %2d, pressed: %5t, repeat: %2d, %v %v %v", e.Type, e.WindowID, e.Pressed, e.Repeat, e.KeySymbol.Scancode, e.KeySymbol.Keycode, e.KeySymbol.Modifiers)
}

func (e MouseMotionEvent) String() string {
	return fmt.Sprintf("%s, window: %2d, device: %2d, X: %4d, Y: %4d, ΔX: %+3d, ΔY: %+3d, buttons: %s", e.Type, e.WindowID, e.Which, e.X, e.Y, e.DeltaX, e.DeltaY, e.Buttons)
}

func (e MouseButtonEvent) String() string {
	return fmt.Sprintf("%s, window: %2d, device: %2d, button: %2d, pressed: %5t, clicks: %2d, X: %4d, Y: %4d", e.Type, e.WindowID, e.Which, e.Button, e.Pressed, e.Clicks, e.X, e.Y)
}

func (e AudioDeviceEvent) String() string {
	name, err := GetAudioDeviceName(e.Which, e.IsCapture)
	if err == nil {
		return fmt.Sprintf("%s, device: %2d, capture: %5t, name: %q", e.Type, e.Which, e.IsCapture, name)
	}
	return fmt.Sprintf("%s, device: %2d, capture: %5t", e.Type, e.Which, e.IsCapture)
}

func (e TouchFingerEvent) String() string {
	return fmt.Sprintf("%s, touch: %2d, finger: %2d, X: %6.2f, Y: %6.2f, ΔX: %+5.2f, ΔY: %+5.2f, pressure: %3.2f", e.Type, e.TouchID, e.FingerID, e.X, e.Y, e.DeltaX, e.DeltaY, e.Pressure)
}

func (e MultiGestureEvent) String() string {
	return fmt.Sprintf("%s, touch: %2d, rotation: %+5.4f, pinch: %+5.4f, X: %6.2f, Y: %6.2f, fingers: %2d, padding: %2d", e.Type, e.TouchID, e.Angle, e.Distance, e.X, e.Y, e.NumFingers, e.Padding)
}

func (e DollarGestureEvent) String() string {
	return fmt.Sprintf("%s, touch: %2d, gesture: %2d, fingers: %2d, error: %3.2f, X: %6.2f, Y: %6.2f", e.Type, e.TouchID, e.GestureID, e.NumFingers, e.Error, e.X, e.Y)
}

func (e ControllerAxisEvent) String() string {
	return fmt.Sprintf("%s, which: %2d, axis: %2d, value: %4d", e.Type, e.Which, e.Axis, e.Value)
}

func (e ControllerButtonEvent) String() string {
	return fmt.Sprintf("%s, which: %2d, button: %2d, state: %2d", e.Type, e.Which, e.Button, e.State)
}

func (e ControllerDeviceEvent) String() string {
	return fmt.Sprintf("%s, which: %2d", e.Type, e.Which)
}

func (e ControllerTouchpadEvent) String() string {
	return fmt.Sprintf("%s, which: %2d, touchpad: %2d, finger: %2d, X: %6.2f, Y: %6.2, pressure: %6.2f", e.Type, e.Which, e.Touchpad, e.Finger, e.X, e.Y, e.Pressure)
}

func (e ControllerSensorEvent) String() string {
	return fmt.Sprintf("%s, which: %2d, sensor: %2d, data: [%6.2f %6.2f %6.2f]", e.Type, e.Which, e.Sensor, e.Data[0], e.Data[1], e.Data[2])
}
