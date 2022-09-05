package sdl

// #include <SDL2/SDL_gamecontroller.h>
import "C"
import (
	"unsafe"
)

type GameController = C.SDL_GameController

func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

func GameControllerNumMappings() int {
	return int(C.SDL_GameControllerNumMappings())
}

func GameControllerMappingForIndex(mappingIndex int) string {
	str := C.SDL_GameControllerMappingForIndex(C.int(mappingIndex))
	if str == nil {
		return ""
	}
	defer C.SDL_free(unsafe.Pointer(str))
	return C.GoString(str)
}

func GameControllerMapping(gameController *GameController) (string, error) {
	str := C.SDL_GameControllerMapping(gameController)
	if str == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(str))
	return C.GoString(str), nil
}

func IsGameController(joystickIndex int) bool {
	return C.SDL_IsGameController(C.int(joystickIndex)) == C.SDL_TRUE
}

func GameControllerOpen(joystickIndex int) (*GameController, error) {
	ptr := C.SDL_GameControllerOpen(C.int(joystickIndex))
	if ptr == nil {
		return nil, GetError()
	}
	return (*GameController)(unsafe.Pointer(ptr)), nil
}

func GameControllerClose(controller *GameController) {
	C.SDL_GameControllerClose(controller)
}

func GameControllerFromInstanceID(joyID JoystickID) (*GameController, error) {
	ptr := C.SDL_GameControllerFromInstanceID((C.SDL_JoystickID)(joyID))
	if ptr == nil {
		return nil, GetError()
	}
	return (*GameController)(unsafe.Pointer(ptr)), nil
}

func GameControllerHasSensor(controller *GameController, typ SensorType) bool {
	return C.SDL_GameControllerHasSensor(controller, (C.SDL_SensorType)(typ)) == C.SDL_TRUE
}

func GameControllerSetSensorEnabled(controller *GameController, typ SensorType, enabled bool) error {
	_enabled := C.SDL_FALSE
	if enabled {
		_enabled = C.SDL_TRUE
	}
	if C.SDL_GameControllerSetSensorEnabled(controller, (C.SDL_SensorType)(typ), (C.SDL_bool)(_enabled)) != 0 {
		return GetError()
	}
	return nil
}

func GameControllerIsSensorEnabled(controller *GameController, typ SensorType) bool {
	return C.SDL_GameControllerIsSensorEnabled(controller, (C.SDL_SensorType)(typ)) == C.SDL_TRUE
}
