package sdl

// #include <SDL2/SDL.h>
import "C"

type KeySymbol struct {
	Scancode  ScanCode
	Keycode   KeyCode
	Modifiers KeyModifiers
}
