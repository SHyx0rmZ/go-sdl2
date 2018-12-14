package sdl
import "C"

// #include <SDL2/SDL.h>
import "C"
import (
	"unsafe"
)

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func GetVersion() Version {
	var v Version
	C.SDL_GetVersion((*C.struct_SDL_version)(unsafe.Pointer(&v)))
	return v
}

func GetRevision() string {
	nativeRevision := C.SDL_GetRevision()
	return C.GoString(nativeRevision)
}

func GetRevisionNumber() int {
	return int(C.SDL_GetRevisionNumber())
}

func VersionAtLeast(major, minor, patch uint8) bool {
	v := GetVersion()
	return  VersionNum(v.Major, v.Minor, v.Patch) >= VersionNum(major, minor, patch)
}

func VersionNum(major, minor, patch uint8) int {
	return int(major) * 1000 + int(minor) * 100 + int(patch)
}
