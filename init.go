package sdl

// #include <SDL2/SDL.h>
import "C"
import "time"

type InitFlag uint32

const (
	InitTimer          InitFlag = 0x00000001
	InitAudio          InitFlag = 0x00000010
	InitVideo          InitFlag = 0x00000020
	InitJoystick       InitFlag = 0x00000200
	InitHaptic         InitFlag = 0x00001000
	InitGameController InitFlag = 0x00002000
	InitEvents         InitFlag = 0x00004000
	InitNoParachute    InitFlag = 0x00100000
	InitEverything     InitFlag = InitTimer | InitAudio | InitVideo | InitJoystick | InitHaptic | InitGameController | InitEvents
)

var timeInit time.Time

func Init(flags InitFlag) error {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return GetError()
	}
	ClearError()
	timeInit = time.Now()
	return nil
}

func InitSubSystem(flags InitFlag) error {
	if C.SDL_InitSubSystem(C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

func Quit() {
	C.SDL_Quit()
}

func QuitSubSystem(flags InitFlag) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}
