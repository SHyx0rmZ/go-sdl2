package sdl

// #include <SDL2/SDL_surface.h>
import "C"
import (
	"image"
	"image/color"
	"image/draw"
	"unsafe"

	"code.witches.io/go/sdl2/compat"
	"code.witches.io/go/sdl2/internal"
)

type Surface C.struct_SDL_Surface

func (s *Surface) FillRect(rect *Rect, color uint32) error {
	r := rect.toInternal()
	if C.SDL_FillRect((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(r)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) FillRects(rects []*Rect, color uint32) error {
	r := make([]*internal.Rect, len(rects))
	for i, rect := range rects {
		r[i] = rect.toInternal()
	}
	if C.SDL_FillRects((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(&r[0])), C.int(len(r)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) Free() {
	if s != nil {
		C.SDL_FreeSurface((*C.struct_SDL_Surface)(s))
	}
}

func (s *Surface) BlitSurface(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	srcR := srcRect.toInternal()
	dstR := dstRect.toInternal()
	r := int(C.SDL_BlitSurface((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(srcR)), (*C.struct_SDL_Surface)(dst), (*C.struct_SDL_Rect)(unsafe.Pointer(dstR))))
	if r != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) AlphaMod() (a uint8, err error) {
	if C.SDL_GetSurfaceAlphaMod((*C.struct_SDL_Surface)(s), (*C.Uint8)(unsafe.Pointer(&a))) != 0 {
		return 0, GetError()
	}
	return a, nil
}

func (s *Surface) ColorMod() (r, g, b uint8, err error) {
	if C.SDL_GetSurfaceColorMod((*C.struct_SDL_Surface)(s), (*C.Uint8)(unsafe.Pointer(&r)), (*C.Uint8)(unsafe.Pointer(&g)), (*C.Uint8)(unsafe.Pointer(&b))) != 0 {
		return 0, 0, 0, GetError()
	}
	return r, g, b, nil
}

func (s *Surface) SetAlphaMod(a uint8) error {
	if C.SDL_SetSurfaceAlphaMod((*C.struct_SDL_Surface)(s), C.Uint8(a)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) SetColorMod(r, g, b uint8) error {
	if C.SDL_SetSurfaceColorMod((*C.struct_SDL_Surface)(s), C.Uint8(r), C.Uint8(g), C.Uint8(b)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) ClipRect() Rect {
	var r internal.Rect
	C.SDL_GetClipRect((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(&r)))
	var rect Rect
	rect.fromInternal(r)
	return rect
}

func (s *Surface) SetClipRect(r *Rect) bool {
	return C.SDL_SetClipRect((*C.struct_SDL_Surface)(s), (*C.struct_SDL_Rect)(unsafe.Pointer(r.toInternal()))) == C.SDL_TRUE
}

func (s *Surface) Lock() error {
	if C.SDL_LockSurface((*C.struct_SDL_Surface)(s)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) Unlock() {
	C.SDL_UnlockSurface((*C.struct_SDL_Surface)(s))
}

func (s *Surface) ColorModel() color.Model {
	return color.NRGBAModel
}

func (s *Surface) Bounds() image.Rectangle {
	r := s.ClipRect()
	return image.Rect(r.X, r.Y, r.X+r.W, r.Y+r.H)
}

func (s *Surface) At(x, y int) color.Color {
	_surface := (*C.struct_SDL_Surface)(s)
	_format := (*PixelFormatS)(unsafe.Pointer(_surface.format))
	_pixel := *(*uint32)(unsafe.Pointer(uintptr(_surface.pixels) + uintptr(y)*uintptr(_surface.pitch) + uintptr(x)*uintptr(_format.bytesPerPixel)))
	r := uint8(((_pixel & _format.RMask) >> _format.RShift) << _format.RLoss)
	g := uint8(((_pixel & _format.GMask) >> _format.GShift) << _format.GLoss)
	b := uint8(((_pixel & _format.BMask) >> _format.BShift) << _format.BLoss)
	a := uint8(((_pixel & _format.AMask) >> _format.AShift) << _format.ALoss)
	//return color.RGBA{
	//	R: uint8(uint16(r) * uint16(a) / 255),
	//	G: uint8(uint16(g) * uint16(a) / 255),
	//	B: uint8(uint16(b) * uint16(a) / 255),
	//	A: a,
	//}
	return color.NRGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (s Surface) Format() PixelFormatS {
	return *(*PixelFormatS)(unsafe.Pointer(s.format))
}

func (s *Surface) Pixels() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(s.pixels)), s.Height()*s.Pitch())
}

func (s *Surface) Width() int {
	return int(s.w)
}

func (s *Surface) Height() int {
	return int(s.h)
}

func (s *Surface) Pitch() int {
	return int(s.pitch)
}

func (s *Surface) SetPixelFormatPalette(palette *Palette) error {
	return SetPixelFormatPalette((*PixelFormatS)(unsafe.Pointer(s.format)), palette)
}

func ConvertSurface(src *Surface, format *PixelFormatS) (*Surface, error) {
	surface := C.SDL_ConvertSurface(
		(*C.struct_SDL_Surface)(src),
		(*C.struct_SDL_PixelFormat)(unsafe.Pointer(format)),
		0,
	)
	if surface == nil {
		return nil, GetError()
	}
	return (*Surface)(surface), nil
}

func (s *Surface) Convert(format *PixelFormatS) (*Surface, error) {
	return ConvertSurface(s, format)
}

func ConvertSurfaceFormat(src *Surface, format PixelFormat) (*Surface, error) {
	surface := C.SDL_ConvertSurfaceFormat(
		(*C.struct_SDL_Surface)(src),
		C.Uint32(format),
		0,
	)
	if surface == nil {
		return nil, GetError()
	}
	return (*Surface)(surface), nil
}

func (s *Surface) ConvertFormat(format PixelFormat) (*Surface, error) {
	return ConvertSurfaceFormat(s, format)
}

func CreateRGBSurfaceWithFormat(flags uint32, width, height, depth int, format PixelFormat) (*Surface, error) {
	surface := C.SDL_CreateRGBSurfaceWithFormat(0, C.int(width), C.int(height), C.int(depth), C.Uint32(format))
	if surface == nil {
		return nil, GetError()
	}
	return (*Surface)(surface), nil
}

func (s *Surface) Image() draw.Image {
	f := s.Format()
	a := imageAdapter{
		surface: s,
	}
	a.model = a.ColorModel()
	a.bounds = a.Bounds()
	p := f.Palette()
	if p != nil {
		a2 := imageAdapterPalette{
			imageAdapter: a,
		}
		p := a2.ColorModel().(color.Palette)
		a2.palette = &p
		return a2
	}
	return a
}

type imageAdapter struct {
	surface *Surface
	model   color.Model
	bounds  image.Rectangle
}

func (a imageAdapter) ColorModel() color.Model {
	if PixelFormat(a.surface.format.format) == PixelFormatARGB2101010 {
		return compat.ARGB2101010Model
	}
	return &genericModel{(*PixelFormatS)(unsafe.Pointer(a.surface.format))}
}

func (a imageAdapter) Bounds() image.Rectangle {
	return image.Rect(
		0,
		0,
		int((*C.struct_SDL_Surface)(a.surface).w),
		int((*C.struct_SDL_Surface)(a.surface).h),
	)
}

func (a imageAdapter) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(a.bounds)) {
		if PixelFormat(a.surface.format.format) == PixelFormatARGB2101010 {
			return compat.ARGB2101010{ARGB: 0}
		}
		return genericColor{data: 0, info: (*PixelFormatS)(unsafe.Pointer(a.surface.format))}
	}
	bpp := int(a.surface.format.BytesPerPixel)
	i := y*int(a.surface.pitch) + x*bpp
	s := (*[4]byte)(unsafe.Add(unsafe.Pointer(a.surface.pixels), i))[:bpp:bpp]
	var data uint32
	switch bpp {
	case 4:
		data = uint32(s[3])<<24 | uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
	case 3:
		data = uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
	case 2:
		data = uint32(s[1])<<8 | uint32(s[0])
	case 1:
		data = uint32(s[0])
	}
	if PixelFormat(a.surface.format.format) == PixelFormatARGB2101010 {
		return compat.ARGB2101010{ARGB: data}
	}
	return genericColor{data: data, info: (*PixelFormatS)(unsafe.Pointer(a.surface.format))}
}

func (a imageAdapter) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(a.bounds)) {
		return
	}
	var data uint32
	switch c := a.model.Convert(c).(type) {
	case genericColor:
		data = c.data
	case compat.ARGB2101010:
		data = c.ARGB
	default:
		panic("unexpected color model")
	}
	bpp := int(a.surface.format.BytesPerPixel)
	i := y*int(a.surface.pitch) + x*bpp
	s := (*[4]byte)(unsafe.Add(unsafe.Pointer(a.surface.pixels), i))[:bpp:bpp]
	switch bpp {
	case 4:
		s[3] = uint8(data>>24) & 0xff
		fallthrough
	case 3:
		s[2] = uint8(data>>16) & 0xff
		fallthrough
	case 2:
		s[1] = uint8(data>>8) & 0xff
		fallthrough
	case 1:
		s[0] = uint8(data) & 0xff
	}
}

type imageAdapterPalette struct {
	imageAdapter
	palette *color.Palette
}

func (a imageAdapterPalette) At(x, y int) color.Color {
	if len(*a.palette) == 0 {
		return nil
	}
	return (*a.palette)[a.ColorIndexAt(x, y)]
}

func (a imageAdapterPalette) ColorIndexAt(x, y int) uint8 {
	if !(image.Point{x, y}.In(a.bounds)) {
		return 0
	}
	bpp := a.surface.format.BitsPerPixel
	by := uint(x*int(bpp)) / 8
	bi := uint(x*int(bpp)) % 8
	i := y*int(a.surface.pitch) + int(by)
	s := (*[2]byte)(unsafe.Add(unsafe.Pointer(a.surface.pixels), i))[:2:2]
	index := uint32(s[1])<<8 | uint32(s[0])
	return uint8(index>>bi) & ((1 << bpp) - 1)
}

func (a imageAdapterPalette) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(a.bounds)) {
		return
	}
	bpp := a.surface.format.BitsPerPixel
	by := uint(x*int(bpp)) / 8
	bi := uint(x*int(bpp)) % 8
	i := y*int(a.surface.pitch) + int(by)
	s := (*[2]byte)(unsafe.Add(unsafe.Pointer(a.surface.pixels), i))[:2:2]
	index := uint32(a.palette.Index(c))
	s[0] &= (1 << bi) - 1
	s[0] |= uint8((index << bi) & 0xff)
	s[1] &= ^((1 << ((bi + uint(bpp)) % 8)) - 1)
	s[1] |= uint8((index << bi) >> 8 & 0xff)
	// todo check if this is correct
}

func (a imageAdapterPalette) ColorModel() color.Model {
	f := a.surface.Format()
	p := f.Palette()
	cs := make(color.Palette, 0, len(p))
	for _, c := range p {
		cs = append(cs, color.NRGBA{R: c.R, G: c.G, B: c.B, A: c.A})
	}
	return cs
}

type genericColor struct {
	data uint32
	info *PixelFormatS
}

func (c genericColor) RGBA() (r, g, b, a uint32) {
	switch {
	case c.info.AMask == 0:
		a = 0xff
	default:
		a = c.data & c.info.AMask
		a >>= c.info.AShift
		a <<= c.info.ALoss
	}
	r = c.data & c.info.RMask
	r >>= c.info.RShift
	r <<= c.info.RLoss
	r |= r << 8
	r *= a
	r /= 0xff
	g = c.data & c.info.GMask
	g >>= c.info.GShift
	g <<= c.info.GLoss
	g |= g << 8
	g *= a
	g /= 0xff
	b = c.data & c.info.BMask
	b >>= c.info.BShift
	b <<= c.info.BLoss
	b |= b << 8
	b *= a
	b /= 0xff
	a |= a << 8
	return
}

type genericModel struct {
	info *PixelFormatS
}

func (m *genericModel) Convert(c color.Color) color.Color {
	if om, ok := c.(genericColor); ok && om.info == m.info {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0 {
		return genericColor{data: 0, info: m.info}
	}
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	r >>= m.info.RLoss
	r <<= m.info.RShift
	r &= m.info.RMask
	g >>= m.info.GLoss
	g <<= m.info.GShift
	g &= m.info.RMask
	b >>= m.info.BLoss
	b <<= m.info.BShift
	b &= m.info.BMask
	a >>= m.info.ALoss
	a <<= m.info.AShift
	a &= m.info.AMask
	return genericColor{data: r | g | b | a, info: m.info}
}
