package sdl

// #include <SDL2/SDL_pixels.h>
import "C"
import (
	"strconv"
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
	PixelFormatRGB888                  = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderXRGB << 20) | (PackedLayout8888 << 16) | (24 << 8) | 4
	PixelFormatRGBX8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderRGBX << 20) | (PackedLayout8888 << 16) | (24 << 8) | 4
	PixelFormatBGR888                  = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderXBGR << 20) | (PackedLayout8888 << 16) | (24 << 8) | 4
	PixelFormatBGRX8888                = (1 << 28) | (PixelFormat(PixelTypePacked32) << 24) | (PackedOrderBGRX << 20) | (PackedLayout8888 << 16) | (24 << 8) | 4
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

var pixelFormatNames = map[PixelFormat]string{
	PixelFormatIndex1LSB:   "INDEX1LSB",
	PixelFormatIndex1MSB:   "INDEX1MSB",
	PixelFormatIndex4LSB:   "INDEX4LSB",
	PixelFormatIndex4MSB:   "INDEX4MSB",
	PixelFormatIndex8:      "INDEX8",
	PixelFormatRGB332:      "RGB332",
	PixelFormatRGB444:      "RGB444",
	PixelFormatRGB555:      "RGB555",
	PixelFormatBGR555:      "BGR555",
	PixelFormatARGB4444:    "ARGB4444",
	PixelFormatRGBA4444:    "RGBA4444",
	PixelFormatABGR4444:    "ABGR4444",
	PixelFormatBGRA4444:    "BGRA4444",
	PixelFormatARGB1555:    "ARGB1555",
	PixelFormatRGBA5551:    "RGBA5551",
	PixelFormatABGR1555:    "ABGR1555",
	PixelFormatBGRA5551:    "BGRA5551",
	PixelFormatRGB565:      "RGB565",
	PixelFormatBGR565:      "BGR565",
	PixelFormatRGB24:       "RGB24",
	PixelFormatBGR24:       "BGR24",
	PixelFormatRGB888:      "RGB888",
	PixelFormatRGBX8888:    "RGBX8888",
	PixelFormatBGR888:      "BGR888",
	PixelFormatBGRX8888:    "BGRX8888",
	PixelFormatARGB8888:    "ARGB8888",
	PixelFormatRGBA8888:    "RGBA8888",
	PixelFormatABGR8888:    "ABGR8888",
	PixelFormatBGRA8888:    "BGRA8888",
	PixelFormatARGB2101010: "ARGB2101010",
	PixelFormatYV12:        "YV12",
	PixelFormatIYUV:        "IYUV",
	PixelFormatYUY2:        "YUY2",
	PixelFormatUYVY:        "UYVY",
	PixelFormatYVYU:        "YVYU",
	PixelFormatNV12:        "NV12",
	PixelFormatNV21:        "NV21",
	PixelFormatExternalOES: "OES",
}

func (f PixelFormat) String() string {
	name, ok := pixelFormatNames[f]
	if !ok {
		return "UNKNOWN" + "(" + strconv.Itoa(int(f)) + ")"
	}
	return name
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
	RMask         uint32
	GMask         uint32
	BMask         uint32
	AMask         uint32

	RLoss    uint8
	GLoss    uint8
	BLoss    uint8
	ALoss    uint8
	RShift   uint8
	GShift   uint8
	BShift   uint8
	AShift   uint8
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

func (f PixelFormatS) Palette() []Color {
	if f.palette == nil {
		return nil
	}
	colors := make([]Color, f.palette.nColors)
	copy(colors, unsafe.Slice(f.palette.colors, f.palette.nColors))
	return colors
}

func AllocPalette(colors int) (*Palette, error) {
	palette := C.SDL_AllocPalette(C.int(colors))
	if palette == nil {
		return nil, GetError()
	}
	return (*Palette)(unsafe.Pointer(palette)), nil
}

func FreePalette(palette *Palette) {
	C.SDL_FreePalette((*C.struct_SDL_Palette)(unsafe.Pointer(palette)))
}

func (p *Palette) Free() {
	FreePalette(p)
}

func SetPaletteColors(palette *Palette, offset int, colors []Color) error {
	if len(colors) == 0 {
		return nil
	}
	if C.SDL_SetPaletteColors(
		(*C.struct_SDL_Palette)(unsafe.Pointer(palette)),
		(*C.struct_SDL_Color)(unsafe.Pointer(&colors[0])),
		C.int(offset),
		C.int(len(colors)),
	) != 0 {
		return GetError()
	}
	return nil
}

func (p *Palette) SetColors(offset int, colors []Color) error {
	return SetPaletteColors(p, offset, colors)
}

func SetPixelFormatPalette(format *PixelFormatS, palette *Palette) error {
	if C.SDL_SetPixelFormatPalette(
		(*C.struct_SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.struct_SDL_Palette)(unsafe.Pointer(palette)),
	) != 0 {
		return GetError()
	}
	return nil
}

func (f *PixelFormatS) SetPalette(palette *Palette) error {
	return SetPixelFormatPalette(f, palette)
}
