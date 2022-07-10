package sdl

// #include <SDL2/SDL.h>
// void callback(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"io"
	"reflect"
	"sync"
	"unsafe"
)

type AudioCallback func(userData unsafe.Pointer, stream *byte, length int32)

type AudioSpec struct {
	Frequency int32
	Format    AudioFormat
	Channels  uint8
	Silence   uint8
	Samples   uint16
	Padding   uint16
	Size      uint32
	Callback  AudioCallback
	UserData  unsafe.Pointer
}

type audioSpec struct {
	Frequency int32
	Format    AudioFormat
	Channels  uint8
	Silence   uint8
	Samples   uint16
	Padding   uint16
	Size      uint32
	Callback  unsafe.Pointer
	UserData  uintptr
}

type AudioFormat uint16

const (
	AudioU8     AudioFormat = 0x0008
	AudioS8     AudioFormat = 0x8008
	AudioU16LSB AudioFormat = 0x0010
	AudioS16LSB AudioFormat = 0x8010
	AudioU16MSB AudioFormat = 0x1010
	AudioS16MSB AudioFormat = 0x9010
	AudioU16                = AudioU16LSB
	AudioS16                = AudioS16LSB
	AudioS32LSB AudioFormat = 0x8020
	AudioS32MSB AudioFormat = 0x9020
	AudioS32                = AudioS32LSB
	AudioF32LSB AudioFormat = 0x8120
	AudioF32MSB AudioFormat = 0x9120
	AudioF32                = AudioF32LSB
)

func GetNumAudioDevices(isCapture bool) int {
	var capture int
	if isCapture {
		capture = 1
	}
	return int(C.SDL_GetNumAudioDevices(C.int(capture)))
}

func GetAudioDeviceName(device int, isCapture bool) (string, error) {
	var capture int
	if isCapture {
		capture = 1
	}
	ptr := C.SDL_GetAudioDeviceName(C.int(device), C.int(capture))
	if ptr == nil {
		return "", GetError()
	}
	return C.GoString(ptr), nil
}

func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

func GetAudioDriver(index int) (string, error) {
	nativePtr := C.SDL_GetAudioDriver(C.int(index))
	if nativePtr == nil {
		return "", GetError()
	}
	return C.GoString(nativePtr), nil
}

type AudioDeviceID int

var callbackMutex sync.Mutex
var callbackToken uintptr
var callbackFuncs = make(map[uintptr]callback)

type callback struct {
	Callback AudioCallback
	UserData unsafe.Pointer
}

//export audioSpecCallback
func audioSpecCallback(userData unsafe.Pointer, stream *uint8, len int32) {
	callbackMutex.Lock()
	f := callbackFuncs[uintptr(userData)]
	callbackMutex.Unlock()
	if f.Callback == nil {
		panic("unknown callback func")
	}
	f.Callback(f.UserData, stream, len)
}

func OpenAudioDevice(device string, isCapture bool, desired, obtained *AudioSpec, allowedChanges int) (AudioDeviceID, error) {
	nativePtr := C.CString(device)
	defer C.free(unsafe.Pointer(nativePtr))
	var capture int
	if isCapture {
		capture = 1
	}
	_desired := audioSpec{
		Frequency: desired.Frequency,
		Format:    desired.Format,
		Channels:  desired.Channels,
		Silence:   desired.Silence,
		Samples:   desired.Samples,
		Padding:   desired.Padding,
		Size:      desired.Size,
	}
	var i uintptr
	if desired.Callback != nil {
		callbackMutex.Lock()
		callbackToken++
		i = callbackToken
		callbackFuncs[i] = struct {
			Callback AudioCallback
			UserData unsafe.Pointer
		}{
			Callback: desired.Callback,
			UserData: desired.UserData,
		}
		callbackMutex.Unlock()
		_desired.Callback = C.callback
		_desired.UserData = i
	}
	var _obtained audioSpec
	result := AudioDeviceID(C.SDL_OpenAudioDevice(
		nativePtr,
		C.int(capture),
		(*C.SDL_AudioSpec)(unsafe.Pointer(&_desired)),
		(*C.SDL_AudioSpec)(unsafe.Pointer(&_obtained)),
		C.int(allowedChanges),
	))
	if result == 0 {
		return 0, GetError()
	}
	*obtained = AudioSpec{
		Frequency: _obtained.Frequency,
		Format:    _obtained.Format,
		Channels:  _obtained.Channels,
		Silence:   _obtained.Silence,
		Samples:   _obtained.Samples,
		Padding:   _obtained.Padding,
		Size:      _obtained.Size,
	}
	if _obtained.Callback != nil {
		if _obtained.Callback != C.callback {
			panic("invalid callback")
		}
		f, ok := callbackFuncs[_obtained.UserData]
		if !ok {
			panic("unknown callback")
		}
		obtained.Callback = f.Callback
		obtained.UserData = f.UserData
	}
	return result, nil
}

func CloseAudioDevice(device AudioDeviceID) {
	C.SDL_CloseAudioDevice(C.SDL_AudioDeviceID(device))
}

func PauseAudioDevice(device AudioDeviceID, pause bool) {
	var pauseOn int
	if pause {
		pauseOn = 1
	}
	C.SDL_PauseAudioDevice(C.SDL_AudioDeviceID(device), C.int(pauseOn))
}

func QueueAudio(device AudioDeviceID, data unsafe.Pointer, length int) error {
	if C.SDL_QueueAudio(
		C.SDL_AudioDeviceID(device),
		data,
		C.Uint32(length),
	) != 0 {
		return GetError()
	}
	return nil
}

func LoadWAVRW(ops *RWOps, freeSrc bool, spec *AudioSpec) ([]byte, error) {
	var local reflect.StringHeader
	var free int
	if freeSrc {
		free = 1
	}
	ptr := (*AudioSpec)(unsafe.Pointer(C.SDL_LoadWAV_RW(
		(*C.struct_SDL_RWops)(unsafe.Pointer(ops)),
		C.int(free),
		(*C.struct_SDL_AudioSpec)(unsafe.Pointer(spec)),
		(**C.Uint8)(noescape(unsafe.Pointer(&local.Data))),
		(*C.Uint32)(noescape(unsafe.Pointer(&local.Len))),
	)))
	if ptr == nil {
		return nil, GetError()
	}
	length := uint32(local.Len)
	data := make([]byte, length)
	copy(data, *(*string)(unsafe.Pointer(&local)))
	C.SDL_FreeWAV(*(**C.Uint8)(unsafe.Pointer(&local.Data)))
	return data, nil
}

type audioStream C.SDL_AudioStream

type AudioStream interface {
	Put([]byte) error
	Get([]byte) (int, error)
	Available() int
	Flush() error
	Clear()
	Free()

	io.Writer
	io.Reader
	io.Closer
}

func (s *audioStream) Put(p []byte) error {
	if C.SDL_AudioStreamPut(
		(*C.SDL_AudioStream)(unsafe.Pointer(s)),
		unsafe.Pointer(&p[0]),
		C.int(len(p)),
	) != 0 {
		return GetError()
	}
	return nil
}

func (s *audioStream) Get(p []byte) (n int, err error) {
	result := int(C.SDL_AudioStreamGet(
		(*C.SDL_AudioStream)(unsafe.Pointer(s)),
		unsafe.Pointer(&p[0]),
		C.int(len(p)),
	))
	if result == -1 {
		return 0, GetError()
	}
	return result, nil
}

func (s *audioStream) Available() int {
	return int(C.SDL_AudioStreamAvailable((*C.SDL_AudioStream)(unsafe.Pointer(s))))
}

func (s *audioStream) Flush() error {
	if C.SDL_AudioStreamFlush((*C.SDL_AudioStream)(unsafe.Pointer(s))) != 0 {
		return GetError()
	}
	return nil
}

func (s *audioStream) Clear() {
	C.SDL_AudioStreamClear((*C.SDL_AudioStream)(unsafe.Pointer(s)))
}

func (s *audioStream) Free() {
	C.SDL_FreeAudioStream((*C.SDL_AudioStream)(unsafe.Pointer(s)))
}

func (s *audioStream) Write(p []byte) (n int, err error) {
	err = s.Put(p)
	return len(p), err
}

func (s *audioStream) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if s.Available() == 0 {
		return 0, io.EOF
	}
	return s.Get(p)
}

func (s *audioStream) Close() error {
	s.Free()
	return nil
}

func NewAudioStream(srcFormat AudioFormat, srcChannels, srcRate int, dstFormat AudioFormat, dstChannels, dstRate int) (AudioStream, error) {
	ptr := (*audioStream)(unsafe.Pointer(C.SDL_NewAudioStream(
		C.SDL_AudioFormat(srcFormat),
		C.Uint8(srcChannels),
		C.int(srcRate),
		C.SDL_AudioFormat(dstFormat),
		C.Uint8(dstChannels),
		C.int(dstRate),
	)))
	if ptr == nil {
		return nil, GetError()
	}
	return ptr, nil
}
