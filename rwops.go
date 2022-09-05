package sdl

// #include <SDL2/SDL_rwops.h>
//
// Sint64 RWsize(struct SDL_RWops *context) {
//   return context->size(context);
// }
//
// Sint64 RWseek(struct SDL_RWops *context, Sint64 offset, int whence) {
//   return context->seek(context, offset, whence);
// }
//
// size_t RWread(struct SDL_RWops *context, void *ptr, size_t size, size_t maxnum) {
//   return context->read(context, ptr, size, maxnum);
// }
//
// size_t RWwrite(struct SDL_RWops *context, const void *ptr, size_t size, size_t num) {
//   return context->write(context, ptr, size, num);
// }
//
// int RWclose(struct SDL_RWops *context) {
//   return context->close(context);
// }
import "C"
import (
	"io"
	"unsafe"
)

type RWOpsType uint32

const (
	RWOpsUnknown RWOpsType = iota
	RWOpsWin32File
	RWOpsStdFile
	RWOpsJNIFile
	RWOpsMemory
	RWOpsMemoryReadOnly
)

type RWOps C.struct_SDL_RWops

func RWFromFile(file, mode string) (*RWOps, error) {
	nativeFile := C.CString(file)
	defer C.free(unsafe.Pointer(nativeFile))

	nativeMode := C.CString(mode)
	defer C.free(unsafe.Pointer(nativeMode))

	nativeRWOps := C.SDL_RWFromFile(nativeFile, nativeMode)
	if nativeRWOps == nil {
		return nil, GetError()
	}

	return (*RWOps)(nativeRWOps), nil
}

func RWFromMem(mem unsafe.Pointer, size int) (*RWOps, error) {
	nativeRWOps := C.SDL_RWFromMem(mem, C.int(size))
	if nativeRWOps == nil {
		return nil, GetError()
	}

	return (*RWOps)(nativeRWOps), nil
}

func RWFromConstMem(mem unsafe.Pointer, size int) (*RWOps, error) {
	nativeRWOps := C.SDL_RWFromConstMem(mem, C.int(size))
	if nativeRWOps == nil {
		return nil, GetError()
	}

	return (*RWOps)(nativeRWOps), nil
}

func AllocRW() (*RWOps, error) {
	nativeRWOps := C.SDL_AllocRW()
	if nativeRWOps == nil {
		return nil, GetError()
	}

	return (*RWOps)(nativeRWOps), nil
}

func FreeRW(area *RWOps) {
	C.SDL_FreeRW((*C.struct_SDL_RWops)(area))
}

func (o *RWOps) Size() int64 {
	nativeRWOps := (*C.struct_SDL_RWops)(o)
	//return nativeRWOps.size.(func(*C.struct_SDL_RWops) C.int64)(nativeRWOps)
	return int64(C.RWsize(nativeRWOps))
}

func (o *RWOps) Seek(offset int64, whence int) (int64, error) {
	nativeRWOps := (*C.struct_SDL_RWops)(o)
	return int64(C.RWseek(nativeRWOps, C.Sint64(offset), C.int(whence))), nil
}

func (o *RWOps) Read(p []byte) (int, error) {
	nativeRWOps := (*C.struct_SDL_RWops)(o)
	n := int(C.RWread(nativeRWOps, unsafe.Pointer(&p[0]), 1, C.size_t(len(p))))
	if n == 0 {
		err := GetError()
		if err == nil {
			return 0, io.EOF
		}
		return 0, err
	}
	return n, nil
}

func (o *RWOps) Write(p []byte) (int, error) {
	nativeRWOps := (*C.struct_SDL_RWops)(o)
	n := int(C.RWwrite(nativeRWOps, unsafe.Pointer(&p[0]), 1, C.size_t(len(p))))
	if n != len(p) {
		return n, GetError()
	}
	return n, nil
}

func (o *RWOps) Close() error {
	nativeRWOps := (*C.struct_SDL_RWops)(o)
	r := C.RWclose(nativeRWOps)
	if r != 0 {
		return GetError()
	}
	return nil
}
