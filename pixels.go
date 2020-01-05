package sdl

// #include <SDL2/SDL.h>
import "C"
import (
	"unsafe"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

const (
	AlphaOpaque      = 255
	AlphaTransparent = 0
)

type PixelType uint8

const (
	PixelTypeUnknown PixelType = iota
	PixelTypeIndex1
	PixelTypeIndex4
	PixelTypeIndex8
	PixelTypePacked8
	PixelTypePacked16
	PixelTypePacked32
	PixelTypeArrayUnsigned8
	PixelTypeArrayUnsigned16
	PixelTypeArrayUnsigned32
	PixelTypeArrayFloat16
	PixelTypeArrayFloat32
)

// type order layout bits bytes

const (
	BitmapOrderNone = iota
	BitmapOrder4321
	BitmapOrder1234
)

const (
	PackedOrderNone = iota
	PackedOrderXRGB
	PackedOrderRGBX
	PackedOrderARGB
	PackedOrderRGBA
	PackedOrderXBGR
	PackedOrderBGRX
	PackedOrderABGR
	PackedOrderBGRA
)

const (
	ArrayOrderNone = iota
	ArrayOrderRGB
	ArrayOrderRGBA
	ArrayOrderARGB
	ArrayOrderBGR
	ArrayOrderBGRA
	ArrayOrderABGR
)

const (
	PackedLayoutNone = iota
	PackedLayout332
	PackedLayout4444
	PackedLayout1555
	PackedLayout5551
	PackedLayout565
	PackedLayout8888
	PackedLayout2101010
	PackedLayout1010102
)

const (
	PixelFormatUnknown     PixelFormat = 0
	PixelFormatIndex1LSB               = (1 << 28) | (PixelFormat(PixelTypeIndex1) << 24) | (BitmapOrder4321 << 20) | (PackedLayoutNone << 16) | (1 << 8) | 0
	PixelFormatIndex1MSB               = (1 << 28) | (PixelFormat(PixelTypeIndex1) << 24) | (BitmapOrder1234 << 20) | (PackedLayoutNone << 16) | (1 << 8) | 0
	PixelFormatIndex4LSB               = (1 << 28) | (PixelFormat(PixelTypeIndex4) << 24) | (BitmapOrder4321 << 20) | (PackedLayoutNone << 16) | (4 << 8) | 0
	PixelFormatIndex4MSB               = (1 << 28) | (PixelFormat(PixelTypeIndex4) << 24) | (BitmapOrder1234 << 20) | (PackedLayoutNone << 16) | (4 << 8) | 0
	PixelFormatIndex8                  = (1 << 28) | (PixelFormat(PixelTypeIndex8) << 24) | (BitmapOrderNone << 20) | (PackedLayoutNone << 16) | (8 << 8) | 1
	PixelFormatRGB332                  = (1 << 28) | (PixelFormat(PixelTypePacked8) << 24) | (PackedOrderXRGB << 20) | (PackedLayout332 << 16) | (8 << 8) | 1
	PixelFormatRGB444                  = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderXRGB << 20) | (PackedLayout4444 << 16) | (12 << 8) | 2
	PixelFormatRGB555                  = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderXRGB << 20) | (PackedLayout1555 << 16) | (15 << 8) | 2
	PixelFormatBGR555                  = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderXBGR << 20) | (PackedLayout1555 << 16) | (15 << 8) | 2
	PixelFormatARGB4444                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderARGB << 20) | (PackedLayout4444 << 16) | (16 << 8) | 2
	PixelFormatRGBA4444                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderRGBA << 20) | (PackedLayout4444 << 16) | (16 << 8) | 2
	PixelFormatABGR4444                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderABGR << 20) | (PackedLayout4444 << 16) | (16 << 8) | 2
	PixelFormatBGRA4444                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderBGRA << 20) | (PackedLayout4444 << 16) | (16 << 8) | 2
	PixelFormatARGB1555                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderARGB << 20) | (PackedLayout1555 << 16) | (16 << 8) | 2
	PixelFormatRGBA5551                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderRGBA << 20) | (PackedLayout5551 << 16) | (16 << 8) | 2
	PixelFormatABGR1555                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderABGR << 20) | (PackedLayout1555 << 16) | (16 << 8) | 2
	PixelFormatBGRA5551                = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderBGRA << 20) | (PackedLayout5551 << 16) | (16 << 8) | 2
	PixelFormatRGB565                  = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderXRGB << 20) | (PackedLayout565 << 16) | (16 << 8) | 2
	PixelFormatBGR565                  = (1 << 28) | (PixelFormat(PixelTypePacked16) << 24) | (PackedOrderXBGR << 20) | (PackedLayout565 << 16) | (16 << 8) | 2
	PixelFormatRGB24                   = (1 << 28) | (PixelFormat(PixelTypeArrayUnsigned8) << 24) | (ArrayOrderRGB << 20) | (PackedLayoutNone << 16) | (24 << 8) | 3
	PixelFormatBGR24                   = (1 << 28) | (PixelFormat(PixelTypeArrayUnsigned8) << 24) | (ArrayOrderBGR << 20) | (PackedLayoutNone << 16) | (24 << 8) | 3
	PixelFormatRGB888                  = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderXRGB << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatRGBX8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderRGBX << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatBGR888                  = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderXBGR << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatBGRX8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderBGRX << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatARGB8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderARGB << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatRGBA8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderRGBA << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatABGR8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderABGR << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatBGRA8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderBGRA << 20) | (PackedLayout8888 << 16) | (32 << 8) | 4
	PixelFormatARGB2101010             = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderARGB << 20) | (PackedLayout2101010 << 16) | (32 << 8) | 4

	// Planar mode: Y + V + U (3 planes)
	PixelFormatYV12 PixelFormat = ('2' << 24) | ('1' << 16) | ('V' << 8) | 'Y'
	// Planar mode: Y + U + V (3 planes)
	PixelFormatIYUV PixelFormat = ('V' << 24) | ('U' << 16) | ('Y' << 8) | 'I'
	// Packed mode: Y0+U0+Y1+V0 (1 plane)
	PixelFormatYUY2 PixelFormat = ('2' << 24) | ('Y' << 16) | ('U' << 8) | 'Y'
	// Packed mode: U0+Y0+V0+Y1 (1 plane)
	PixelFormatUYVY PixelFormat = ('Y' << 24) | ('V' << 16) | ('Y' << 8) | 'U'
	// Packed mode: Y0+V0+Y1+U0 (1 plane)
	PixelFormatYVYU PixelFormat = ('U' << 24) | ('Y' << 16) | ('V' << 8) | 'Y'
	// Planar mode: Y + U/V interleaved (2 planes)
	PixelFormatNV12 PixelFormat = ('2' << 24) | ('1' << 16) | ('V' << 8) | 'N'
	// Planar mode: Y + V/U interleaved (2 planes)
	PixelFormatNV21 PixelFormat = ('1' << 24) | ('2' << 16) | ('V' << 8) | 'N'
	// Android video texture format
	PixelFormatExternalOES PixelFormat = (' ' << 24) | ('S' << 16) | ('E' << 8) | 'O'
)

var (
	PixelFormatRGBA32 PixelFormat
	PixelFormatARGB32 PixelFormat
	PixelFormatBGRA32 PixelFormat
	PixelFormatABGR32 PixelFormat
)

func init() {
	switch "LittleEndian" {
	case "BigEndian":
		PixelFormatRGBA32 = PixelFormatRGBA8888
		PixelFormatARGB32 = PixelFormatARGB8888
		PixelFormatBGRA32 = PixelFormatBGRA8888
		PixelFormatABGR32 = PixelFormatABGR8888
	case "LittleEndian":
		PixelFormatRGBA32 = PixelFormatABGR8888
		PixelFormatARGB32 = PixelFormatBGRA8888
		PixelFormatBGRA32 = PixelFormatARGB8888
		PixelFormatABGR32 = PixelFormatRGBA8888
	default:
		panic("invalid endianness")
	}
}

type PixelFormat uint32

func (f PixelFormat) IsFourCC() bool {
	return f != 0 && ((f>>28)&0xf) != 1
}

func (f PixelFormat) PixelType(format int) PixelType {
	return PixelType((format >> 24) & 0xf)
}

func PixelOrder(format int) int {
	return (format >> 20) & 0xf
}

func PixelLayout(format int) int {
	return (format >> 16) & 0xf
}

func BitsPerPixel(format int) int {
	return (format >> 8) & 0xf
}

type Palette struct {
	nColors  int32
	colors   *Color
	version  uint32
	refCount int32
}

type PixelFormatS struct {
	format        PixelFormat
	palette       *Palette
	bitsPerPixel  uint8
	bytesPerPixel uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32

	rLoss    uint8
	gLoss    uint8
	bLoss    uint8
	aLoss    uint8
	rShift   uint8
	gShift   uint8
	bShift   uint8
	aShift   uint8
	refCount int32
	next     *PixelFormatS
}

func AllocFormat(format PixelFormat) (*PixelFormatS, error) {
	nativeFormat := C.SDL_AllocFormat(C.Uint32(format))
	if nativeFormat == nil {
		return nil, GetError()
	}
	return (*PixelFormatS)(unsafe.Pointer(nativeFormat)), nil
}

func (f *PixelFormatS) Free() {
	C.SDL_FreeFormat((*C.struct_SDL_PixelFormat)(unsafe.Pointer(f)))
}

func MapRGB(format *PixelFormatS, r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB((*C.struct_SDL_PixelFormat)(unsafe.Pointer(format)), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

func MapRGBA(format *PixelFormatS, r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA((*C.struct_SDL_PixelFormat)(unsafe.Pointer(format)), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

func (f PixelFormatS) Format() PixelFormat {
	return f.format
}

func (f PixelFormatS) BytesPerPixel() int {
	return int(f.bytesPerPixel)
}
